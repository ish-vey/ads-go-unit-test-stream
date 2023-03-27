package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"
	pb "unit_test_stream/unit_test_stream"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func runTestStream(client pb.UnitTestStreamClient) {
	data_sample := pb.VideoStreamRequest{
		FileName: "Dummy File",
		Data:     []byte("Here is some dummy data!"),
		Complete: true,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.TestStream(ctx)

	if err != nil {
		log.Fatalf("client.TestStream failed: %v", err)
	}

	stream.Send(&data_sample)

	reply, err := stream.CloseAndRecv()
	fmt.Print(reply.Success)

}

var (
	serverAddr = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
)

func main() {
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(*serverAddr, opts)

	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewUnitTestStreamClient(conn)

	runTestStream(client)
}
