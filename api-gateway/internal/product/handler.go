package product

import (
	"context"
	"net/http"

	"github.com/AhmadKusumahDEV/go-post-micro/api-gateway/internal/types"
	"github.com/AhmadKusumahDEV/go-post-micro/api-gateway/pkg/utils"
)

type HandlerProductImpl struct {
	Services ServicesProduct
}

// AddProduct implements HandlerProduct.
func (h *HandlerProductImpl) AddProduct(writer http.ResponseWriter, request *http.Request) {
	panic("unimplemented")
}

// GetProductList implements HandlerProduct.
func (h *HandlerProductImpl) GetProductList(writer http.ResponseWriter, request *http.Request) {
	ctx := context.WithValue(request.Context(), types.HeaderKey, request.Header)
	h.Services.GetProductList(ctx)
	utils.Encode_Json(writer, struct {
		Name string
	}{
		Name: "test",
	})
}

// ModifyProduct implements HandlerProduct.
func (h *HandlerProductImpl) ModifyProduct(writer http.ResponseWriter, request *http.Request) {
	panic("unimplemented")
}

// RemoveProduct implements HandlerProduct.
func (h *HandlerProductImpl) RemoveProduct(writer http.ResponseWriter, request *http.Request) {
	panic("unimplemented")
}

func NewHandlerProductImp(services ServicesProduct) HandlerProduct {
	return &HandlerProductImpl{
		Services: &ServicesProductImpl{},
	}
}
