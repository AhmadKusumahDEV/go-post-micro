package product

import (
	"context"
	"net/http"

	"github.com/AhmadKusumahDEV/go-post-micro/api-gateway/internal/domain"
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
func (r *RepositoryProductImpl) ListProduct(ctx context.Context) (domain.Product, error) {
	panic("unimplemented")
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
