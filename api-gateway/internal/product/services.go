package product

import (
	"context"

	"github.com/AhmadKusumahDEV/go-post-micro/api-gateway/internal/domain"
	"github.com/redis/go-redis/v9"
)

type ServicesProductImpl struct {
	Repository  RepositoryProduct
	RedisClient *redis.Client
}

// AddProduct implements ServicesProduct.
func (s *ServicesProductImpl) AddProduct(ctx context.Context, product domain.Product) {
	panic("unimplemented")
}

// GetProductList implements ServicesProduct.
func (s *ServicesProductImpl) GetProductList(ctx context.Context) (domain.Product, error) {
	panic("unimplemented")
}

// ModifyProduct implements ServicesProduct.
func (s *ServicesProductImpl) ModifyProduct(ctx context.Context, product domain.Product) {
	panic("unimplemented")
}

// RemoveProduct implements ServicesProduct.
func (s *ServicesProductImpl) RemoveProduct(ctx context.Context, id string) {
	panic("unimplemented")
}

func NewServicesProduct(repository RepositoryProduct, redisClient *redis.Client) ServicesProduct {
	return &ServicesProductImpl{
		Repository:  repository,
		RedisClient: redisClient,
	}
}
