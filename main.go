package main

import (
	"fmt"
	"net"

	"github.com/PrakharSrivastav/artist-service-grpc/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func main() {

	// register rpc server
	rpc := rpc.New(nil)
	server := NewServer()
	server.Add(rpc)
	server.Start()
}

type Server struct {
	Health  *health.Server
	Server  *grpc.Server
	options []grpc.ServerOption
}

type Handler interface {
	Register(server *grpc.Server)
}

func NewServer() *Server {
	s := Server{}
	s.Health = health.NewServer()

	s.Server = grpc.NewServer(s.options...)
	grpc_health_v1.RegisterHealthServer(s.Server, s.Health)
	reflection.Register(s.Server)

	return &s
}

func (s *Server) getPort() string {
	return "6565"
}

func (s *Server) Start() {
	port := s.getPort()
	addr := fmt.Sprintf(":%s", port)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("Failed to start gRPC server", err.Error())
	}

	fmt.Println("Starting gRPC server on port %s", port)
	s.Server.Serve(lis)
}

func (s *Server) Add(handler Handler) {
	handler.Register(s.Server)
}
