package product

import (
	"context"
	"io"
	"net/http"

	"github.com/AhmadKusumahDEV/go-post-micro/api-gateway/config"
	"github.com/AhmadKusumahDEV/go-post-micro/api-gateway/internal/domain"
	"github.com/AhmadKusumahDEV/go-post-micro/api-gateway/pkg/utils"
)

type RepositoryProductImpl struct {
	http *http.Client
}

// CreateProduct implements RepositoryProduct.
func (r *RepositoryProductImpl) CreateProduct(ctx context.Context, product domain.Product) {
	panic("unimplemented")
}

// DeleteProduct implements RepositoryProduct.
func (r *RepositoryProductImpl) DeleteProduct(ctx context.Context, id string) {
	panic("unimplemented")
}

// ListProduct implements RepositoryProduct.
func (r *RepositoryProductImpl) ListProduct(ctx context.Context) ([]byte, error) {
	c := ctx.Value(config.HeaderKey).(map[string][]string)
	req, err := http.NewRequest("GET", "https://mongokopikan.vercel.app/products", nil)
	utils.Err(err, "error on Repository product Services")

	for key, values := range c {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}
	resp, err := r.http.Do(req)

	defer func() {
		resp.Body.Close()
	}()

	utils.Err(err, "error on Repository product Services when get response data")
	read, err := io.ReadAll(resp.Body)
	utils.Err(err, "error on Repository product Services when read response data")
	return read, nil
}

// UpdateProduct implements RepositoryProduct.
func (r *RepositoryProductImpl) UpdateProduct(ctx context.Context, product domain.Product) {
	panic("unimplemented")
}

func NewRepositoryProductImpl(http *http.Client) RepositoryProduct {
	return &RepositoryProductImpl{
		http: http,
	}
}
