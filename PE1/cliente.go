
// Package main implements a client for Greeter service.
package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
	"strconv"
	"context"
	"log"
	"time"
	"encoding/csv"
	"google.golang.org/grpc"
	pb "github.com/dmedelba/sd-tarea1/PE1/protos"
)

const (
	cantidad_pedidos = 60;
	puerto     = "dist70:6969"
	idPaquete = "1";
    tipo = "pyme";
    nombre = "caca";
    valor = "1400";
    origen = "tienda-a";
    destino = "tienda-b";
)
func contarPedidos(nombre_archivo string)(int){
	cantidad_lineas := 0
	csvfile, err := os.Open(nombre_archivo)
		if err != nil {
			log.Fatalln("No se pudo leer el archivo", err)
		}
		r := csv.NewReader(bufio.NewReader(csvfile))
		for{
			line, err := r.Read()
			if err == io.EOF{
				break
			}else if err != nil && line==nil{ //line puesto porque o si no se queja
				log.Fatal(err)
				continue
			}
			cantidad_lineas++
		}
	return cantidad_lineas
}

func enviarPedido(conn *grpc.ClientConn, id_pedido int, tipo_cliente string)(int){
	//conexión con el servidor 
	c := pb.NewProtosClient(conn)
	if (tipo_cliente == "1"){
		//Se lee csv Pyme
		csvfile, err := os.Open("./archivos/pymes.csv")
		log.Printf("Pyme leido")
		if err != nil {
			log.Fatalln("No se pudo leer el archivo", err)
		}
			//r := csv.NewReader(csvfile)
		r := csv.NewReader(bufio.NewReader(csvfile))
		for i:=0; true; i++{
			line, err := r.Read()
			if (i == id_pedido){
				log.Printf(line[0])
				log.Printf(line[1])
				log.Printf(line[2])
				log.Printf(line[3])
				log.Printf(line[4])
				log.Printf(line[5])

				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				r, err := c.SolicitarPedidoPyme(ctx, &pb.SolicitudPedidoPyme{
					IdPaquete:line[0],
					Tipo: line[5], 
					Nombre: line[1], 
					Valor: line[2], 
					Origen:line[3], 
					Destino:line[4]})
				if err != nil {
					log.Fatalf("No se pudo enviar el pedido. ERROR: %v", err)
				}
				return 1
			}
			if err == io.EOF{
				break
			}else if err != nil{
				log.Fatal(err)
				continue
			}
			
		}
	}else{
		//Se lee csv Retail
		csvfile, err := os.Open("./archivos/retail.csv")
		if err != nil {
			log.Fatalln("No se pudo leer el archivo", err)
		}
		//r := csv.NewReader(csvfile)
		r := csv.NewReader(bufio.NewReader(csvfile))
		for i:=0; true; i++{
			line, err := r.Read()
			if (i == id_pedido){
				log.Printf(line[0])
				return 1
			}
			if err == io.EOF{
				break
			}else if err != nil{
				log.Fatal(err)
				continue
			}
			
		}
	}

	
	
	return 1
	//c := pb.NewProtosClient(conn)
}
func main() {
	//var tiempo_pedidos string
	var tipo_cliente string
	var accion string
	var enviado int
	cantidadPedidosPyme := contarPedidos("./archivos/pymes.csv") - 1 
	cantidadPedidosRetail := contarPedidos("./archivos/retail.csv") - 1

	log.Printf(strconv.Itoa(cantidadPedidosPyme))
	log.Printf(strconv.Itoa(cantidadPedidosRetail))

	total_pedidos := 0 
	pedido_pymes := 1
	pedido_retail := 1
	log.Printf("Estableciendo conexión con logistica...")
	
	//Establecemos conexión con logisitica dist70:6969
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(puerto, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo establecer la conexión. ERROR: %v", err)
	}
	defer conn.Close()

	//Interfaz inicial
	log.Printf("[Cliente] Ingrese el tiempo entre pedidos:")
	//fmt.Scanln(&tiempo_pedidos)

	log.Printf("[Cliente] Seleccione el tipo de cliente que corresponde: ")
	log.Printf("1. Pyme")
	log.Printf("2. Retail")
	fmt.Scanln(&tipo_cliente)

	log.Printf("[Cliente] ¿Qué desea hacer ?")
	log.Printf("1. Enviar un pedido")
	log.Printf("2. Consultar estado de un pedido")
	fmt.Scanln(&accion)
	
	switch accion {
	case "1":
		//Enviar un pedido , pyme o retail dado por el tipo de cliente	
		if (tipo_cliente == "1"){
			enviado = enviarPedido(conn, pedido_pymes, tipo_cliente)
			if (enviado == 1){
				log.Printf("Pedido pyme enviado")
			}
			pedido_pymes++;
			total_pedidos++;
		}else if (tipo_cliente == "2"){
			enviado = enviarPedido(conn, pedido_retail, tipo_cliente)
			if (enviado == 1){
				log.Printf("Pedido retail enviado")
			}
			pedido_retail++;
			total_pedidos++;
		}
	case "2":
		//consultar estado del pedido con el codigo
		
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
