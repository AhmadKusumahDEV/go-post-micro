package api

import (
	"github.com/AhmadKusumahDEV/go-post-micro/api-gateway/config"
	"github.com/AhmadKusumahDEV/go-post-micro/api-gateway/internal/product"
)

func ProductRouter(serv *config.Server, products product.HandlerProduct) {
	serv.Addhandler("/Product", products.GetProductList)
}
