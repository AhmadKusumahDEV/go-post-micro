package tests

import (
	"bytes"
	"compress/gzip"
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

	dataset, err := r.Get(context.Background(), "product").Bytes()
	fmt.Println("jedor")

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

func TestPost(t *testing.T) {
	ctx := context.Background()
	product := domain.PostProduct{
		Image:    "https://i.ytimg.com/vi/6rkaGCZFlp0/maxresdefault.jpg",
		Desc:     "test",
		Price:    1000,
		Title:    "test",
		Category: "coffe",
	}

	data, err := json.Marshal(product)
	if err != nil {
		fmt.Println(err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://mongokopikan.vercel.app/products", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(dat))
	defer resp.Body.Close()

	// request list data ulang
	req, err = http.NewRequestWithContext(ctx, http.MethodGet, "https://mongokopikan.vercel.app/products", nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Accept-Encoding", "gzip")

	resp, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var dataa []byte
	if resp.Header.Get("Content-Encoding") == "gzip" {
		gzipReader, err := gzip.NewReader(bytes.NewReader(dat))
		if err != nil {
			fmt.Println(err)
		}
		defer gzipReader.Close()

		dataa, err = io.ReadAll(gzipReader)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(dataa))
	}
}

func TestUpdated(t *testing.T) {
	ctx := context.Background()
	dataProduct := domain.UpdateProduct{
		Id:       "67b36c18723eaee4691c4e7d",
		Image:    "https://i.ytimg.com/vi/6rkaGCZFlp0/maxresdefault.jpg",
		Desc:     "test updateddddddd",
		Price:    1000,
		Title:    "test updated 12321131323123123123123123",
		Category: "coffee late to update",
	}

	url := fmt.Sprintf("https://mongokopikan.vercel.app/products/%s", dataProduct.Id)

	// Marshal dataProduct ke JSON
	data, err := json.Marshal(dataProduct)
	if err != nil {
		t.Fatalf("Failed to marshal data: %v", err)
	}
	fmt.Println("Request Body:", string(data)) // Debug: Cetak body request

	// Buat HTTP request
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, url, bytes.NewBuffer(data))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "TestClient/1.0") // Tambahkan header User-Agent

	// Kirim request menggunakan HTTP client dengan timeout
	client := &http.Client{
		Timeout: 15 * time.Second, // Timeout untuk mencegah request menggantung
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	// Baca response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(body)) // Debug: Cetak response body

	// Cek status code
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Unexpected status code: %d, Body: %s", resp.StatusCode, string(body))
	}
}
