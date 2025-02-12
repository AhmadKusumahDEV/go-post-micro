package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

func Decode_Json(req *http.Request, result any) {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(result)
	if err != nil {
		log.Fatal(err)
	}
}

func Encode_Json(w http.ResponseWriter, result any) {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(result)
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
