package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/tesarwijaya/navillera/app/procedures/entity"
	"github.com/tesarwijaya/navillera/config"

	pb "github.com/tesarwijaya/navillera/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// Server struct for instantiate serve function gRPC.
type Server struct {
}

// GetEntity ...
func (v *Server) GetEntity(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	return entity.Get(request)
}

// unaryInterceptor to use as authentication example
func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, _ := metadata.FromIncomingContext(ctx)

	fmt.Println("metadata: ", md)
	fmt.Println("req: ", req)

	// Auth logic here!

	m, err := handler(ctx, req)
	if err != nil {
		fmt.Println("RPC failed with error ", err)
	}
	return m, err
}

// Init ...
func Init() {
	listening, err := net.Listen("tcp", fmt.Sprintf("%s:%s", config.C.GrpcHost, config.C.GrpcPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(unaryInterceptor))
	pb.RegisterCallServer(grpcServer, &Server{})
	grpcServer.Serve(listening)
}
