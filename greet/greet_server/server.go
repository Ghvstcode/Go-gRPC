package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	greetpb "github.com/GhvstCode/Grpc-course/greet"
)

type Server struct{}

func (*Server) Greet(ctx context.Context, in *greetpb.GreetRequest) (*greetpb.GreetResponse, error){
	fmt.Printf("Greet Function was Invoked with %v", in)
	firstName := in.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := greetpb.GreetResponse{
		Result: result,
	}
	return &res, nil
}

func main() {
	fmt.Println("Hello world")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err  != nil {
		log.Fatalf("Unable to listen %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &Server{})


	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}