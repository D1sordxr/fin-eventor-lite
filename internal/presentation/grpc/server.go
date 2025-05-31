package grpc

import (
	"github.com/D1sordxr/fin-eventor-lite/internal/presentation/grpc/pb/services"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	Server         *grpc.Server
	accountService services.AccountServiceServer
}

func NewServer(
	config any, // TODO: Config struct for gRPC server settings
	accSvc services.AccountServiceServer,
) *Server {
	return &Server{
		Server:         grpc.NewServer(),
		accountService: accSvc,
	}
}

func (s *Server) StartServer() error {
	s.registerServices()

	listener, err := net.Listen("tcp", ":50051") // TODO: Use config for port
	if err != nil {
		return err
	}

	if err = s.Server.Serve(listener); err != nil {
		return err
	}

	return nil
}

func (s *Server) Shutdown() {
	s.Server.GracefulStop()
}

func (s *Server) registerServices() {
	services.RegisterAccountServiceServer(s.Server, s.accountService)
}
