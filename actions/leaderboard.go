package actions

import (
	"context"
	"io"
	"log"

	leaderboardClient "github.com/cpjudge/cpjudge_webserver/proto/leaderboard"
	"github.com/gobuffalo/buffalo"
	"google.golang.org/grpc"
)

// LeaderboardHandler : fetch leaderboard for contest
func LeaderboardHandler(c buffalo.Context) error {
	contestID := c.Request().URL.Query().Get("contest_id")
	if contestID != "" {
		contest := &leaderboardClient.Contest{
			ContestId: contestID,
		}
		handleConnection(leaderboardClient.GetLeaderboard(contest), contest, c)

	}
	return c.Render(400, r.JSON(map[string]interface{}{
		"message": "contest_id parameter missing",
	}))
}
func handleConnection(conn *grpc.ClientConn, contest *leaderboardClient.Contest,
	c buffalo.Context) {
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
		log.Println("Send", contest)
		if contest != nil {
			if err := stream.Send(contest); err != nil {
				log.Printf("can not send %v", err)
			}
		}
	}()
	closeReceive := make(chan bool)
	// Receive leaderboard
	go func() {
		select {
		case <-closeReceive:
			log.Println("Close receive")
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
		log.Println("closed end due to buffalo")
	}()
	// Closes the bi-directional stream
	go func() {
		<-ctx.Done()
		if err := ctx.Err(); err != nil {
			log.Println(err)
		}
		close(done)
		log.Println("closed end")
	}()
	<-done
	log.Println("closed")
}
