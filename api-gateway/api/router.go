package api

import (
	"net/http"

	"github.com/AhmadKusumahDEV/go-post-micro/api-gateway/config"
	"github.com/AhmadKusumahDEV/go-post-micro/api-gateway/internal/product"
)

func ProductRouter(serv *config.Server, products product.HandlerProduct) {
	serv.Addhandler("/product", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			products.GetProductList(w, r)
		case http.MethodPost:
			products.AddProduct(w, r)
		case "PUT":
			products.ModifyProduct(w, r)
		case "DELETE":
			products.RemoveProduct(w, r)
		}
	})
}
