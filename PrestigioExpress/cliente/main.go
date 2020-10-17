
// Package main implements a client for Greeter service.
package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	pb "github.com/dmedelba/sd-tarea1/tree/master/PrestigioExpress/ordenCliente/grpc"
)

const (
	address     = "dist70:6969"
	idPaquete = "1";
    tipo = "pyme";
    nombre = "caca";
    valor = "1400";
    origen = "tienda-a";
    destino = "tienda-b";
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetPedido(ctx, &pb.HelloRequest{idPaquete: idPaquete,tipo: tipo, nombre: nombre, valor: valor, origen:origen, destino:destino})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
