package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"

	greetpb2 "github.com/GhvstCode/Grpc-course/greet/greetpb"
)

func main() {
	fmt.Println("Hello, I'm a Client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf(" Could not connect: %v", err)
	}
	defer cc.Close()

	c := greetpb2.NewGreetServiceClient(cc)
	//fmt.Printf("Created Client : %f", c)
	//doUnary(c)
	doServerStreaming(c)
}
func doServerStreaming(c greetpb2.GreetServiceClient){
	req := &greetpb2.GreetManyTimesRequest{
		Greeting: &greetpb2.Greeting{
			FirstName: "Ghvst",
			LastName: "Code",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error caught while calling ServerStreaming Rpc : %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF{
			break
		}

		if err != nil {
			log.Fatalf("Error caught while calling looping ServerStreaming Rpc : %v", err)
		}

		log.Printf(" ServerStreaming Rpc Message: %v", msg.GetResult())
	}

}

func doUnary(c greetpb2.GreetServiceClient){
	fmt.Println("Hello, I'm a Unary")
	req := &greetpb2.GreetRequest{
		Greeting: &greetpb2.Greeting{
			FirstName: "Ghvst",
			LastName: "Code",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error caught while calling Greet Rpc : %v", err)
	}

	log.Printf("Response from Greet : %v", res.Result)
}