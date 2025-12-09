package repository

import (
	"context"
	"trading-api/internal/domain/entity"

	uuid "github.com/tentone/mssql-uuid"
)

type CommissionRepository interface {
	Create(ctx context.Context, ent *entity.Commission) error
	Update(ctx context.Context, ent *entity.Commission) (*entity.Commission, error)
	FindByID(ctx context.Context, commissionId uuid.UUID) (*entity.Commission, error)
	FindAll(ctx context.Context) ([]entity.Commission, error)
}
