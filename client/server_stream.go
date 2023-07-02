package main

import (
	"context"
	"io"
	"log"

	pb "github.com/TulioMeran/go_example_grpc/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Println("Streaming started")

	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while streaming %v", err)
		}

		log.Println(message)
	}
	log.Println("Streaming finished")

}
