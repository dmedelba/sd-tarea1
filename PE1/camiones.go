package main

import (
	//"os"
	"fmt"
	//"bufio"
	//"io"
	"strconv"
	"context"
	"log"
	"time"
	//"encoding/csv"
	"google.golang.org/grpc"
	pb "github.com/dmedelba/sd-tarea1/PE1/protos"
)

//id-paquete, tipo de paquete, valor, origen, destino, n´umero de intentos y fecha de entrega.
type paquete struct {
	Idpaquete    string
	Tipo         string // retail,prioritario,normal
	Valor        int // valor del productito
	Origen       string
	Destino      string
	Intentos     string //3,2,1
	Fechaentrega string // 0 -> no entregado
}

// Estrucctura de un camion
type camion struct {
	Tipo     string // Retail o normal
	Paquete1 paquete
	Paquete2 paquete
	Estado   int // 0 = vacio; 1 = medio; 2 = lleno
}

// recibo un paquete y se lo asigno a un camion

func agregarpaquetes(conn *grpc.ClientConn, camioncito *camion ) {
	//conexión con el servidor 
	c := pb.NewProtosClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	r, err := c.SolicitarPaquete(ctx, &pb.SolicitudPaquete{Tipo: camioncito.Tipo})

	if (camioncito.Estado == 0){
		camioncito.Paquete1 = paquete{ Idpaquete:r.IdPaquete, Tipo:r.Tipo, Valor:r.Valor , Origen:r.Origen, 
										Destino:r.Destino, Intentos: "0", Fechaentrega: "0"}
		camioncito.Estado = 1
	}else {
		camioncito.Paquete2 = paquete{ Idpaquete:r.IdPaquete, Tipo:r.Tipo, Valor:r.Valor , Origen:r.Origen, 
			Destino:r.Destino, Intentos: "0", Fechaentrega: "0"}
		camioncito.Estado = 2
	}
	// No se considerará el caso en que el estado es 2 puesto que con ese estado esta en proceso de envio y no
	// de asignación de paquetes
}



func main() {
	var tiempo_pedidos string
	// se define cada camion (2 retail y 1 pyme)
	camion1 := &camion{Tipo: "Retail", Estado: 0}
	camion2 := &camion{Tipo: "Retail", Estado: 0}
	camion3 := &camion{Tipo: "Pyme", Estado: 0}

	//Establecemos conexión con logisitica dist70:6970
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("dist70:7080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo establecer la conexión. ERROR: %v", err)
	}
	defer conn.Close()

	//Interfaz inicial
	log.Printf("[Camion] Ingrese el tiempo entre entregas de pedidos:")
	fmt.Scanln(&tiempo_pedidos)
	tiempoEspera, err := strconv.Atoi(tiempo_pedidos)

	// A veces, es posible que no conozca el número de iteraciones que necesitará 
	//para completar una tarea concreta. En ese caso, puede omitir todas las instrucciones
	// y usar la palabra clave break para cerrar la ejecución.
	for{

		// IF QUEDAN PAQUETES EN LA COLA RETAIL;
			// SI NO PREGUNTAR SI QUEDAN PAQUETES PRIORITARIOS EN LA COLA PYMES;
		//cargamos camion 1 
		agregarpaquetes(conn, camion1) // AGREGO PAQUETE 1 (ESTADO = 1) (MEDIO)
		//VOLVER A PREGUNTAR SI QUEDAN PAQUETES EN COLA
		agregarpaquetes(conn, camion1) // AGREGO PAQUETE 2 (ESTADO = 2) (LLENO)

		// IF QUEDAN PAQUETES EN LA COLA RETAIL;
			// SI NO PREGUNTAR SI QUEDAN PAQUETES PRIORITARIOS EN LA COLA PYMES;
		//cargamos camion 2 
		agregarpaquetes(conn, camion2) // AGREGO PAQUETE 1 (ESTADO = 1) (MEDIO)
		//VOLVER A PREGUNTAR SI QUEDAN PAQUETES EN COLA
		agregarpaquetes(conn, camion2) // AGREGO PAQUETE 2 (ESTADO = 2) (LLENO)

		// IF QUEDAN PAQUETES EN LA COLA PYMES
		//cargamos camion 3 
		agregarpaquetes(conn, camion3) // AGREGO PAQUETE 1 (ESTADO = 1) (MEDIO)
		agregarpaquetes(conn, camion3) // AGREGO PAQUETE 2 (ESTADO = 2) (LLENO)

		// IMPLEMENTAR UNA FUNCION QUE HAGA EL PROCESO DE REPARTO (LO DIFICIL ESTA AQUI)

	}

}
