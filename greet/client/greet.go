package main

import (
	"context"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
)

func doGreet(c pb.SumServiceClient) {
	log.Println("doSum was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNum:  5,
		SecondNum: 20,
	})

	if err != nil {
		log.Fatalf("Could not Sum: %v\n", err)
	}

	log.Printf("Suming: %d\n", res.Result)
}
