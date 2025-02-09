package main

import (
	"fmt"
	"net/http"
)

func main() {

	muxHandler := http.NewServeMux()
	muxHandler.HandleFunc("/", handler)

	server := http.Server{
		Addr:    ":8085",
		Handler: muxHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}

func handler(w http.ResponseWriter, _ *http.Request) {

	fmt.Fprintln(w, "Hello World")
}
