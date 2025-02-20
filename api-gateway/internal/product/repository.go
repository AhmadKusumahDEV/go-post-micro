package product

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/AhmadKusumahDEV/go-post-micro/api-gateway/internal/domain"
	"github.com/AhmadKusumahDEV/go-post-micro/api-gateway/internal/types"
)

type RepositoryProductImpl struct {
	http *http.Client
}

// CreateProduct implements RepositoryProduct.
func (r *RepositoryProductImpl) CreateProduct(ctx context.Context, ch types.ChannelGroup) []byte {
	header := ctx.Value(types.HeaderKey).(http.Header)
	c := ctx.Value(types.PostKeyCtx).(domain.PostProduct)
	sendctx := context.Background()
	data, err := json.Marshal(c)
	if err != nil {
		log.Printf("failed marshal json (repository): %v", err)
		return nil
	}

	// create request post data
	req, err := http.NewRequestWithContext(sendctx, http.MethodPost, "https://mongokopikan.vercel.app/products", bytes.NewBuffer(data))
	if err != nil {
		log.Printf("failed change format http request to post (repository): %v", err)
		return nil
	}

	// set header
	for key, values := range header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	req.Header.Set("Content-Type", "application/json")

	// send post data
	_, err = r.http.Do(req)
	if err != nil {
		log.Println("error send http post (repository): ", err)
		return nil
	}

	// change request format
	req, err = http.NewRequestWithContext(sendctx, http.MethodGet, "https://mongokopikan.vercel.app/products", nil)
	req.Header.Set("Accept-Encoding", "gzip")
	if err != nil {
		log.Printf("failed change format http request (repository): %v", err)
		return nil
	}

	resp, err := r.http.Do(req)
	if err != nil {
		log.Println("error send http request (repository): ", err)
		return nil
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Println("error read data request (repository): ", err)
		return nil
	}

	var dataByte []byte
	if resp.Header.Get("Content-Encoding") == "gzip" {
		gzipReader, err := gzip.NewReader(bytes.NewReader(data))
		if err != nil {
			log.Println("error read gzip request (repository): ", err)
			return nil
		}
		defer gzipReader.Close()

		dataByte, err = io.ReadAll(gzipReader)
		if err != nil {
			log.Println("error convert gzip (repository): ", err)
			return nil
		}
	}

	return dataByte
}

// DeleteProduct implements RepositoryProduct.
func (r *RepositoryProductImpl) DeleteProduct(ctx context.Context, id string) []byte {
	header := ctx.Value(types.HeaderKey).(http.Header)
	sendctx := context.Background()

	url := fmt.Sprintf("https://mongokopikan.vercel.app/products/%s", id)
	// create request post data
	req, err := http.NewRequestWithContext(sendctx, http.MethodDelete, url, nil)
	if err != nil {
		log.Printf("failed change format http request to post (repository): %v", err)
		return nil
	}

	// set header
	for key, values := range header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	req.Header.Set("Content-Type", "application/json")

	// send delete data
	_, err = r.http.Do(req)
	if err != nil {
		log.Println("error send http post (repository): ", err)
		return nil
	}

	// change request format
	req, err = http.NewRequestWithContext(sendctx, http.MethodGet, "https://mongokopikan.vercel.app/products", nil)
	req.Header.Set("Accept-Encoding", "gzip")
	if err != nil {
		log.Printf("failed change format http request (repository): %v", err)
		return nil
	}

	resp, err := r.http.Do(req)
	if err != nil {
		log.Println("error send http request (repository): ", err)
		return nil
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("error read data request (repository): ", err)
		return nil
	}

	var dataByte []byte
	if resp.Header.Get("Content-Encoding") == "gzip" {
		gzipReader, err := gzip.NewReader(bytes.NewReader(data))
		if err != nil {
			log.Println("error read gzip request (repository): ", err)
			return nil
		}
		defer gzipReader.Close()

		dataByte, err = io.ReadAll(gzipReader)
		if err != nil {
			log.Println("error convert gzip (repository): ", err)
			return nil
		}
	}

	return dataByte
}

