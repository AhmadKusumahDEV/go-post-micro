package product

import (
	"context"
	"net/http"

	"github.com/AhmadKusumahDEV/go-post-micro/api-gateway/internal/domain"
)

type RepositoryProduct interface {
	ListProduct(ctx context.Context) (domain.Product, error)
	CreateProduct(ctx context.Context, product domain.Product)
	UpdateProduct(ctx context.Context, product domain.Product)
	DeleteProduct(ctx context.Context, id string)
}

type ServicesProduct interface {
	GetProductList(ctx context.Context) (domain.Product, error)
	AddProduct(ctx context.Context, product domain.Product)
	ModifyProduct(ctx context.Context, product domain.Product)
	RemoveProduct(ctx context.Context, id string)
}

type HandlerProduct interface {
	GetProductList(writer http.ResponseWriter, request *http.Request)
	AddProduct(writer http.ResponseWriter, request *http.Request)
	ModifyProduct(writer http.ResponseWriter, request *http.Request)
	RemoveProduct(writer http.ResponseWriter, request *http.Request)
}
