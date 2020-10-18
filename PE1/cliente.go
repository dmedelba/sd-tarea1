
// Package main implements a client for Greeter service.
package main

import (
	"os"
	"fmt"
	"bufio"
	//"context"
	"log"
	//"time"
	"encoding/csv"
	"google.golang.org/grpc"
	//pb "github.com/dmedelba/sd-tarea1/PE1/protos"
)

const (
	puerto     = "dist70:6969"
	idPaquete = "1";
    tipo = "pyme";
    nombre = "caca";
    valor = "1400";
    origen = "tienda-a";
    destino = "tienda-b";
)
//conn *grpc.ClientConn, parametro
func enviarPedidoPyme(id_pedido int)(int){
	//Se lee csv Pyme
	csvfile, err := os.Open("./archivos/pymes.csv")
	if err != nil {
		log.Fatalln("No se pudo leer el archivo", err)
	}
	//r := csv.NewReader(csvfile)
	r := csv.NewReader(bufio.NewReader(csvfile))
	for (i:=0; true; i++){
		linea, err := r.Read()
		if (i==id_pedido){
			log.Printf(linea)
		}
		if err == io.EOF{
			break
		}
		
	}
	
	log.Printf(linea)
	return 1
	//c := pb.NewProtosClient(conn)
}
func main() {

	var tipo_cliente string
	var accion string
	var enviado int
	id_pedido := 0 

	log.Printf("Estableciendo conexión con logistica...")

	//Establecemos conexión con logisitica dist70:6969
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(puerto, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo establecer la conexión. ERROR: %v", err)
	}
	defer conn.Close()

	//Interfaz inicial
	log.Printf("[Cliente] Seleccione el tipo de cliente que corresponde: ")
	log.Printf("1. Pyme")
	log.Printf("2. Retail")
	fmt.Scanln(&tipo_cliente)

	log.Printf("[Cliente]¿Qué desea hacer ?")
	log.Printf("1. Enviar un pedido")
	log.Printf("2. Consultar estado de un pedido")
	fmt.Scanln(&accion)
	
	switch tipo_cliente {
	case "1":
		//Cliente pyme
		if(accion == "1"){
			//leer csv pyme
			enviado = enviarPedidoPyme(id_pedido)
			//enviar pedido
		}else if (accion == "2"){
			//consultar estado de un pedido
		}
	case "2":
		//Cliente Retail
		if(accion == "1"){
			//enviar pedido retail
		}else if (accion == "2"){
			//consultar estado de un pedido
		}
	}
	/*
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SolicitarPedidoPyme(ctx, &pb.SolicitudPedidoPyme{IdPaquete:idPaquete,Tipo: tipo, Nombre: nombre, Valor: valor, Origen:origen, Destino:destino})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf(r)
	*/
	log.Printf("Conexión basica")
}
