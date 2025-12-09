package dto

import (
	uuid "github.com/tentone/mssql-uuid"
)

type UserRequest struct {
	Username string `json:"username" binding:"required"`
}

type UserResponse struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Balance  float64   `json:"balance"`
}

type BalanceRequest struct {
	Amount float64 `json:"amount"  binding:"required"`
}

type BalanceResponse struct {
	UserID  uuid.UUID `json:"user_id"`
	Balance float64   `json:"balance"`
}

type UserListResponse struct {
	Data       []UserResponse `json:"data"`
	Page       int            `json:"page"`
	TotalPage  int            `json:"total_page"`
	Count      int            `json:"count"`
	TotalCount int64          `json:"total_count"`
}
