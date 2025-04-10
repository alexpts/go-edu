package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/alexpts/edu-go/examples/grpc/contract"
	"github.com/alexpts/edu-go/examples/grpc/server"
)

func main() {
	grpcServer := grpc.NewServer()
	service := server.ServiceA{}

	contract.RegisterAServer(grpcServer, &service)

	tcp := tcpListener()
	if err := grpcServer.Serve(tcp); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func tcpListener() net.Listener {
	tcpServer, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("tcp server listening at %v", tcpServer.Addr())

	return tcpServer
}
