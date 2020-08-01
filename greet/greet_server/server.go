package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"

	greetpb2 "github.com/GhvstCode/Grpc-course/greet/greetpb"
)

type Server struct{}

func (*Server) Greet(ctx context.Context, in *greetpb2.GreetRequest) (*greetpb2.GreetResponse, error){
	fmt.Printf("Greet Function was Invoked with %v", in)
	firstName := in.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := greetpb2.GreetResponse{
		Result: result,
	}
	return &res, nil
}

func (*Server) GreetManyTimes(in *greetpb2.GreetManyTimesRequest, stream greetpb2.GreetService_GreetManyTimesServer) error {
	firstName := in.GetGreeting().GetFirstName();
	for i := 0; i <10; i++ {
		result := "Hello" + firstName + "number " + strconv.Itoa(i)
		res := &greetpb2.GreetManyTimesResponse{
			Result: result,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep( 2*time.Second)
	}
	return nil
}

func main() {
	fmt.Println("Hello world")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err  != nil {
		log.Fatalf("Unable to listen %v", err)
	}

	s := grpc.NewServer()
	greetpb2.RegisterGreetServiceServer(s, &Server{})


	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}