package http

import (
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
