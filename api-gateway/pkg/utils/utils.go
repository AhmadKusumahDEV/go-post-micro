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
	Err(err, "json decode error")
}

func Encode_Json(w http.ResponseWriter, result any) {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(result)
	Err(err, "json encode error")
}

func Err(e error, msg string) {
	if e != nil && msg == "" {
		log.Println(e)
	} else if e != nil {
		log.Fatalln(msg, e)
	}
}

func UrlParse(ul string) *url.URL {
	u, err := url.Parse(ul)
	Err(err, "url parse error")
	return u
}

func CopyHeaderRequest(r *http.Request, header http.Header) {
	for k, v := range header {
		for _, vv := range v {
			r.Header.Add(k, vv)
		}
	}
}
