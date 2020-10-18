// Package main implements a server for Greeter service.
package main

import (
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	pb "github.com/dmedelba/sd-tarea1/PE1/protos"
)

const (
	port = ":6969"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

/*
// SayHello implements helloworld.GreeterServer
func (s *server) GetPedido(ctx context.Context, in *pb.solicitudPedido) (*pb.respuestaPedido, error) {
	log.Printf("Received: %v", in.GetNombre())
	return 
}
*/
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Servidor falla al escuchar. ERROR: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProtosServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Printf("Servidor OK")
}
