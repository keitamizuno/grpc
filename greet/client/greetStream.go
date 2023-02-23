package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
)

func doGreetStream(c pb.GreetServiceClient) {
	log.Printf("doGreetStream was invoked")

	req := &pb.GreetRequest{
		FirstName: "Keita",
	}

	stream, err := c.GreetStream(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not GreatStream: %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}
		log.Printf("GreetStream: %d\n", msg.Result)
	}

}
