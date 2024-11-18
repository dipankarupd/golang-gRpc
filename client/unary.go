package main

import (
	"context"
	"log"

	pb "github.com/dipankarupd/grpc-demo/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx := context.Background()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("%s", res.Message)
}
