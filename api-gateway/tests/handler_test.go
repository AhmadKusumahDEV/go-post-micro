package tests

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

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
