package service

import (
	"context"
	"trading-api/internal/dto"

	uuid "github.com/tentone/mssql-uuid"
)

type CommissionService interface {
	GetCommissionByID(ctx context.Context, commissionId uuid.UUID) (*dto.CommissionResponse, error)
	GetListCommission(ctx context.Context) ([]dto.CommissionResponse, error)
}
