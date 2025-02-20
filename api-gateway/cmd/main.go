package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AhmadKusumahDEV/go-post-micro/api-gateway/api"
	"github.com/AhmadKusumahDEV/go-post-micro/api-gateway/config"
	"github.com/AhmadKusumahDEV/go-post-micro/api-gateway/internal/domain"
	"github.com/AhmadKusumahDEV/go-post-micro/api-gateway/internal/product"
	"github.com/AhmadKusumahDEV/go-post-micro/api-gateway/pkg/utils"
	"github.com/go-playground/validator/v10"
)

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
	data := initContex.Value(headerkey).(http.Header)
	fmt.Println(data)
	utils.Encode_Json(writer, struct {
		Status string
		Code   int
		Data   any
	}{
		Status: "ok",
		Code:   200,
		Data:   "sangat baik",
	})
}

func handlerer(writer http.ResponseWriter, _ *http.Request) {
	// 1. Kirim response ke client terlebih dahulu
	log.Println("lama2")
	utils.Encode_Json(writer, domain.Response{
		Status: "ok",
		Code:   http.StatusOK,
		Data:   "berhasil post data",
	})
	log.Println("berhasil post dan bila gagal maka send to message broker")

	// 2. Jalankan proses async di goroutine (TANPA channel di handler)
	go func() {
		// Simulasi proses async (e.g., kirim ke message broker)
		time.Sleep(2 * time.Second)
		log.Println("Proses async selesai")

		// Jika gagal, kirim ke message broker di sini
		// ...
	}()
}

func main() {
	redis := config.InitRedis()
	httpClient := config.InitClient()
	validasi := validator.New()

	productRepository := product.NewRepositoryProductImpl(httpClient)
	ProductServices := product.NewServicesProduct(productRepository, redis, validasi)
	productHandler := product.NewHandlerProductImp(ProductServices)

	server := config.NewServer(":9090")
	server.Use(&api.CorsMiddelware{Handler: server.Mux.MuxHandler})
	api.ProductRouter(server, productHandler)
	server.Addhandler("/test", GetProductList)
	server.Addhandler("/test2", handlerer)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
