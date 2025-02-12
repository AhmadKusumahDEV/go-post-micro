package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/AhmadKusumahDEV/go-post-micro/api-gateway/config"
	"github.com/AhmadKusumahDEV/go-post-micro/api-gateway/internal/domain"
	"github.com/AhmadKusumahDEV/go-post-micro/api-gateway/internal/product"
)

func TestHandler(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	writer := httptest.NewRecorder()

	hadnler := product.HandlerProductImpl{}

	hadnler.GetProductList(writer, request)

	response := writer.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

}

func TestConvert(t *testing.T) {
	r := config.InitRedis()

	fmt.Println("test")

	var data []domain.Product
	resp, err := http.Get("https://mongokopikan.vercel.app/products")
	if err != nil {
		fmt.Println(err)
	}

	readAll, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(readAll, &data)
	if err != nil {
		fmt.Println(err)
	}
	r.Set(context.Background(), "babi2", readAll, 5*time.Minute)
	defer resp.Body.Close()
	for _, v := range data {
		fmt.Println("id: ", v.Id)
		fmt.Println("price: ", v.Price)
		fmt.Println("desc: ", v.Desc)
		fmt.Println("image: ", v.Image)
	}

	dataset, err := r.Get(context.Background(), "babi2").Bytes()
	if dataset != nil {
		fmt.Println(dataset)
	}
	if err != nil {
		fmt.Println(err)
	}

}

func TestGetDataRedis(t *testing.T) {
	r := config.InitRedis()

	var data []domain.Product

	dataset, err := r.Get(context.Background(), "babi2").Bytes()

	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(dataset, &data)
	if err != nil {
		fmt.Println(err)
	}

	if dataset != nil {
		fmt.Println(data)
	}
}
