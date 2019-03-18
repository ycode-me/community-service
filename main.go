package main

import (
	"context"
	"log"
	"net"

	pb "github.com/ycode-me/ycodeapis/community/v1"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement pb Service.
type server struct{}

func (s *server) CreateMoment(ctx context.Context,
	in *pb.CreateMomentRequest) (*pb.CreateMomentResponse, error) {
	return &pb.CreateMomentResponse{ResourceName: "foo"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to Listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCommunityServiceServer(grpcServer, &server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve: %v", err)
	}
}
