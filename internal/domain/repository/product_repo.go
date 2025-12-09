package repository

import (
	"context"
	"trading-api/internal/domain/entity"

	uuid "github.com/tentone/mssql-uuid"
)

type ProductRepository interface {
	Create(ctx context.Context, ent *entity.Product) error
	Update(ctx context.Context, ent *entity.Product) (*entity.Product, error)
	FindByID(ctx context.Context, productId uuid.UUID) (*entity.Product, error)
	FindAll(ctx context.Context) ([]entity.Product, error)
}
