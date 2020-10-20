// Package main implements a server for Greeter service.
package main

import (
	"bufio"
	"container/list"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"

	pb "github.com/dmedelba/sd-tarea1/PE1/protos"
	"google.golang.org/grpc"
)

const (
	portCliente = ":6070"
	portCamion  = ":6071"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedProtosServer
}

// paquete
type Paquete struct {
	IdPaquete         string
	CodigoSeguimiento string
	Tipo              string
	Valor             string
	Origen            string
	Destino           string
	Intentos          string
	Estado            string
}

//Colas de los pedidos en logistica
var cPrioritario = list.New()
var cNormal = list.New()
var cRetail = list.New()

func conexionCamiones() {
	lis, err := net.Listen("tcp", portCamion)
	if err != nil {
		log.Fatalf("Servidor falla al escuchar camiones. ERROR: %v", err)
	}
	log.Printf("[Servidor] Esperando comunicación camiones.")
	s := grpc.NewServer()
	pb.RegisterProtosServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

//Crear codigos de seguimiento unicos.
func crearCodigoSeguimientoPyme(in *pb.SolicitudPedidoPyme) string {
	var codigo string
	codigo = "PY" + in.IdPaquete
	return codigo
}
func crearCodigoSeguimientoRetail(in *pb.SolicitudPedidoRetail) string {
	var codigo string
	codigo = "RT" + in.IdPaquete

	return codigo
}

//Registros de paquetes en logistica, solicitado en enunciado.
//Guardamos los pedidos en un archivo general para obtener su estado e ir actualizando
//tambien se aprovecha de asignar el pedio a la cola correspondiente

func guardarPaquetesLogisticaPY(in *pb.SolicitudPedidoPyme, codigoSeguimiento string) {
	file, err := os.Create("./logistica_files/pyme/" + codigoSeguimiento + ".csv")
	if err != nil {
		log.Println(err)
	}
	now := time.Now().String()
	var paquete = [][]string{
		{now, in.IdPaquete, in.Tipo, in.Nombre, in.Valor, in.Origen, in.Destino, codigoSeguimiento},
	}
	w := csv.NewWriter(file)
	w.WriteAll(paquete)
	file.Close()

	//guardamos el paquete en nuestro archivo general
	archivo, error := os.OpenFile("./logistica_files/pedidosGeneral.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if error != nil {
		log.Fatal(error)
	}
	defer archivo.Close()

	var paqueteG [][]string
	intentos := "0"
	estado := "En Bodega"
	paqueteG = append(paqueteG, []string{in.IdPaquete, codigoSeguimiento, in.Tipo, in.Valor, intentos, estado})
	ww := csv.NewWriter(archivo)
	ww.WriteAll(paqueteG)
	archivo.Close()

	//agrego paquete a la cola
	pedido := Paquete{
		IdPaquete:         in.IdPaquete,
		CodigoSeguimiento: codigoSeguimiento,
		Tipo:              in.Tipo,
		Valor:             in.Valor,
		Origen:            in.Origen,
		Destino:           in.Destino,
		Intentos:          intentos,
		Estado:            estado,
	}

	if pedido.Tipo == "prioritario" {
		cPrioritario.PushBack(pedido)
	} else {
		cNormal.PushBack(pedido)
	}

}

func guardarPaquetesLogisticaRT(in *pb.SolicitudPedidoRetail, codigoSeguimiento string) {
	file, err := os.Create("./logistica_files/retail/" + codigoSeguimiento + ".csv")
	if err != nil {
		log.Println(err)
	}
	now := time.Now().String()
	tipo := "retail"
	var paquete = [][]string{
		{now, in.IdPaquete, tipo, in.Nombre, in.Valor, in.Origen, in.Destino, codigoSeguimiento},
	}
	w := csv.NewWriter(file)
	w.WriteAll(paquete)
	file.Close()

	//guardamos el paquete en nuestro archivo general
	archivo, error := os.OpenFile("./logistica_files/pedidosGeneral.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if error != nil {
		log.Fatal(error)
	}
	defer archivo.Close()

	var paqueteG [][]string
	intentos := "0"
	estado := "En Bodega"
	paqueteG = append(paqueteG, []string{in.IdPaquete, codigoSeguimiento, tipo, in.Valor, intentos, estado})
	ww := csv.NewWriter(archivo)
	ww.WriteAll(paqueteG)
	archivo.Close()

	//agrego paquete a la cola
	pedido := Paquete{
		IdPaquete:         in.IdPaquete,
		CodigoSeguimiento: codigoSeguimiento,
		Tipo:              tipo,
		Valor:             in.Valor,
		Origen:            in.Origen,
		Destino:           in.Destino,
		Intentos:          intentos,
		Estado:            estado,
	}
	cRetail.PushBack(pedido)
	log.Printf("PAQUETE AGREGADO A LA COLA, SE SUPONE")
	fmt.Println(cRetail.Front())

}

//solicitar el estado de un pedido desde el cliente
func (s *server) ObtenerCodigoSeguimiento(ctx context.Context, in *pb.SolicitudSeguimiento) (*pb.RespuestaSeguimiento, error) {
	log.Printf("[Servidor] Consultando el codigo: %d", in.CodigoSeguimiento)
	var estadoPedido string
	csvFile, _ := os.Open("./logistica_files/pedidosGeneral.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		if line[1] == in.CodigoSeguimiento {
			csvFile.Close()
			estadoPedido = line[5]
			break
		}
	}
	return &pb.RespuestaSeguimiento{EstadoPedido: estadoPedido}, nil
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

//enviamos paquete a el camion dependiendo de su tipo
func (s *server) SolicitarPaquete(ctx context.Context, in *pb.SolicitudPaquete) (*pb.RespuestaPaquete, error) {
	var paqueteEnviar Paquete
	var vacio string
	tipodeCamion := in.Tipo

	if tipodeCamion == "retail" {
		if cRetail.Front() != nil {
			primerElemento := cRetail.Front()
			paqueteEnviar = Paquete(primerElemento.Value.(Paquete))
			cRetail.Remove(primerElemento)

		} else if cPrioritario.Front() != nil {
			primerElemento := cPrioritario.Front()
			paqueteEnviar = Paquete(primerElemento.Value.(Paquete))
			cPrioritario.Remove(primerElemento)

		} else {
			log.Printf("[Servidor] No hay paquetes disponibles")
			return &pb.RespuestaPaquete{
				IdPaquete:         vacio,
				CodigoSeguimiento: vacio,
				Tipo:              vacio,
				Valor:             vacio,
				Origen:            vacio,
				Destino:           vacio,
			}, nil
		}
	} else {
		if cPrioritario.Front() != nil {
			primerElemento := cPrioritario.Front()
			paqueteEnviar = Paquete(primerElemento.Value.(Paquete))
			cPrioritario.Remove(primerElemento)

		} else if cNormal.Front() != nil {
			primerElemento := cNormal.Front()
			paqueteEnviar = Paquete(primerElemento.Value.(Paquete))
			cNormal.Remove(primerElemento)

		} else {
			log.Printf("[Servidor] No hay paquetes disponibles")
			return &pb.RespuestaPaquete{
				IdPaquete:         vacio,
				CodigoSeguimiento: vacio,
				Tipo:              vacio,
				Valor:             vacio,
				Origen:            vacio,
				Destino:           vacio,
			}, nil

		}
	}
	//cambiar estado del paquete a "En camino" en el csv.
	return &pb.RespuestaPaquete{
		IdPaquete:         paqueteEnviar.IdPaquete,
		CodigoSeguimiento: paqueteEnviar.CodigoSeguimiento,
		Tipo:              paqueteEnviar.Tipo,
		Valor:             paqueteEnviar.Valor,
		Origen:            paqueteEnviar.Origen,
		Destino:           paqueteEnviar.Destino,
	}, nil
}

//Recibimos el estado del paquete de los camiones
func (s *server) ObtenerEstado(ctx context.Context, in *pb.SolicitudEstado) (*pb.RespuestaEstado, error){
	//recibimos el estado y actualizamos el csv de los pedidos.

	log.Printf(in.IdPaquete)
	log.Printf(in.Intentos)
	log.Printf(in.Estado)
	log.Printf(in.Fecha)
	return &pb.RespuestaEstado{Confirmacion: "1"}, nil

}

func main() {
	go conexionCamiones()

	lis, err := net.Listen("tcp", portCliente)
	if err != nil {
		log.Fatalf("Servidor falla al escuchar cliente. ERROR: %v", err)
	}
	log.Printf("[Servidor] Esperando comunicación cliente.")
	s := grpc.NewServer()
	pb.RegisterProtosServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
