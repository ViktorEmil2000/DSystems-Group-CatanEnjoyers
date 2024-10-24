package main

import (
	"log"
	"net"
	"os"

	"github.com/ViktorEmil2000/DSystems-Group-CatanEnjoyers/api/proto"
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

	cs := proto.ChitChatServer{}
	proto.RegisterServicesServer(grpcserver, &cs)

	grpcserver.Serve(listen)

}
