
// Package main implements a client for Greeter service.
package main

import (
	"log"
	"google.golang.org/grpc"
)

const (
	address     = "dist70:6969"
	defaultName = "caca"
)

func main() {
	// Set up a connection to the server.
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()


	// Contact the server and print out its response.
	//name := defaultName
	//if len(os.Args) > 1 {
	//	name = os.Args[1]
	//}
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	
	log.Printf("Greeting: HOLAAAAA MACA")
}
