package main

import (
	"io"
	"log"

	pb "github.com/dipankarupd/grpc-demo/proto"
)

func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error receiving %v", err)
		}
		log.Printf("Got stream with name %v", req.Name)

		resp := &pb.HelloResponse{
			Message: "Hello" + req.Name,
		}
		if err := stream.Send(resp); err != nil {
			return err
		}
	}
}
