package main

import (
	"chat-service/internal/api/note"
	"chat-service/internal/config"
	noteRepository "chat-service/internal/repository/note"
	noteService "chat-service/internal/service/note"
	"chat-service/pkg/chat_v1"
	"chat-service/pkg/note_v1"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
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

	// Loading env files
	ctx := context.Background()

	err := config.Load(".env")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	grpcConfig, err := config.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	pgConfig, err := config.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %v", err)
	}

	lis, err := net.Listen("tcp", grpcConfig.GRPCAddress())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pool, err := pgxpool.Connect(
		ctx,
		pgConfig.DSN(),
	)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	noteRepo := noteRepository.NewRepository(pool)
	noteServ := noteService.NewService(noteRepo)
	s := grpc.NewServer()
	reflection.Register(s)
	note_v1.RegisterNoteV1Server(s, note.NewImplementation(noteServ))
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
