package repository

import (
	"context"
	"trading-api/internal/domain/entity"

	uuid "github.com/tentone/mssql-uuid"
)

type UserRepository interface {
	Create(ctx context.Context, ent *entity.User) error
	Update(ctx context.Context, ent *entity.User) (*entity.User, error)
	FindByID(ctx context.Context, userId uuid.UUID) (*entity.User, error)
	DeductBalanceTx(ctx context.Context, userID uuid.UUID, amount float64) (*entity.User, error)
	AddBalanceTx(ctx context.Context, userID uuid.UUID, amount float64) (*entity.User, error)
	GetAllUsersByPagination(ctx context.Context, limit, offset int) ([]entity.User, int64, error)
}
