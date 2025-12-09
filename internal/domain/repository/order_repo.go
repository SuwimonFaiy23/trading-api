package repository

import (
	"context"
	"trading-api/internal/domain/entity"

	uuid "github.com/tentone/mssql-uuid"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(ctx context.Context, ent *entity.Order) error
	Update(ctx context.Context, ent *entity.Order) (*entity.Order, error)
	FindByID(ctx context.Context, orderId uuid.UUID) (*entity.Order, error)
	DB() *gorm.DB
}
