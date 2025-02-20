package product

import (
	"context"
	"fmt"
	"log"

	"github.com/AhmadKusumahDEV/go-post-micro/api-gateway/internal/domain"
	"github.com/AhmadKusumahDEV/go-post-micro/api-gateway/internal/types"
	"github.com/AhmadKusumahDEV/go-post-micro/api-gateway/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
)

type ServicesProductImpl struct {
	Repository  RepositoryProduct
	RedisClient *redis.Client
	Validator   *validator.Validate
}

// AddProduct implements ServicesProduct.
func (s *ServicesProductImpl) AddProduct(ctx context.Context, ch types.ChannelGroup) {
	data := ctx.Value(types.PostKeyCtx).(domain.PostProduct)
	if err := s.Validator.Struct(data); err != nil {
		log.Println(err)
		ch["err"] <- false
		return
	} else {
		ch["err"] <- true
		DataProduct := s.Repository.CreateProduct(ctx, ch)
		if DataProduct != nil {
			log.Println("gagal post and send to message broker")
		}
		utils.InsertRedisByteSync(DataProduct)
		return
	}
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
func (s *ServicesProductImpl) ModifyProduct(ctx context.Context, product domain.UpdateProduct, ch types.ChannelGroup) {
	ch["err"] <- true
	DataProduct := s.Repository.UpdateProduct(ctx, product)
	if DataProduct != nil {
		log.Println("failed Updated (services) and send to message broker")
	}
	utils.InsertRedisByteSync(DataProduct)
}

// RemoveProduct implements ServicesProduct.
func (s *ServicesProductImpl) RemoveProduct(ctx context.Context, id domain.DeleteProduct, ch types.ChannelGroup) {
	if err := s.Validator.Struct(id); err != nil {
		log.Println("error on validator (services) :", err)
		ch["err"] <- false
		return
	} else {
		ch["err"] <- true
		DataProduct := s.Repository.DeleteProduct(ctx, id.Id)
		if DataProduct != nil {
			log.Println("gagal post and send to message broker")
		}
		utils.InsertRedisByteSync(DataProduct)
		return
	}
}

func NewServicesProduct(repository RepositoryProduct, redisClient *redis.Client, validator *validator.Validate) ServicesProduct {
	return &ServicesProductImpl{
		Repository:  repository,
		RedisClient: redisClient,
		Validator:   validator,
	}
}
