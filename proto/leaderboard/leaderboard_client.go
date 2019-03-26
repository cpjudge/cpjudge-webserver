package leaderboard

import (
	"flag"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"
)

var (
	tls                = flag.Bool("tls_leaderboard", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file_leaderboard", "", "The file containning the CA root cert file")
	serverAddr         = flag.String("server_addr_leaderboard", "172.17.0.1:11000", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override_leaderboard", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")
)

// Internal function to call the leaderboard service and fetch participants.
// func getLeaderboard(client LeaderboardClient, contest *Contest) *Participants {
// 	log.Printf("Leaderboard for  contest_id: %s", contest.GetContestId())
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	particpants, err := client.GetLeaderboard(ctx, contest)
// 	if err != nil {
// 		log.Fatalf("%v.GetLeaderboard(_) = _, %v: ", client, err)
// 	}
// 	return particpants
// }

// GetLeaderboard : Spawns a connection to evaluator service and calls Evaluate()
func GetLeaderboard() *grpc.ClientConn {
	// Establish a connection with the leaderboard service.
	flag.Parse()
	var opts []grpc.DialOption
	if *tls {
		if *caFile == "" {
			*caFile = testdata.Path("ca.pem")
		}
		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	return conn
}
