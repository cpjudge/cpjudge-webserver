package actions

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"

	leaderboardClient "github.com/cpjudge/cpjudge_webserver/proto/leaderboard"
	"github.com/gobuffalo/buffalo"
	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
)

type websocketWithChannels struct {
	websocket      *websocket.Conn
	sendChannel    chan *leaderboardClient.Contest
	receiveChannel chan *leaderboardClient.Participants
	contestID      string
}

var leaderboardMap map[string]([]*websocketWithChannels)
var leaderboardConn *grpc.ClientConn
var client leaderboardClient.LeaderboardClient

// LeaderboardHandler : fetch leaderboard for contest
func LeaderboardHandler(c buffalo.Context) error {
	if leaderboardMap == nil {
		leaderboardMap = make(map[string]([]*websocketWithChannels))
	}
	contestID := c.Request().URL.Query().Get("contest_id")
	if contestID != "" {
		contest := &leaderboardClient.Contest{
			ContestId: contestID,
		}
		log.Printf(contest.String())
		// 	if leaderboardMap[contest.ContestId] == nil {
		// 		leaderboardMap[contest.ContestId] = make([]chan *leaderboardClient.Contest, 0)
		// 	}
		// 	channel := make(chan *leaderboardClient.Contest)
		// 	leaderboardMap[contest.ContestId] = append(leaderboardMap[contest.ContestId], channel)
		// 	handleConnection(leaderboardClient.GetLeaderboard(contest), contest, c, channel)

		// }
		// return c.Render(400, r.JSON(map[string]interface{}{
		// 	"message": "contest_id parameter missing",
		// }))
		if leaderboardConn == nil {
			leaderboardConn = leaderboardClient.GetLeaderboard()
			client = leaderboardClient.NewLeaderboardClient(leaderboardConn)
		}
		websocketConn, err := GetWebsocketConnection(c)
		if err != nil {
			return c.Render(500, r.JSON(map[string]interface{}{
				"message": err.Error(),
			}))
		}
		wwc := &websocketWithChannels{
			websocket:      websocketConn,
			sendChannel:    make(chan *leaderboardClient.Contest),
			receiveChannel: make(chan *leaderboardClient.Participants),
			contestID:      contestID,
		}
		leaderboardMap[contest.ContestId] = append(leaderboardMap[contest.ContestId], wwc)
		go handleConnection(wwc)
		wwc.sendChannel <- contest
		go onReceiveLeaderboard(wwc)
		go onSendContestID(wwc)
	}
	return nil
}

// TriggerLeaderboards : Trigger all leaderboards to update their status
func TriggerLeaderboards(contestID string) {
	contest := &leaderboardClient.Contest{
		ContestId: contestID,
	}
	for _, wwc := range leaderboardMap[contestID] {
		log.Println("Websocket", wwc)
		wwc.sendChannel <- contest
	}
}

func onSendContestID(wwc *websocketWithChannels) {
	for {
		messageType, contestID, err := wwc.websocket.ReadMessage()
		if err != nil {
			log.Println(err)
			log.Println("Error", err.Error())
			removeAndCloseWWC(wwc)
			return
		}
		if messageType == 1 {
			fmt.Println("onSendContestId read : ", string(contestID))
			contest := &leaderboardClient.Contest{
				ContestId: string(contestID),
			}
			wwc.sendChannel <- contest
		}
	}
}

func onReceiveLeaderboard(wwc *websocketWithChannels) {
	for {
		participants := <-wwc.receiveChannel
		if participants != nil {
			participantsJson, err := json.Marshal(participants)
			if err != nil {
				log.Println("Error while marshalling", err.Error())
				continue
			}
			log.Println("Writing to websocket")
			if err := wwc.websocket.WriteMessage(1, []byte(participantsJson)); err != nil {
				log.Println("onReceiveLeaderboard Error", err.Error())
				removeAndCloseWWC(wwc)
				return
			}
		}
	}

}

func removeAndCloseWWC(wwc *websocketWithChannels) {
	log.Println("Attempting closure of websocket")
	for k, wwcFromMap := range leaderboardMap[wwc.contestID] {
		log.Println(k)
		if wwc.websocket == wwcFromMap.websocket {
			log.Println("Closing websocket")
			err := wwc.websocket.Close()
			if err != nil {
				log.Println("Closing websocket err ", err.Error())
			}
			leaderboardMap[wwc.contestID] = append(leaderboardMap[wwc.contestID][:k],
				leaderboardMap[wwc.contestID][k+1:]...)
			log.Println("Closed websocket")
			break
		}
	}
	close(wwc.receiveChannel)
	close(wwc.sendChannel)
}

func handleConnection(wwc *websocketWithChannels) {
	// Call internal function to fetch leaderboard from leaderboard service
	// create stream
	stream, err := client.GetLeaderboard(context.Background())
	ctx := stream.Context()
	if err != nil {
		log.Printf("open stream error %v", err)
	}
	done := make(chan bool)
	go func() {
		for {
			contest := <-wwc.sendChannel
			log.Println("Send to leaderboard_client", contest)
			if contest != nil {
				if err := stream.Send(contest); err != nil {
					log.Printf("can not send %v", err)
				}
			} else {
				log.Println("contest is nil?")
				return
			}
		}
	}()
	// Receive leaderboard
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				close(done)
				log.Println("Closing leaderboard stream")
				return
			}
			if resp == nil {
				log.Println("Response nil")
				return
			}
			if err != nil {
				log.Printf("can not receive %v", err)
			}
			log.Println("Received from leaderboard_client", resp.String())
			wwc.receiveChannel <- resp
		}

	}()
	// Closes the bi-directional stream
	go func() {
		<-ctx.Done()
		if err := ctx.Err(); err != nil {
			log.Println(err)
			close(wwc.sendChannel)
			close(done)
			log.Println("channel done and sendChannel closed (stream closed)")
		}
	}()
	<-done
	log.Println("connection will be closed")
}
