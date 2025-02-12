package product

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/AhmadKusumahDEV/go-post-micro/api-gateway/internal/domain"
	"github.com/AhmadKusumahDEV/go-post-micro/api-gateway/internal/types"
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
func (r *RepositoryProductImpl) ListProduct(ctx context.Context) ([]byte, context.Context, error) {
	c := ctx.Value(types.HeaderKey).(http.Header)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://mongokopikan.vercel.app/products", nil)
	if err != nil {
		return nil, nil, fmt.Errorf(types.ErrCreateRequest.Error(), err)
	}

	// copy header from context
	for key, values := range c {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	resp, err := r.http.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf(types.ErrSendRequest.Error(), err)
	}

	Header := context.WithValue(context.Background(), types.RespHeader, resp.Header)

	defer func() {
		resp.Body.Close()
	}()

	read, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf(types.ErrRead.Error(), err)
	}

	return read, Header, nil
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
