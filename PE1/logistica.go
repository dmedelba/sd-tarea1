// Package main implements a server for Greeter service.
package main

import (
	"context"
	"log"
	"os"
	"encoding/csv"
	"time"
	"net"
	"google.golang.org/grpc"
	pb "github.com/dmedelba/sd-tarea1/PE1/protos"
)

const (
	port = ":7071"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedProtosServer
}

//Crear codigos de seguimiento unicos.
func crearCodigoSeguimientoPyme(in *pb.SolicitudPedidoPyme)(string){
	var codigo string
	codigo = "PY" + in.IdPaquete
	return codigo
}
func crearCodigoSeguimientoRetail(in *pb.SolicitudPedidoRetail)(string){
	var codigo string
	codigo = "RT" + in.IdPaquete
	
	return codigo
}
//Registros de paquetes en logistica.
func guardarPaquetesLogisticaPY(in *pb.SolicitudPedidoPyme, codigoSeguimiento string){
	file, err := os.OpenFile("./pedidos_logistica/pedidos.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	now := time.Now().String()
	var paquete [][]string
	paquete = append(paquete, []string{now, in.IdPaquete, in.Tipo, in.Nombre, in.Valor, in.Origen, in.Destino, codigoSeguimiento})
	w := csv.NewWriter(file)
	w.WriteAll(paquete)
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
	log.Printf("Agregado paquete al archivo paquetes.csv")
}

func guardarPaquetesLogisticaRT(in *pb.SolicitudPedidoRetail, codigoSeguimiento string){
	file, err := os.OpenFile("./pedidos_logistica/pedidos.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	now := time.Now().String()
	var paquete [][]string
	paquete = append(paquete, []string{now, in.IdPaquete, "retail", in.Nombre, in.Valor, in.Origen, in.Destino, codigoSeguimiento})
	w := csv.NewWriter(file)
	w.WriteAll(paquete)
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
	log.Printf("Agregado paquete al archivo paquetes.csv")
}

//recibir pedidoPyme
func (s *server) SolicitarPedidoPyme(ctx context.Context, in *pb.SolicitudPedidoPyme) (*pb.RespuestaPedido, error) {
	log.Printf("[Servidor] Pedido recibido: %+v", in)

	//Registrar el pedido y crear codigo de seguimiento
	codigo := crearCodigoSeguimientoPyme(in)
	guardarPaquetesLogisticaPY(in, codigo)
	return &pb.RespuestaPedido{CodigoSeguimiento: codigo}, nil
}
//recibir pedidoRetail
func (s *server) SolicitarPedidoRetail(ctx context.Context, in *pb.SolicitudPedidoRetail) (*pb.RespuestaPedido, error) {
	log.Printf("[Servidor] Pedido recibido: %+v", in)

	//Registrar el pedido y crear codigo de seguimiento
	codigo := crearCodigoSeguimientoRetail(in)
	guardarPaquetesLogisticaRT(in, codigo)
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
