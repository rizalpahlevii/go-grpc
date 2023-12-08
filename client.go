package main

import (
	"github.com/rizalpahlevii/go-grpc/chat"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not connect to gRPC server: %v", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "Hello From the Client!",
	}

	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from Server: %s", response.Body)
}