// ListProduct implements RepositoryProduct.
func (r *RepositoryProductImpl) ListProduct(ctx context.Context) ([]byte, context.Context, error) {
	c := ctx.Value(types.HeaderKey).(http.Header)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://mongokopikan.vercel.app/products", nil)
	if err != nil {
		return nil, nil, fmt.Errorf(types.ErrCreateRequest.Error(), err)
	}

	// dari saya 15 02 2025
	// terjadi error pada encoding di response header bila lu nemu error kaya gitu lagi
	// tinggal make req.Header.Set("Accept-Encoding", "gzip") kenttodd gara gaara ginian doang
	// semaleman kontotl
	// catatan tambahan baik nya itu melaukan set accepth-encoding gzip ketika melakukan request
	//  namun bila pada response header nya tidak ngaco maka req.header gzip tidak di perlukan

	// copy header from context
	for key, values := range c {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}
	req.Header.Set("Accept-Encoding", "gzip")

	resp, err := r.http.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf(types.ErrSendRequest.Error(), err)
	}

	Header := context.WithValue(context.Background(), types.RespKeyHeader, resp.Header)

	defer func() {
		resp.Body.Close()
	}()

	read, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf(types.ErrRead.Error(), err)
	}

	if resp.Header.Get("Content-Encoding") == "gzip" {
		gzipReader, err := gzip.NewReader(bytes.NewReader(read))
		if err != nil {
			return nil, nil, fmt.Errorf("gzip invalid: %v", err)
		}
		defer gzipReader.Close()

		data, err := io.ReadAll(gzipReader)
		if err != nil {
			return nil, nil, fmt.Errorf("gagal baca gzip: %v", err)
		}
		return data, Header, nil
	}
	return read, Header, nil
}

// UpdateProduct implements RepositoryProduct.
func (r *RepositoryProductImpl) UpdateProduct(ctx context.Context, product domain.UpdateProduct) []byte {
	header := ctx.Value(types.HeaderKey).(http.Header)
	sendctx := context.Background()
	data, err := json.Marshal(product)
	if err != nil {
		log.Printf("failed marshal json (repository): %v", err)
		return nil
	}

	url := fmt.Sprintf("https://mongokopikan.vercel.app/products/%s", product.Id)

	// create request post data
	req, err := http.NewRequestWithContext(sendctx, http.MethodPatch, url, bytes.NewBuffer(data))
	if err != nil {
		log.Printf("failed change format http request to patch (repository): %v", err)
		return nil
	}

	// set header
	for key, values := range header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	req.Header.Set("Content-Type", "application/json")

	// send post data
	_, err = r.http.Do(req)
	if err != nil {
		log.Println("error send http Updated request (repository): ", err)
		return nil
	}

	// change request format
	req, err = http.NewRequestWithContext(sendctx, http.MethodGet, "https://mongokopikan.vercel.app/products", nil)
	req.Header.Set("Accept-Encoding", "gzip")
	if err != nil {
		log.Printf("failed change format http request (repository): %v", err)
		return nil
	}

	resp, err := r.http.Do(req)
	if err != nil {
		log.Println("error send http request (repository): ", err)
		return nil
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Println("error read data request (repository): ", err)
		return nil
	}

	var dataByte []byte
	if resp.Header.Get("Content-Encoding") == "gzip" {
		gzipReader, err := gzip.NewReader(bytes.NewReader(data))
		if err != nil {
			log.Println("error read gzip request (repository): ", err)
			return nil
		}
		defer gzipReader.Close()

		dataByte, err = io.ReadAll(gzipReader)
		if err != nil {
			log.Println("error convert gzip (repository): ", err)
			return nil
		}
	}

	return dataByte
}

func NewRepositoryProductImpl(http *http.Client) RepositoryProduct {
	return &RepositoryProductImpl{
		http: http,
	}
}
