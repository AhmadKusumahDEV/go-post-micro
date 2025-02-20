package product

import (
	"context"
	"net/http"

	"github.com/AhmadKusumahDEV/go-post-micro/api-gateway/internal/domain"
	"github.com/AhmadKusumahDEV/go-post-micro/api-gateway/internal/types"
)

type RepositoryProduct interface {
	ListProduct(ctx context.Context) ([]byte, context.Context, error)
	CreateProduct(ctx context.Context, ch types.ChannelGroup) []byte
	UpdateProduct(ctx context.Context, product domain.UpdateProduct) []byte
	DeleteProduct(ctx context.Context, id string) []byte
}

type ServicesProduct interface {
	GetProductList(ctx context.Context) ([]byte, context.Context, error)
	AddProduct(ctx context.Context, ch types.ChannelGroup)
	ModifyProduct(ctx context.Context, product domain.UpdateProduct, ch types.ChannelGroup)
	RemoveProduct(ctx context.Context, id domain.DeleteProduct, ch types.ChannelGroup)
}

type HandlerProduct interface {
	GetProductList(writer http.ResponseWriter, request *http.Request)
	AddProduct(writer http.ResponseWriter, request *http.Request)
	ModifyProduct(writer http.ResponseWriter, request *http.Request)
	RemoveProduct(writer http.ResponseWriter, request *http.Request)
}
