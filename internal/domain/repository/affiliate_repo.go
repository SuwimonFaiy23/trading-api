package repository

import (
	"context"
	"trading-api/internal/domain/entity"

	uuid "github.com/tentone/mssql-uuid"
)

type AffiliateRepository interface {
	Create(ctx context.Context, ent *entity.Affiliate) error
	Update(ctx context.Context, ent *entity.Affiliate) (*entity.Affiliate, error)
	FindByID(ctx context.Context, affiliateId uuid.UUID) (*entity.Affiliate, error)
	FindAll(ctx context.Context) ([]entity.Affiliate, error)
}
