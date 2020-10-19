
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
	puerto     = "dist70:7071"
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
			var tipo_pedido string
			if (i == id_pedido){
				if (line[5]=="1"){
					tipo_pedido = "prioritario"
				}else{
					tipo_pedido = "normal"
				}
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()
				r, err := c.SolicitarPedidoPyme(ctx, &pb.SolicitudPedidoPyme{
					IdPaquete:line[0],
					Tipo: tipo_pedido, 
					Nombre: line[1], 
					Valor: line[2], 
					Origen:line[3], 
					Destino:line[4]})
				if err != nil {
					log.Fatalf("No se pudo enviar el pedido. ERROR: %v", err)
				}
				//respuesta servidor
				log.Printf("[Servidor] Codigo de seguimiento: " + r.CodigoSeguimiento)
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
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()
				r, err := c.SolicitarPedidoRetail(ctx, &pb.SolicitudPedidoRetail{
					IdPaquete:line[0], 
					Nombre: line[1], 
					Valor: line[2], 
					Origen:line[3], 
					Destino:line[4]})
				if err != nil {
					log.Fatalf("No se pudo enviar el pedido. ERROR: %v", err)
				}
				//respuesta servidor
				log.Printf("[Servidor] Codigo de seguimiento: " + r.CodigoSeguimiento)
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

func consultarEstadoPedido(conn *grpc.ClientConn, codigoSeguimiento string){
	c := pb.NewProtosClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.ObtenerCodigoSeguimiento(ctx, &pb.SolicitudSeguimiento{
		CodigoSeguimiento:codigoSeguimiento})
	if err != nil {
		log.Fatalf("No se pudo solicitar el estado del pedido. ERROR: %v", err)
	}
	//respuesta servidor
	log.Printf("[Servidor] Estado del paquete " + codigoSeguimiento + " :"  + r.EstadoPedido)
}
func main() {
	var tiempo_pedidos string
	var tipo_cliente string
	var accion string
	var enviado int
	var code string

	cantidadPedidosPyme := contarPedidos("./archivos/pymes.csv")  
	cantidadPedidosRetail := contarPedidos("./archivos/retail.csv") 
	pedidosMax := cantidadPedidosPyme + cantidadPedidosRetail - 2

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
	fmt.Scanln(&tiempo_pedidos)
	tiempoEspera, err := strconv.Atoi(tiempo_pedidos)
	for true{
		log.Printf("[Cliente] Seleccione el tipo de cliente que corresponde: ")
		log.Printf("1. Pyme")
		log.Printf("2. Retail")
		log.Printf("0. Para finalizar conexión")
		fmt.Scanln(&tipo_cliente)

		if (tipo_cliente == "0"){
			break
		}

		log.Printf("[Cliente] ¿Qué desea hacer ?")
		log.Printf("1. Enviar un pedido")
		log.Printf("2. Consultar estado de un pedido")
		fmt.Scanln(&accion)

		switch accion {
		case "1":
			//Enviar un pedido , pyme o retail dado por el tipo de cliente
			if (total_pedidos == pedidosMax){
				log.Printf("[Cliente] No quedan más pedidos disponibles. ")
				break
			}

			switch tipo_cliente {
			case "1":
				if (pedido_pymes < cantidadPedidosPyme){
					enviado = enviarPedido(conn, pedido_pymes, tipo_cliente)
					if (enviado == 1){
						log.Printf("---------------------------------")
					}
					pedido_pymes++;
				}else{
					log.Printf("[Cliente] No quedan más pedidos pyme")
				}
			case "2":
				if (pedido_retail < cantidadPedidosRetail){
					enviado = enviarPedido(conn, pedido_retail, tipo_cliente)
					if (enviado == 1){ 
						log.Printf("---------------------------------")
					}
					pedido_retail++;
	
				}else{
					log.Printf("[Cliente] No quedan más pedidos retail")
				}
			}	
			 
			total_pedidos ++
			
		case "2":
			//consultar estado del pedido con el codigo
			if (total_pedidos>0){
				log.Printf("[Cliente] Ingrese el código de seguimiento a consultar: ")
				fmt.Scanln(&code)
				consultarEstadoPedido(conn, code)
			} else{
				log.Printf("[Cliente] No se ha realizado ningún pedido aún.")
			}
			
		}
		//tiempo entre pedidos
		log.Printf("[Espera] Esperando " + tiempo_pedidos +" segundos para realizar otro pedido")
		time.Sleep(time.Duration(tiempoEspera) * time.Second)
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
	log.Printf("Conexión finalizada")
}
