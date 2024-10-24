package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {

	Port := os.Getenv("PORT")
	if Port == "" {
		Port = "50051"
	}

	listen, _ := net.Listen("tcp", ":"+Port)
	log.Println("Listening @ : " + Port)

	grpcserver := grpc.NewServer()

	grpcserver.Serve(listen)
	log.Println("Server Running")

}
