package service

import (
	"context"
	"trading-api/internal/dto"

	uuid "github.com/tentone/mssql-uuid"
)

type UserService interface {
	CreateUser(ctx context.Context, reqBody dto.UserRequest) (*dto.UserResponse, error)
	UpdateUser(ctx context.Context, userId uuid.UUID, reqBody dto.UserRequest) (*dto.UserResponse, error)
	GetUserByID(ctx context.Context, userId uuid.UUID) (*dto.UserResponse, error)
	AddBalanceUser(ctx context.Context, userId uuid.UUID, reqBody dto.BalanceRequest) (*dto.BalanceResponse, error)
	DeductBalanceUser(ctx context.Context, userId uuid.UUID, reqBody dto.BalanceRequest) (*dto.BalanceResponse, error)
	GetUserAllByPagination(ctx context.Context, limit, page int) (*dto.UserListResponse, error)
}
