package client

import (
	"context"
	"log"

	"github.com/grpcdemo/chat/chat"
	"google.golang.org/grpc"
)

func StartGRPCClient() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":7000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	response, err := c.SayHello(context.Background(), &chat.Message{Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)
}
