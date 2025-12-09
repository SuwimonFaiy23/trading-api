package service

import (
	"context"
	"trading-api/internal/domain/entity"
	"trading-api/internal/domain/repository"
	"trading-api/internal/dto"

	log "github.com/sirupsen/logrus"
	uuid "github.com/tentone/mssql-uuid"
)

type ProductService struct {
	productRepository repository.ProductRepository
}

type ProductServiceDependencies struct {
	ProductRepository repository.ProductRepository
}

func NewProductService(d ProductServiceDependencies) *ProductService {
	return &ProductService{
		productRepository: d.ProductRepository,
	}
}

func (s *ProductService) CreateProduct(ctx context.Context, reqBody dto.ProductRequest) (*dto.ProductResponse, error) {
	product := &entity.Product{
		ID:       uuid.NewV4(),
		Name:     reqBody.Name,
		Quantity: reqBody.Quantity,
		Price:    reqBody.Price,
	}
	if err := s.productRepository.Create(ctx, product); err != nil {
		log.WithFields(log.Fields{
			"layer":    "service",
			"function": "CreateProduct",
			"reqBody":  reqBody,
		}).WithError(err).Error("unable to create product")
		return nil, err
	}
	result := dto.ProductResponse{
		ID:       product.ID,
		Name:     product.Name,
		Quantity: product.Quantity,
		Price:    product.Price,
	}
	return &result, nil
}

func (s *ProductService) GetProductByID(ctx context.Context, productId uuid.UUID) (*dto.ProductResponse, error) {
	product, err := s.productRepository.FindByID(ctx, productId)
	if err != nil {
		log.WithFields(log.Fields{
			"layer":     "service",
			"function":  "GetProductByID",
			"productId": productId,
		}).WithError(err).Error("unable to get product by id")
		return nil, err
	}

	result := dto.ProductResponse{
		ID:       product.ID,
		Name:     product.Name,
		Quantity: product.Quantity,
		Price:    product.Price,
	}

	return &result, nil
}

func (s *ProductService) GetListProduct(ctx context.Context) ([]dto.ProductResponse, error) {
	var result []dto.ProductResponse
	products, err := s.productRepository.FindAll(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"layer":    "service",
			"function": "GetProductAll",
		}).WithError(err).Error("unable to get product all")
		return nil, err
	}

	if len(products) > 0 {
		for _, product := range products {
			data := dto.ProductResponse{
				ID:       product.ID,
				Name:     product.Name,
				Quantity: product.Quantity,
				Price:    product.Price,
			}
			result = append(result, data)
		}
	}

	return result, nil
}
