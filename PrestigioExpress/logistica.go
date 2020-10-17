// Package main implements a server for Greeter service.
package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "github.com/dmedelba/sd-tarea1/tree/master/PrestigioExpress/ordenCliente/grpc"
)

const (
	port = ":6969"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) GetPedido(ctx context.Context, in *pb.solicitudPedido) (*pb.respuestaPedido, error) {
	log.Printf("Received: %v", in.GetNombre())
	return &pb.respuestaPedido{Message: "Hello " + in.GetNombre()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
