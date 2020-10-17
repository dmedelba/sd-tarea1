
// Package main implements a client for Greeter service.
package main

import (
	"context"
	"log"
	"time"
	"google.golang.org/grpc"
	"github.com/dmedelVVJd2aj3ba/sd-tarea1/PE1/protos"
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
	c := protos.NewGreeterClient(conn)

	// Contact the server and print out its response.

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetPedido(ctx, &protos.solicitudPedido{idPaquete: idPaquete,tipo: tipo, nombre: nombre, valor: valor, origen:origen, destino:destino})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
