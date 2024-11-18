package main

import (
	"context"
	"io"
	"log"

	pb "github.com/dipankarupd/grpc-demo/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {

	log.Printf("Streaming started: \n")

	streams, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("could not send names %v", err)
	}

	for {
		msg, err := streams.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while streaming %v", err)
		}

		log.Println(msg)
	}

	log.Printf("Streaming finished")

}
