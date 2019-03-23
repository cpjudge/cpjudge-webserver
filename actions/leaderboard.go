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
	go func() {
		for {
			log.Println("Send")
			contest := <-sendChannel
			log.Println("Send", contest)
			if contest != nil {
				if err := stream.Send(contest); err != nil {
					log.Printf("can not send %v", err)
				}
			} else {
				return
			}
		}
	}()
	sendChannel <- contest
	// Receive leaderboard
	go func() {
		for {
			log.Println("Receive")
			resp, err := stream.Recv()
			if err == io.EOF {
				close(done)
				return
			}
			if resp == nil {
				return
			}
			if err != nil {
				log.Printf("can not receive %v", err)
			}
			log.Println("Received", resp.String())
		}

	}()
	// If request is cancelled by client(end user) this function ensures the
	// grpc bi-directional stream is closed properly.
	go func() {
		<-c.Done()
		sendChannel <- nil
		done <- true
		if err := c.Err(); err != nil {
			log.Println("Buffalo", err)
		}
		log.Println("close call(reason: buffalo)")
	}()
	// Closes the bi-directional stream
	go func() {
		<-ctx.Done()
		if err := ctx.Err(); err != nil {
			log.Println(err)
			for i, v := range leaderboardMap[contest.ContestId] {
				if v == sendChannel {
					log.Println("Removing channel", v)
					leaderboardMap[contest.ContestId] = append(
						leaderboardMap[contest.ContestId][:i],
						leaderboardMap[contest.ContestId][i+1:]...)
				}
			}
			close(sendChannel)
			close(done)
			log.Println("channel done and sendChannel closed (stream closed)")
		}
	}()
	<-done
	log.Println("connection will be closed")
}
