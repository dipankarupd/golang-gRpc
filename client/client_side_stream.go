package main

import (
	"context"
	"log"
	"time"

	pb "github.com/dipankarupd/grpc-demo/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Client stream started. \n")

	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names %v", err)
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("error while sending %v", err)
		}

		log.Printf("sent request with name %s", name)
		time.Sleep(2 * time.Second)
	}

	resp, err := stream.CloseAndRecv()
	log.Printf("client streaming finished")
	if err != nil {
		log.Fatalf("error while receiving")
	}
	log.Printf("%v", resp.Messages)
}
