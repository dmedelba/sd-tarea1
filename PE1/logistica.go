// Package main implements a server for Greeter service.
package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"github.com/dmedelba/sd-tarea1/PE1/protos"
)

const (
	port = ":6969"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	protos.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) GetPedido(ctx context.Context, in *protos.solicitudPedido) (*protos.respuestaPedido, error) {
	log.Printf("Received: %v", in.GetNombre())
	return &protos.respuestaPedido{Message: "Hello " + in.GetNombre()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	protos.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
