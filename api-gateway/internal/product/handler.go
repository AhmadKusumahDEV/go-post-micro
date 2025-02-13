package product

import (
	"context"
	"net/http"

	"github.com/AhmadKusumahDEV/go-post-micro/api-gateway/internal/domain"
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
	data, contex, err := h.Services.GetProductList(ctx)
	if err != nil {
		http.Error(writer, "internal server error", http.StatusInternalServerError)
		return
	}

	// data di ambil langsung dari product services
	if contex != nil {
		go utils
		headerResponse := contex.Value(types.RespHeader).(http.Header)
		for key, values := range headerResponse {
			for _, value := range values {
				writer.Header().Add(key, value)
			}
		}
		response := domain.Response{
			Status: "selamat anda mendapatkan datanya",
			Code:   http.StatusOK,
			Data:   data,
		}
		utils.Encode_Json(writer, response)
		return
	}
	// data yang diambil dari redis server
	response := domain.Response{
		Status: "selamat anda mendapatkan datanya",
		Code:   http.StatusOK,
		Data:   data,
	}
	utils.Encode_Json(writer, response)
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
