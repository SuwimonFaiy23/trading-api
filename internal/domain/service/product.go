package service

import (
	"context"
	"trading-api/internal/dto"

	uuid "github.com/tentone/mssql-uuid"
)

type ProductService interface {
	CreateProduct(ctx context.Context, reqBody dto.ProductRequest) (*dto.ProductResponse, error)
	GetProductByID(ctx context.Context, productId uuid.UUID) (*dto.ProductResponse, error)
	GetListProduct(ctx context.Context) ([]dto.ProductResponse, error)
}
