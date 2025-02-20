package config

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	Engine     *http.Server
	Mux        *CustomMux
	Middleware []http.Handler
}

func (s *Server) GetListHandler() {
	for _, route := range s.Mux.Listhandler {
		fmt.Printf("Route: %s\n", route)
	}
}

func (s *Server) ListenAndServe() error {
	s.GetListHandler()
	rang := len(s.Middleware)
	s.Engine.Handler = s.Middleware[rang-1]
	log.Println("Server is running on port", s.Engine.Addr)
	return s.Engine.ListenAndServe()
}

func (s *Server) Addhandler(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	s.PushListHandler(pattern)
	s.Mux.MuxHandler.HandleFunc(pattern, handler)
}

func (s *Server) PushListHandler(pattern string) {
	s.Mux.Listhandler = append(s.Mux.Listhandler, pattern)
}

func (s *Server) Use(middleware http.Handler) {
	s.Middleware = append(s.Middleware, middleware)
}

type CustomMux struct {
	MuxHandler  *http.ServeMux
	Listhandler []string
}

func NewServer(port string) *Server {
	customMux := CustomMux{
		MuxHandler:  http.NewServeMux(),
		Listhandler: []string{},
	}
	return &Server{
		Engine: &http.Server{
			Addr:    port,
			Handler: customMux.MuxHandler,
		},
		Mux:        &customMux,
		Middleware: []http.Handler{},
	}
}
