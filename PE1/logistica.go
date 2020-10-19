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
/*
type colas struct { 
	var numeros1 []i
}
*/

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


//Registros de paquetes en logistica, solicitado en enunciado.
//hacer una lista con estos archivos
func guardarPaquetesLogisticaPY(in *pb.SolicitudPedidoPyme, codigoSeguimiento string){
	file,err:=os.Create("./logistica_files/pyme/"+codigoSeguimiento+".csv")
	if err!=nil{
		log.Println(err)
	}
	now := time.Now().String()
	var paquete = [][]string{
		{now, in.IdPaquete, in.Tipo, in.Nombre, in.Valor, in.Origen, in.Destino, codigoSeguimiento},
	}
	w:=csv.NewWriter(file)
	w.WriteAll(paquete)
	file.Close()

	//guardamos el paquete en nuestro archivo general
	archivo, error := os.OpenFile("./logistica_files/pyme/pedidosGeneral.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if error != nil {
		log.Fatal(error)
	}
	defer archivo.Close()

	var paqueteG [][]string
	intentos := "0"
	estado := "En Bodega"
	paqueteG = append(paqueteG, []string{in.IdPaquete, codigoSeguimiento, in.Tipo, in.Valor, intentos, estado})
	ww := csv.NewWriter(archivo)
	ww.WriteAll(data)
	archivo.Close()
}

func guardarPaquetesLogisticaRT(in *pb.SolicitudPedidoRetail, codigoSeguimiento string){
	file,err:=os.Create("./logistica_files/retail/"+codigoSeguimiento+".csv")
	if err!=nil{
		log.Println(err)
	}
	now := time.Now().String()
	tipo := "retail"
	var paquete = [][]string{
		{now, in.IdPaquete, tipo, in.Nombre, in.Valor, in.Origen, in.Destino, codigoSeguimiento},
	}
	w:=csv.NewWriter(file)
	w.WriteAll(paquete)
	file.Close()

	//guardamos el paquete en nuestro archivo general
	archivo, error := os.OpenFile("./logistica_files/pyme/pedidosGeneral.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if error != nil {
		log.Fatal(error)
	}
	defer archivo.Close()

	var paqueteG [][]string
	intentos := "0"
	estado := "En Bodega"
	paqueteG = append(paqueteG, []string{in.IdPaquete, codigoSeguimiento, tipo, in.Valor, intentos, estado})
	ww := csv.NewWriter(archivo)
	ww.WriteAll(data)
	archivo.Close()

}

//Guardamos los pedidos en un archivo general para obtener su estado e ir actualizando

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
