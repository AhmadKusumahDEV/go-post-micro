package api

import "net/http"

type CorsMiddelware struct {
	Handler http.Handler
}

func (c *CorsMiddelware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Credentials", "true")
	writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")
	if request.Method == "OPTIONS" {
		writer.WriteHeader(http.StatusNoContent)
		return
	}
	c.Handler.ServeHTTP(writer, request)
}
