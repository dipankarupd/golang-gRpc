package main

import (
	"io"
	"log"

	pb "github.com/dipankarupd/grpc-demo/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {

	var messages []string

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessagesList{
				Messages: messages,
			})
		}
		if err != nil {
			log.Fatalf("error %v", err)
		}
		log.Printf("Got request: %v", req.Name)
		messages = append(messages, req.Name)
	}
}
