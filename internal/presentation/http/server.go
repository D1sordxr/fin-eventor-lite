package http

import (
	"context"
	"errors"
	"net/http"

	"github.com/D1sordxr/fin-eventor-lite/internal/domain/ports"
	cfg "github.com/D1sordxr/fin-eventor-lite/internal/infrastructure/config/http"
)

type Handler interface {
	RegisterRoutes(mux *http.ServeMux)
}

type Server struct {
	log    ports.Log
	Server *http.Server
	Routes []Handler
}

func NewServer(
	log ports.Log,
	config *cfg.Config,
	routes ...Handler,
) *Server {
	log.Info("Initializing HTTP server", "port", config.Port)

	return &Server{
		log: log,
		Server: &http.Server{
			Addr:              ":" + config.Port,
			ReadHeaderTimeout: config.Timeout,
			ReadTimeout:       config.Timeout,
			WriteTimeout:      config.Timeout,
		},
		Routes: routes,
	}
}

func (s *Server) StartServer() error {
	s.log.Info("Registering HTTP routes...")
	mux := http.NewServeMux()
	for _, route := range s.Routes {
		route.RegisterRoutes(mux)
	}
	s.Server.Handler = mux

	s.log.Info("Starting HTTP server...", "address", s.Server.Addr)
	if err := s.Server.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			s.log.Info("HTTP server closed gracefully")
			return nil
		}
		s.log.Error("HTTP server stopped with error", "error", err.Error())
		return err
	}

	s.log.Info("HTTP server exited unexpectedly")
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.log.Info("Shutting down HTTP server...")
	err := s.Server.Shutdown(ctx)
	if err != nil {
		s.log.Error("Failed to gracefully shutdown HTTP server", "error", err.Error())
		return err
	}
	s.log.Info("HTTP server shutdown complete")
	return nil
}
