package product

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AhmadKusumahDEV/go-post-micro/api-gateway/internal/domain"
	"github.com/AhmadKusumahDEV/go-post-micro/api-gateway/internal/types"
	"github.com/AhmadKusumahDEV/go-post-micro/api-gateway/pkg/utils"
)

type HandlerProductImpl struct {
	Services ServicesProduct
}

// AddProduct implements HandlerProduct.
func (h *HandlerProductImpl) AddProduct(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	// init channel
	CheckErr := make(chan bool)
	Channel := types.ChannelGroup{
		"err": CheckErr,
	}

	dto := domain.PostProduct{}
	// decode json
	err := utils.Decode_Json(request, &dto)
	if err != nil {
		utils.Encode_Json(writer, domain.ResponseErr{
			Message: "error decode json",
			Status:  "internal server error",
			Code:    500,
		})
	}

	ctx = context.WithValue(ctx, types.PostKeyCtx, dto)
	ctx = context.WithValue(ctx, types.HeaderKey, request.Header)

	go h.Services.AddProduct(ctx, Channel)

	defer func() {
		close(CheckErr)
	}()

	select {
	case errPost := <-Channel["err"]:
		if !errPost {
			utils.Encode_Json(writer, domain.ResponseErr{
				Message: "Ada beberapa data yang tidak lengkap",
				Status:  "bad request",
				Code:    http.StatusBadRequest,
			})
			return
		} else {
			utils.Encode_Json(writer, domain.Response{
				Status: "ok",
				Code:   http.StatusOK,
				Data:   "berhasil post data",
			})
			log.Println("berhasil atau gagal akan tetap send to message broker")
			return
		}
	case <-time.After(1 * time.Second):
		utils.Encode_Json(writer, domain.ResponseErr{
			Message: "Request timeout",
			Status:  "error",
			Code:    http.StatusRequestTimeout,
		})
		return
	}

}

// GetProductList implements HandlerProduct.
func (h *HandlerProductImpl) GetProductList(writer http.ResponseWriter, request *http.Request) {
	var DataConvert []domain.Product
	ctx := context.WithValue(request.Context(), types.HeaderKey, request.Header)

	data, contex, err := h.Services.GetProductList(ctx)
	if err != nil {
		log.Println("error get data", err)
		http.Error(writer, "internal server error", http.StatusInternalServerError)
		return
	}
	if err := json.Unmarshal(data, &DataConvert); err != nil {
		http.Error(writer, "internal server error when decode", http.StatusInternalServerError)
		fmt.Print(err)
		return
	}

	// data di ambil langsung dari product services
	if contex != nil {
		utils.InsertRedisByte(data)
		headerResponse := contex.Value(types.RespKeyHeader).(http.Header)
		for key, values := range headerResponse {
			for _, value := range values {
				writer.Header().Set(key, value)
			}
		}
		writer.Header().Del("Content-Encoding")

		response := domain.Response{
			Status: "ok",
			Code:   http.StatusOK,
			Data:   DataConvert,
		}
		utils.Encode_Json(writer, response)
		return
	}
	// data yang diambil dari redis server
	response := domain.Response{
		Status: "cahce data",
		Code:   http.StatusOK,
		Data:   DataConvert,
	}

	utils.Encode_Json(writer, response)
}

// ModifyProduct implements HandlerProduct.
func (h *HandlerProductImpl) ModifyProduct(writer http.ResponseWriter, request *http.Request) {
	ctx := context.Background()
	// init channel
	CheckErr := make(chan bool)
	Channel := types.ChannelGroup{
		"err": CheckErr,
	}

	dto := domain.UpdateProduct{}
	// decode json
	err := utils.Decode_Json(request, &dto)
	if err != nil {
		utils.Encode_Json(writer, domain.ResponseErr{
			Message: "error decode json",
			Status:  "internal server error",
			Code:    500,
		})
	}

	ctx = context.WithValue(ctx, types.HeaderKey, request.Header)

	go h.Services.ModifyProduct(ctx, dto, Channel)

	defer func() {
		close(CheckErr)
	}()

	select {
	case errPost := <-Channel["err"]:
		utils.Encode_Json(writer, domain.Response{
			Status: "ok",
			Code:   http.StatusOK,
			Data:   "berhasil update data",
		})
		log.Println("berhasil atau gagal akan tetap send to message broker")
		fmt.Println(errPost)
		return
	case <-time.After(1 * time.Second):
		utils.Encode_Json(writer, domain.ResponseErr{
			Message: "Request timeout",
			Status:  "error",
			Code:    http.StatusRequestTimeout,
		})
		return
	}
}

// RemoveProduct implements HandlerProduct.
func (h *HandlerProductImpl) RemoveProduct(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	// init channel
	CheckErr := make(chan bool)
	Channel := types.ChannelGroup{
		"err": CheckErr,
	}

	dto := domain.DeleteProduct{}
	// decode json
	err := utils.Decode_Json(request, &dto)
	if err != nil {
		log.Println(types.ErrDecode.Error(), "handler", err)
		utils.Encode_Json(writer, domain.ResponseErr{
			Message: "error decode json",
			Status:  "internal server error",
			Code:    500,
		})
	}
	ctx = context.WithValue(ctx, types.HeaderKey, request.Header)

	go h.Services.RemoveProduct(ctx, dto, Channel)

	defer func() {
		close(CheckErr)
	}()

	select {
	case errPost := <-Channel["err"]:
		if !errPost {
			utils.Encode_Json(writer, domain.ResponseErr{
				Message: "id tidak terkirim",
				Status:  "bad request",
				Code:    http.StatusBadRequest,
			})
			return
		} else {
			utils.Encode_Json(writer, domain.Response{
				Status: "ok",
				Code:   http.StatusOK,
				Data:   "berhasil Hapus data",
			})
			log.Println("berhasil atau gagal akan tetap send to message broker")
			return
		}
	case <-time.After(1 * time.Second):
		utils.Encode_Json(writer, domain.ResponseErr{
			Message: "Request timeout",
			Status:  "error",
			Code:    http.StatusRequestTimeout,
		})
		return
	}
}

func NewHandlerProductImp(services ServicesProduct) HandlerProduct {
	return &HandlerProductImpl{
		Services: services,
	}
}

// func NewHandlerProductImp(services ServicesProduct) HandlerProduct {
// 	return &HandlerProductImpl{
// 		Services: &ServicesProductImpl{},
// 	}
// }
// ini code yang membuat error menjadi aneh yaitu
// runtime error: invalid memory address or nil pointer dereference
// kontotoll
