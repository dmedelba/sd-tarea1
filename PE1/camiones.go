package main

import (
	//"os"
	"fmt"
	"strconv"

	//"bufio"
	//"io"
	//"strconv"
	"context"
	"log"
	"math/rand"
	"time"

	//"encoding/csv"
	pb "github.com/dmedelba/sd-tarea1/PE1/protos"
	"google.golang.org/grpc"
)

//id-paquete, tipo de paquete, valor, origen, destino, n´umero de intentos y fecha de entrega.
type paquete struct {
	Idpaquete    string
	Tipo         string // retail,prioritario,normal
	Valor        string // valor del productito
	Origen       string
	Destino      string
	Intentos     string //3,2,1
	Fechaentrega string // 0 -> no entregado
}

// Estrucctura de un camion
type camion struct {
	Tipo     string // retail o normal
	Paquete1 paquete
	Paquete2 paquete
	Estado   int // 0 = vacio; 1 = medio; 2 = lleno
}

// recibo un paquete desde logistica (cola)y se lo asigno a un camion
func agregarpaquetes(conn *grpc.ClientConn, camioncito *camion) {
	//conexión con el servidor
	c := pb.NewProtosClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	r, err := c.SolicitarPaquete(ctx, &pb.SolicitudPaquete{Tipo: camioncito.Tipo})
	defer cancel()
	if err != nil {
		log.Fatalf("No se pudo agregar el paquete. ERROR: %v", err)
	}

	if camioncito.Estado == 0 {
		camioncito.Paquete1 = paquete{Idpaquete: r.IdPaquete, Tipo: r.Tipo, Valor: r.Valor, Origen: r.Origen,
			Destino: r.Destino, Intentos: "0", Fechaentrega: "0"}
		camioncito.Estado = 1
	} else {
		camioncito.Paquete2 = paquete{Idpaquete: r.IdPaquete, Tipo: r.Tipo, Valor: r.Valor, Origen: r.Origen,
			Destino: r.Destino, Intentos: "0", Fechaentrega: "0"}
		camioncito.Estado = 2
	}
	// No se considerará el caso en que el estado es 2 puesto que con ese estado esta en proceso de envio y no
	// de asignación de paquetes
}

//comunicamos el estado del paquete a logistica
func enviarEstadoPaquetes(conn *grpc.ClientConn,  camioncito *camion, estado string, paquete string){
	c := pb.NewProtosClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	var id_enviar string
	var intento_enviar string
	var fecha_enviar string
	if (paquete == "paquete1"){
		id_enviar = camioncito.Paquete1.Idpaquete
		intento_enviar = camioncito.Paquete1.Intentos
		fecha_enviar = camioncito.Paquete1.Fechaentrega
	}else if (paquete == "paquete2"){
		id_enviar = camioncito.Paquete2.Idpaquete
		intento_enviar = camioncito.Paquete2.Intentos
		fecha_enviar = camioncito.Paquete2.Fechaentrega
	}
	r, err := c.ObtenerEstado(ctx, &pb.SolicitudEstado{
		IdPaquete: id_enviar,
		Intentos: intento_enviar,
		Estado: estado,
		Fecha: fecha_enviar,
	})
	defer cancel()
	if err != nil {
		log.Fatalf("No se pudo enviar el estado del paquete a logistica. ERROR: %v", err)
	}

	if (r.Confirmacion == "1"){
		log.Printf("[Camion] Estado del pedido entregado a logistica.")
	}

}

//sumamos intento en string.
func sumarintento(cantidad_actual string) string {
	intente, _ := strconv.Atoi(cantidad_actual)
	return strconv.Itoa(intente + 1)
}

