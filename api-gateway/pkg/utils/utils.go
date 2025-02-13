package utils

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/AhmadKusumahDEV/go-post-micro/api-gateway/config"
)

func Decode_Json(req *http.Request, result any) {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(result)
	if err != nil {
		log.Fatal(err)
	}
}

func Encode_Json(w http.ResponseWriter, result any) {
	err := json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Fatal(err)
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

func InsertRedisByte(data []byte, ch chan bool) {
	redis := config.InitRedis()
	redis.Set(context.Background(), "product", data, 5*time.Minute)
	defer func() {
		ch <- true
	}()
}
