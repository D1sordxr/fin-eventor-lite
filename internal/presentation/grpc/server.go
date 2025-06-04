package grpc

import (
	"context"
	cfg "github.com/D1sordxr/fin-eventor-lite/internal/infrastructure/config/grpc"
	"github.com/D1sordxr/fin-eventor-lite/internal/presentation/grpc/pb/services"
	"github.com/D1sordxr/fin-eventor-lite/internal/shared/ports"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"net"
)

type Server struct {
	log            ports.Log
	Port           string
	Server         *grpc.Server
	accountService services.AccountServiceServer
}

func NewServer(
	log ports.Log,
	config *cfg.Config,
	accountService services.AccountServiceServer,
) *Server {
	log.Info("Initializing gRPC server", "port", config.Port)

	server := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		Time:              config.Time,
		Timeout:           config.Timeout,
		MaxConnectionIdle: config.MaxConnectionIdle,
		MaxConnectionAge:  config.MaxConnectionAge,
	}))

	return &Server{
		log:            log,
		Port:           config.Port,
		Server:         server,
		accountService: accountService,
	}
}

func (s *Server) StartServer() error {
	s.log.Info("Registering gRPC services...")
	s.registerServices()

	s.log.Info("Starting gRPC server...", "port", s.Port)
	listener, err := net.Listen("tcp", ":"+s.Port)
	if err != nil {
		s.log.Error("Failed to start TCP listener", "error", err)
		return err
	}

	err = s.Server.Serve(listener)
	if err != nil {
		s.log.Error("gRPC server stopped with error", "error", err)
		return err
	}

	s.log.Info("gRPC server stopped gracefully")
	return nil
}

func (s *Server) Shutdown(_ context.Context) error {
	s.log.Info("Shutting down gRPC server...")
	s.Server.GracefulStop()
	s.log.Info("gRPC server shutdown complete")
	return nil
}

func (s *Server) registerServices() {
	services.RegisterAccountServiceServer(s.Server, s.accountService)
}
