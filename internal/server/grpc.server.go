package server

import (
	"log"
	"net"

	"go.onRunCRM/internal/handler"
	pb "go.onRunCRM/pkg/api"
	"google.golang.org/grpc"
)

const PORT = ":50051"

func RunGRPCServer() {
	lis, err := net.Listen("tcp", PORT)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	tenantServer := &handler.TenantHandlerServer{}

	pb.RegisterTenantsServiceServer(grpcServer, tenantServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Printf("Success run")
}
