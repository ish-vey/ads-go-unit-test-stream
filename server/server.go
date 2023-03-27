package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	pb "unit_test_stream/unit_test_stream"

	"google.golang.org/grpc"
)

type videoStreamRequest struct {
	fileName string
	data     []byte
	complete bool
}

type videoStreamReply struct {
	success bool
}

type unitTestStream struct {
	request  videoStreamRequest
	response videoStreamReply
}

func (s *unitTestStream) TestStream(stream pb.UnitTestStream_TestStreamServer) error {
	var raw_data []byte
	var fileName string
	complete := false

	for {
		data, err := stream.Recv()
		if err == io.EOF {
			fmt.Printf(fileName, string(raw_data), complete)
			return stream.SendAndClose(&pb.VideoStreamReply{
				Success: true,
			})
		}

		if err != nil {
			return err
		}

		raw_data = append(raw_data, data.Data...)

		fileName = data.FileName
		complete = data.Complete

	}
}

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	server := grpc.NewServer()
	fmt.Printf("Starting up server at %d", *port)
	server.Serve(lis)

}
