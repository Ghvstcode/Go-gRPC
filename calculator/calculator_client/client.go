package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	"github.com/GhvstCode/Grpc-course/calculator/calculatorpb"
)

func main() {
	fmt.Println("Hello, I'm a Client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf(" Could not connect: %v", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculatorResponseClient(cc)
	//fmt.Printf("Created Client : %f", c)
	doUnary(c)
}

func doUnary(c calculatorpb.CalculatorResponseClient){
	fmt.Println("Hello, I'm a Unary")
	req := &calculatorpb.SumRequest{
			FirstNumber: 3,
			LastNumber: 10,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error caught while calling Greet Rpc : %v", err)
	}

	log.Printf("Response from Calculator : %v", res.SumResult)
}