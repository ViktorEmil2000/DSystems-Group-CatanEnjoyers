package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	listen, _ := net.Listen("tcp", ":50051")
	log.Println("WOWOWOW WE HEar this boys!")

	grpcserver := grpc.NewServer()

	grpcserver.Listen(listen)
	log.Println("Eyy da surver runnin")

}
