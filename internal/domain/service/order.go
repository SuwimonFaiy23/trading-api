package service

import (
	"context"
	"trading-api/internal/dto"

	uuid "github.com/tentone/mssql-uuid"
)

type OrderService interface {
	CreateOrder(ctx context.Context, reqBody dto.OrderRequest) (*dto.OrderResponse, error)
	GetOrderByID(ctx context.Context, orderId uuid.UUID) (*dto.OrderResponse, error)
}
