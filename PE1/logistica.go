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
	port = ":7070"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedProtosServer
}
func crearCodigoSeguimiento(tipo_cliente string, in *SolicitudPedidoPyme)(string){
	if (tipo_cliente == "pyme"){
		codigo := "PY" + in.IdPaquete
	}else if (tipo_cliente == "retail"){
		codigo := "RT" + in.IdPaquete 
	}	
	return codigo
}
//recibir pedidoPyme
func (s *server) SolicitarPedidoPyme(ctx context.Context, in *pb.SolicitudPedidoPyme) (*pb.RespuestaPedido, error) {
	log.Printf("[Servidor] Pedido recibido: %+v", in)

	//Registrar el pedido y crear codigo de seguimiento
	codigo := crearCodigoSeguimiento("pyme", in)
	return &pb.RespuestaPedido{CodigoSeguimiento: codigo}, nil
}
//recibir pedidoRetail
func (s *server) SolicitarPedidoRetail(ctx context.Context, in *pb.SolicitudPedidoRetail) (*pb.RespuestaPedido, error) {
	log.Printf("[Servidor] Pedido recibido: %+v", in)
	//Registrar el pedido y crear codigo de seguimiento
	codigo := crearCodigoSeguimiento("retail", in)
	return &pb.RespuestaPedido{CodigoSeguimiento: codigo}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Servidor falla al escuchar. ERROR: %v", err)
	}
	log.Printf("[Servidor] Esperando comunicación.")
	s := grpc.NewServer()
	pb.RegisterProtosServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	
}
