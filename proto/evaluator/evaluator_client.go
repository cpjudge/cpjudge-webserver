package evaluator

import (
	"context"
	"flag"
	"log"
	"time"

	spb "github.com/cpjudge/cpjudge_webserver/proto/submission"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containning the CA root cert file")
	serverAddr         = flag.String("server_addr", "172.17.0.1:12000", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")
)

// Internal function to call the evaluator service and submit the code.
func evaluateCode(client EvaluatorClient, submission *spb.Submission) *CodeStatus {
	log.Printf("Evaluating code with submission_id: %s", submission.GetSubmissionId())
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	codeStatus, err := client.EvaluateCode(ctx, submission)
	if err != nil {
		log.Fatalf("%v.EvaluateCode(_) = _, %v: ", client, err)
	}
	log.Println("After evaluation: ", codeStatus)
	return codeStatus
}

// EvaluateCode : Spawns a connection to evaluator service and calls evaluateCode()
func EvaluateCode(submission *spb.Submission) *CodeStatus {
	// Establish a connection with the evaluator service.
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
	defer conn.Close()
	client := NewEvaluatorClient(conn)
	// Call internal function to submit code to evaluator service
	codeStatus := evaluateCode(client, submission)
	return codeStatus
}
