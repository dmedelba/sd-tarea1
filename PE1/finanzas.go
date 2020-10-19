package main

import (

	"fmt"
	"log"
	"github.com/streadway/amqp"
	"encoding/json"
	"strconv"
)


//  función auxiliar para verificar el valor 
//de retorno de cada llamada amqp
func failOnError(err error, msg string) {
	if err != nil {
	  log.Fatalf("%s: %s", msg, err)
	  panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

//se declara la estructa de los registros que se reciben
type registro struct{
	Idpaquete string
	Fechaentrega string // 0 -> no entregado 
	Tipo string // retail,prioritario,normal
	Intentos string //3,2,1,0
	Valor string // valor del productito
}

func main() {

	//se define la variable que recibe el registro
	var regis registro 

	//se define la variable que tendra el balance total
	var baltotal int
	baltotal = 0

	//conexion del servidor rabbit
	// reemplazar localhost con IP de logistica 
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	//canal, que es donde reside la mayor parte de la API 
	//para hacer las cosas
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// declaracion de cola que debe ser la 
	// misma que envia uwu
	q, err := ch.QueueDeclare(
		"hello-queue", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	
	//recibo los mensajes de la cola
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")
	
	forever := make(chan bool)
	
	go ​func() {
		​for d := range msgs {
			​log.Printf("Received a message: %s", d.Body) // d.body es el mensaje
			json.Unmarshal(d.body, &regis) // guardo en regis el msg

			//AQUI DEFINO LA FUNCION PARA CALCULAR GANANCIAS Y PERDIDAS
			
			// defino variables auxiliares
			var val int
			var cantintento int 

			id := regis.Idpaquete 
			fechaent := regis.Fechaentrega // 0 -> no entregado  
			tipe := regis.Tipo  // retail,prioritario,normal
			intento := regis.Intentos //3,2,1
			valor := regis.Valor // valor  del productito

			val,_ = strconv.Atoi(valor) //se pasa el valor a int
			cantintento,_ = strconv.Atoi(intento) //se pasa el intento a int


			// se define las variables a utilizar para el calculo
			var ganancia int
			ganancia = 0
			var perdida int
			perdida = 0
			var total int

			if (fechaent != "0"){
				ganancia = val 
				perdida = (cantintento-1)*10 
			}else {
				if(tipe == "normal"){
					ganancia = 0
					perdida = (cantintento-1)*10
				}else if (tipe == "prioritario"){
					ganancia = int(val*0.3)
					perdida = (cantintento-1)*10
				}else{
					ganancia = val
					perdida = (cantintento-1)*10
				}
			}

			total = ganancia - perdida
			baltotal = baltotal + total

			// FALTA METER LA WEA AL ARCHIVO PARA BALANCE GRAL.
		​}

	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}