//funcion que recibe el camion y los tiempo de espera entre entregas de paquete
//la logica por dentras es poder entregar los paquetes y si falla aumenta la cantidad de intentos
func entregarpedidos(conn *grpc.ClientConn, camioncito *camion, tiempoEspera1 int, tiempoEspera2 int, paquete_listo string) {
	//se busca cual de los dos paquetes se entregara primero en caso de estar lleno
	// defino variables auxiliares
	var val1 int
	var val2 int

	if camioncito.Tipo == "retail" {
		if camioncito.Estado == 2 {
			valor1 := camioncito.Paquete1.Valor
			valor2 := camioncito.Paquete2.Valor
			intentoPaquete1, _ := strconv.Atoi(camioncito.Paquete1.Intentos)
			intentoPaquete2, _ := strconv.Atoi(camioncito.Paquete2.Intentos)
			val1, _ = strconv.Atoi(valor1)
			val2, _ = strconv.Atoi(valor2)

			// 2 PAQUETES  , 3 INTENTOS CADA PAQUETE,  
			//selecciona el paquete mas valioso, random segun la probabilidad
			if val1 > val2 {
				if (rand.Intn(100) < 80) && (intentoPaquete1 < 3) && (camioncito.Paquete1.Fechaentrega == "0") {
					time.Sleep(time.Duration(tiempoEspera1) * time.Second)
					camioncito.Paquete1.Fechaentrega = time.Now().String()
					camioncito.Estado = 1
					paquete_listo = "paquete1"
					enviarEstadoPaquetes(conn,  camioncito, "Recibido", paquete_listo)
					entregarpedidos(conn, camioncito, tiempoEspera1, tiempoEspera2, paquete_listo)
	
				} else {
					if intentoPaquete1 < 3 {
						var paquetito string
						time.Sleep(time.Duration(tiempoEspera1) * time.Second)
						camioncito.Paquete1.Intentos = sumarintento(camioncito.Paquete1.Intentos)
						time.Sleep(time.Duration(tiempoEspera2) * time.Second)
						entregarpedidos(conn, camioncito, tiempoEspera1, tiempoEspera2, paquetito)
					} else {
						//no pudo entregar el paquete => comunicar enviar			
						camioncito.Estado = 1
						paquete_listo = "paquete1"
						enviarEstadoPaquetes(conn,  camioncito, "No Recibido", paquete_listo)
						time.Sleep(time.Duration(tiempoEspera2) * time.Second)
						entregarpedidos(conn, camioncito, tiempoEspera1, tiempoEspera2,paquete_listo)
					}
				}
			} else {
				if (rand.Intn(100) < 80) && (intentoPaquete2 < 3) && (camioncito.Paquete2.Fechaentrega == "0") {
					time.Sleep(time.Duration(tiempoEspera1) * time.Second)
					camioncito.Paquete2.Fechaentrega = time.Now().String()
					camioncito.Estado = 1
					paquete_listo = "paquete2"
					enviarEstadoPaquetes(conn,  camioncito, "Recibido", paquete_listo)
					entregarpedidos(conn, camioncito, tiempoEspera1, tiempoEspera2, paquete_listo)
					// logistica y avisar que la entrega del paquete finalizo, y cambia su estado.
					
				} else {
					if intentoPaquete2 < 3 {
						var paquetito string
						time.Sleep(time.Duration(tiempoEspera1) * time.Second)
						camioncito.Paquete2.Intentos = sumarintento(camioncito.Paquete2.Intentos)
						time.Sleep(time.Duration(tiempoEspera2) * time.Second)
						entregarpedidos(conn, camioncito, tiempoEspera1, tiempoEspera2,paquetito)
					} else {
						paquete_listo = "paquete2"  //=> Se acabaron los intentos
						camioncito.Estado = 1
						enviarEstadoPaquetes(conn,  camioncito, "No Recibido", paquete_listo)
						time.Sleep(time.Duration(tiempoEspera2) * time.Second)
						entregarpedidos(conn, camioncito, tiempoEspera1, tiempoEspera2, paquete_listo)
					}
				}
			}
		} else {
			//ya entregamos un paquete, solo queda uno en el camión.
			//log.Printf("entro al else estado = 1")
			if (paquete_listo == "paquete2" || paquete_listo == ""){
				valor1 := camioncito.Paquete1.Valor
				intentoPaquete1, _ := strconv.Atoi(camioncito.Paquete1.Intentos)
				val1, _ = strconv.Atoi(valor1)

				if (rand.Intn(100) < 80) && (intentoPaquete1 < 3) {
					time.Sleep(time.Duration(tiempoEspera1) * time.Second)
					camioncito.Paquete1.Fechaentrega = time.Now().String()
					camioncito.Estado = 0
				} else{
					if intentoPaquete1 < 3 {
						var paquetito string
						time.Sleep(time.Duration(tiempoEspera1) * time.Second)
						camioncito.Paquete1.Intentos = sumarintento(camioncito.Paquete1.Intentos)
						time.Sleep(time.Duration(tiempoEspera2) * time.Second)
						entregarpedidos(conn,camioncito, tiempoEspera1, tiempoEspera2, paquetito)
					} else {
						camioncito.Estado = 0
					}
				}
			} else if (paquete_listo == "paquete1" || paquete_listo == "") {
				valor2 := camioncito.Paquete2.Valor
				intentoPaquete2, _ := strconv.Atoi(camioncito.Paquete2.Intentos)
				val2, _ = strconv.Atoi(valor2)

				if (rand.Intn(100) < 80) && (intentoPaquete2 < 3) {
					time.Sleep(time.Duration(tiempoEspera1) * time.Second)
					camioncito.Paquete1.Fechaentrega = time.Now().String()
					camioncito.Estado = 0
				} else {
					if intentoPaquete2 < 3 {
						var paquetito string
						time.Sleep(time.Duration(tiempoEspera1) * time.Second)
						camioncito.Paquete2.Intentos = sumarintento(camioncito.Paquete2.Intentos)
						time.Sleep(time.Duration(tiempoEspera2) * time.Second)
						entregarpedidos(conn, camioncito, tiempoEspera1, tiempoEspera2, paquetito)
					} else {
						camioncito.Estado = 0
					}
				}
			}

		}
	}else if (camioncito.Tipo == "normal"){

	}
}

