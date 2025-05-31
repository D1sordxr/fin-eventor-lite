package http

import (
	"context"
	"errors"
	"net/http"
)

type Handler interface {
	RegisterRoutes(mux *http.ServeMux)
}

type Server struct {
	Server *http.Server
	Routes []Handler
}

func NewServer(
	port string,
	routes ...Handler,
) *Server {
	return &Server{
		Server: &http.Server{
			Addr: ":" + port,
		},
		Routes: routes,
	}
}

func (s *Server) StartServer() error {
	mux := http.NewServeMux()
	for _, route := range s.Routes {
		route.RegisterRoutes(mux)
	}

	s.Server.Handler = mux

	if err := s.Server.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			// Server closed
			return nil
		}
		return err
	}

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.Server.Shutdown(ctx)
}

// TODO: Client routes

type ClientHandler interface {
	RegisterClientRoutes(mux *http.ServeMux)
}

func (s *Server) RegisterClientRoutes(mux *http.ServeMux, routes ...ClientHandler) {
	for _, route := range routes {
		if clientHandler, ok := route.(ClientHandler); ok {
			clientHandler.RegisterClientRoutes(mux)
		}
	}
}
