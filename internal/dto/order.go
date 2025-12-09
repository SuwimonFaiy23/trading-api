package dto

import (
	uuid "github.com/tentone/mssql-uuid"
)

type OrderRequest struct {
	UserID   uuid.UUID       `json:"user_id" binding:"required"`
	Products []ProductDetail `json:"products" binding:"required"`
}

type ProductDetail struct {
	ProductID uuid.UUID `json:"product_id"`
	Name      string    `json:"name"`
	Quantity  int64     `json:"quantity"`
	Price     float64   `json:"price"`
}

type OrderResponse struct {
	ID              uuid.UUID       `json:"id"`
	UserID          uuid.UUID       `json:"user_id"`
	Products        []ProductDetail `json:"products"`
	TotalAmount     float64         `json:"total_amount"`
	TotalCommission float64         `json:"total_commission"`
}
