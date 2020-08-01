package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/GhvstCode/Grpc-course/calculator/calculatorpb"
)

type Server struct{}

func (*Server) Sum(ctx context.Context, in *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Received Sum Rpc : %v", in)
	firstNumber := in.FirstNumber
	secondNumber := in.LastNumber

	sum := firstNumber + secondNumber
	res := &calculatorpb.SumResponse{
		SumResult: sum,
	}

	return res, nil
}

func main() {
	fmt.Println("Calculator Server Started")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err  != nil {
		log.Fatalf("Unable to listen %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorResponseServer(s, &Server{})


	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}