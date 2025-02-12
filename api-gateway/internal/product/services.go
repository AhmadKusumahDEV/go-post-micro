package product

import (
	"context"
	"fmt"
	"log"

	"github.com/AhmadKusumahDEV/go-post-micro/api-gateway/internal/domain"
	"github.com/AhmadKusumahDEV/go-post-micro/api-gateway/internal/types"
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
func (s *ServicesProductImpl) GetProductList(ctx context.Context) ([]byte, context.Context, error) {
	data, err := s.RedisClient.Get(context.Background(), "product").Bytes()
	if err != nil {
		log.Printf(types.ErrRedis.Error(), err)
	}

	if data != nil {
		return data, nil, nil
	}

	by, context, err := s.Repository.ListProduct(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf(types.ErrRepository.Error(), err)
	}
	return by, context, nil
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
