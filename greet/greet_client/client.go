package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	greetpb "github.com/GhvstCode/Grpc-course/greet"
)

func main() {
	fmt.Println("Hello, I'm a Client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf(" Could not connect: %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	//fmt.Printf("Created Client : %f", c)
	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient){
	fmt.Println("Hello, I'm a Unary")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
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