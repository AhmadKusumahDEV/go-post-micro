package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/AhmadKusumahDEV/go-post-micro/api-gateway/pkg/utils"
)

type Server struct {
	Engine *http.Server
	Mux    *CustomMux
}

func (s *Server) GetListHandler() {
	for _, route := range s.Mux.Listhandler {
		fmt.Printf("Route: %s\n", route)
	}
}

func (s *Server) ListenAndServe() error {
	s.GetListHandler()
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
		Mux: &customMux,
	}
}

func GetProductList(writer http.ResponseWriter, request *http.Request) {
	type setkey string
	initContex := request.Context()
	req, err := http.NewRequest("GET", "https://mongokopikan.vercel.app/products", nil)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	const headerkey setkey = "header"

	for key, values := range request.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	fmt.Println(req.Header)
	initContex = context.WithValue(initContex, headerkey, req.Header)
	data := initContex.Value(headerkey)
	fmt.Println(data)
	utils.Encode_Json(writer, struct {
		Name    string
		Context any
	}{
		Name:    "test",
		Context: data,
	})
}

func main() {
	server := NewServer(":9090")

	server.Addhandler("/product", GetProductList)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
