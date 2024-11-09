package main

import (
	"chat-service/pkg/chat_v1"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
)

const grpcPort = 50051

type server struct {
	*chat_v1.UnimplementedChatV1Server
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	chat_v1.RegisterChatV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) Create(context.Context, *chat_v1.CreateRequest) (*chat_v1.CreateResponse, error) {
	return &chat_v1.CreateResponse{Id: 50505050050505}, nil
}
func (s *server) Delete(context.Context, *chat_v1.DeleteRequest) (*emptypb.Empty, error) {
	return nil, nil
}
func (s *server) SendMessage(context.Context, *chat_v1.SendMessageRequest) (*emptypb.Empty, error) {
	return nil, nil
}
