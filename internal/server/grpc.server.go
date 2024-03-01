package server

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

const PORT = ":50051"

func RunGRPCServer() {
	lis, err := net.Listen("tcp", PORT)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
