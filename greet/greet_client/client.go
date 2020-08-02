package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

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
	//doServerStreaming(c)
	//doClientStreaming(c)
	doBidiStreaming(c)
}
func doBidiStreaming (c greetpb2.GreetServiceClient) {
	request := []*greetpb2.GreetEveryoneRequest{
		&greetpb2.GreetEveryoneRequest{
			Greeting: &greetpb2.Greeting{
				FirstName: "Ab",
			},
		},
		&greetpb2.GreetEveryoneRequest{
			Greeting: &greetpb2.Greeting{
				FirstName: "Ac",
			},
		},
		&greetpb2.GreetEveryoneRequest{
			Greeting: &greetpb2.Greeting{
				FirstName: "Ad",
			},
		},
		&greetpb2.GreetEveryoneRequest{
			Greeting: &greetpb2.Greeting{
				FirstName: "Ae",
			},
		},
	}

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while Creating stream:  %v", err)
	}
	//waitc := make(chan struct{})
	var wg sync.WaitGroup

	wg.Add(2)
	go func () {
		for _, req := range request {
			fmt.Printf("Sending Message: %v \n", req)
			_ = stream.Send(req)
			time.Sleep(2 *time.Second)
		}
		wg.Done()
		if err := stream.CloseSend(); err != nil {
			fmt.Printf("An error occured while closing send: %v \n", err)
		}
	}()
	go func () {
		for {
			res, err2 := stream.Recv()
			if err2 == io.EOF{
				fmt.Printf("An error occured: %v \n", err2)
				break
			}
			if err2 != nil {
				fmt.Printf("An error occured: %v", err)
				break
			}
			fmt.Printf("Receiving Message: %v \n", res.GetResult())

		}
		//close(waitc)
		wg.Done()
	}()
	//<-waitc
	wg.Wait()
}
func doClientStreaming (c greetpb2.GreetServiceClient) {
	request := []*greetpb2.LongGreetRequest{
		&greetpb2.LongGreetRequest{
			Greeting: &greetpb2.Greeting{
				FirstName: "Ab",
			},
		},
		&greetpb2.LongGreetRequest{
			Greeting: &greetpb2.Greeting{
				FirstName: "Ac",
			},
		},
		&greetpb2.LongGreetRequest{
			Greeting: &greetpb2.Greeting{
				FirstName: "Ad",
			},
		},
		&greetpb2.LongGreetRequest{
			Greeting: &greetpb2.Greeting{
				FirstName: "Ae",
			},
		},
	}
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling long greet: %v", err)
	}
	for _, req := range request{
		fmt.Printf("Sending Request : %v", req)
		_ = stream.Send(req)
		time.Sleep(3 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving : %v", err)
	}
	fmt.Printf("Receiving Request : %v", res)
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