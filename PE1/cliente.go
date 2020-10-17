
// Package main implements a client for Greeter service.
package main

import (
	"context"
	"log"
	"time"
	"google.golang.org/grpc"
	"github.com/dmedelba/sd-tarea1/PE1/protos"
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
	
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := protos.NewGreeterClient(conn)

	// Contact the server and print out its response.
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetPedido(ctx, &ordenCliente.solicitudPedido{idPaquete: idPaquete,tipo: tipo, nombre: nombre, valor: valor, origen:origen, destino:destino})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Conexión basica")
}
