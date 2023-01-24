package server

import (
	"log"
	"net"

	"github.com/grpcdemo/chat/chat"
	"google.golang.org/grpc"
)

func StartGRPCServer() {
	lis, err := net.Listen("tcp", ":7000")
	if err != nil {
		log.Fatalf("error starting TCP server: %v", err)
	}

	grpcServer := grpc.NewServer()
	chatServer := chat.Server{}
	chat.RegisterChatServiceServer(grpcServer, chatServer)
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("error starting GRPC server: %v", err)
	}
}
