package server

import (
	"net"

	gGrpc "google.golang.org/grpc"

	"github.com/alexpts/edu-go/examples/grpc/contract"
)

// Server декоратор на над grpc.Server
type Server struct {
	GrpcServer *gGrpc.Server
	listener   *net.Listener
}

func New(server *gGrpc.Server, listener *net.Listener, a ServiceA) *Server {
	wrapServer := &Server{
		GrpcServer: server,
		listener:   listener,
	}

	// регистрируем все сервисы в сервере
	contract.RegisterAServer(server, &a)

	return wrapServer
}

func (s *Server) Serve() error {
	return s.GrpcServer.Serve(*s.listener)
}
