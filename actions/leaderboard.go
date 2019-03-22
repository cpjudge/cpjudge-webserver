package actions

import (
	"context"
	"io"
	"log"

	leaderboardClient "github.com/cpjudge/cpjudge_webserver/proto/leaderboard"
	"github.com/gobuffalo/buffalo"
	"google.golang.org/grpc"
)

var leaderboardMap map[string]([]chan *leaderboardClient.Contest)

// LeaderboardHandler : fetch leaderboard for contest
func LeaderboardHandler(c buffalo.Context) error {
	if leaderboardMap == nil {
		leaderboardMap = make(map[string]([]chan *leaderboardClient.Contest))
	}
	contestID := c.Request().URL.Query().Get("contest_id")
	if contestID != "" {
		contest := &leaderboardClient.Contest{
			ContestId: contestID,
		}
		if leaderboardMap[contest.ContestId] == nil {
			leaderboardMap[contest.ContestId] = make([]chan *leaderboardClient.Contest, 0)
		}
		channel := make(chan *leaderboardClient.Contest)
		leaderboardMap[contest.ContestId] = append(leaderboardMap[contest.ContestId], channel)
		handleConnection(leaderboardClient.GetLeaderboard(contest), contest, c, channel)

	}
	return c.Render(400, r.JSON(map[string]interface{}{
		"message": "contest_id parameter missing",
	}))
}

// TriggerLeaderboards : Trigger all leaderboards to update their status
func TriggerLeaderboards(contestID string) {
	contest := &leaderboardClient.Contest{
		ContestId: contestID,
	}
	for _, channel := range leaderboardMap[contestID] {
		log.Println("Channel", channel)
		channel <- contest
	}
}

func handleConnection(conn *grpc.ClientConn, contest *leaderboardClient.Contest,
	c buffalo.Context,
	sendChannel chan *leaderboardClient.Contest) {
	client := leaderboardClient.NewLeaderboardClient(conn)
	// Call internal function to fetch leaderboard from leaderboard service
	// create stream
	defer conn.Close()
	stream, err := client.GetLeaderboard(context.Background())
	ctx := stream.Context()
	if err != nil {
		log.Printf("open stream error %v", err)
	}
	done := make(chan bool)
	closeSend := make(chan bool)
	go func() {
		select {
		case <-closeSend:
			log.Println("Close receive stream")
			return
		default:
			for {
				contest := <-sendChannel
				log.Println("Send", contest)
				if contest != nil {
					if err := stream.Send(contest); err != nil {
						log.Printf("can not send %v", err)
					}
				}
			}
		}
	}()
	sendChannel <- contest
	closeReceive := make(chan bool)
	// Receive leaderboard
	go func() {
		select {
		case <-closeReceive:
			log.Println("Close receive stream")
			return
		default:
			for {
				resp, err := stream.Recv()
				if err == io.EOF {
					close(done)
					return
				}
				if err != nil {
					log.Printf("can not receive %v", err)
				}
				log.Println("Received", resp.String())
			}
		}

	}()
	// If request is cancelled by client(end user) this function ensures the
	// grpc bi-directional stream is closed properly.
	go func() {
		<-c.Done()
		if err := c.Err(); err != nil {
			log.Println(err)
		}
		closeReceive <- true
		closeSend <- true
		log.Println("close call(reason: buffalo)")
	}()
	// Closes the bi-directional stream
	go func() {
		<-ctx.Done()
		if err := ctx.Err(); err != nil {
			log.Println(err)
		}
		close(done)
		close(sendChannel)
		close(closeReceive)
		close(closeSend)
		log.Println("stream closed")
		log.Println("channel done closeReceive and sendChannel closed")
	}()
	<-done
	log.Println("connection closed")
}
