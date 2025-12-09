package service

import (
	"context"
	"trading-api/internal/dto"

	uuid "github.com/tentone/mssql-uuid"
)

type AffiliateService interface {
	CreateAffiliate(ctx context.Context, reqBody dto.AffiliateRequest) (*dto.AffiliateResponse, error)
	GetAffiliateByID(ctx context.Context, affiliateId uuid.UUID) (*dto.AffiliateResponse, error)
	GetListAffiliate(ctx context.Context) ([]dto.AffiliateResponse, error)
}
