package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/alexpts/edu-go/examples/grpc/server"
)

func main() {
	grpcServer := server.New(
		grpc.NewServer(),
		tcpListener(),
		server.ServiceA{},
	)

	if err := grpcServer.Serve(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func tcpListener() *net.Listener {
	tcpServer, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("tcp server listening at %v", tcpServer.Addr())

	return &tcpServer
}
