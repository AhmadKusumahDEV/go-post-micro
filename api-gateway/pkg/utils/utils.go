package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/AhmadKusumahDEV/go-post-micro/api-gateway/config"
	"github.com/AhmadKusumahDEV/go-post-micro/api-gateway/internal/domain"
)

func Decode_Json(req *http.Request, result any) error {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(result)
	if err != nil {
		return err
	}
	return nil
}

// func Encode_Json(w http.ResponseWriter, result any) {
// 	err := json.NewEncoder(w).Encode(result)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func Encode_Json(w http.ResponseWriter, result any) {
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewEncoder(w)
	err := decoder.Encode(result)
	if err != nil {
		json.NewEncoder(w).Encode(domain.ResponseErr{
			Message: "error",
			Status:  "gagal",
			Code:    500,
		})
	}
}

func UrlParse(ul string) *url.URL {
	u, err := url.Parse(ul)
	if err != nil {
		log.Fatal(err)
	}
	return u
}

func CopyHeaderRequest(r *http.Request, header http.Header) {
	for k, v := range header {
		for _, vv := range v {
			r.Header.Add(k, vv)
		}
	}
}

func InsertRedisByte(data []byte) {
	redis := config.InitRedis()
	err := redis.Set(context.Background(), "product", data, 30*time.Second).Err()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("berhasil")
	return
}

func InsertRedisByteSync(data []byte) {
	redis := config.InitRedis()
	err := redis.Set(context.Background(), "product", data, 30*time.Second).Err()
	if err != nil {
		log.Println(err)
	}
}

// func InitCors(w http.ResponseWriter) {
// 	w.Header().se
// }