func main() {
	var tiempo_pedidos string
	var tiempo_pedidos2 string
	// se define cada camion (2 retail y 1 pyme)
	camion1 := &camion{Tipo: "retail", Estado: 0}
	camion2 := &camion{Tipo: "retail", Estado: 0}
	camion3 := &camion{Tipo: "pyme", Estado: 0}

	//Establecemos conexión con logisitica dist70:6071
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("dist70:6071", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo establecer la conexión. ERROR: %v", err)
	}
	defer conn.Close()

	//Interfaz inicial
	log.Printf("[Camion] Ingrese el tiempo máximo para la entrega de un pedido:")
	fmt.Scanln(&tiempo_pedidos)
	log.Printf("[Camion] Ingrese el tiempo entre entregas de pedidos:")
	fmt.Scanln(&tiempo_pedidos2)

	tiempoEspera1, _ := strconv.Atoi(tiempo_pedidos)
	tiempoEspera2, _ := strconv.Atoi(tiempo_pedidos2)

	//cargamos camion 1
	if camion1.Estado == 0 {
		agregarpaquetes(conn, camion1) // AGREGO PAQUETE 1 (ESTADO = 1) (MEDIO)
		agregarpaquetes(conn, camion1) // AGREGO PAQUETE 2 (ESTADO = 2) (LLENO)
	}
	//cargamos camion 2
	if camion2.Estado == 0{
		agregarpaquetes(conn, camion2) // AGREGO PAQUETE 1 (ESTADO = 1) (MEDIO)
		agregarpaquetes(conn, camion2) // AGREGO PAQUETE 2 (ESTADO = 2) (LLENO)
	}

	//cargamos camion 3
	if camion3.Estado == 0{
		agregarpaquetes(conn, camion3) // AGREGO PAQUETE 1 (ESTADO = 1) (MEDIO)
		agregarpaquetes(conn, camion3) // AGREGO PAQUETE 2 (ESTADO = 2) (LLENO)
	}

	//mandamos lo camiones a ruta
	var paquetito string
	
	entregarpedidos(conn, camion1, tiempoEspera1, tiempoEspera2, paquetito)
	entregarpedidos(conn, camion2, tiempoEspera1, tiempoEspera2, paquetito)
	entregarpedidos(conn, camion3, tiempoEspera1, tiempoEspera2, paquetito)

	//agregamos el primer intento por paquete a los camiones
	camion1.Paquete1.Intentos = sumarintento(camion1.Paquete1.Intentos)
	camion1.Paquete2.Intentos = sumarintento(camion1.Paquete2.Intentos)

	camion2.Paquete1.Intentos = sumarintento(camion2.Paquete1.Intentos)
	camion2.Paquete2.Intentos = sumarintento(camion2.Paquete2.Intentos)

	camion3.Paquete1.Intentos = sumarintento(camion3.Paquete1.Intentos)
	camion3.Paquete2.Intentos = sumarintento(camion3.Paquete2.Intentos)


	// guargar los pedidos de camion el csv

}
