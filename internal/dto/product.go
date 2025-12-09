package dto

import uuid "github.com/tentone/mssql-uuid"

type ProductRequest struct {
	Name     string  `json:"name" binding:"required"`
	Quantity int64   `json:"quantity" binding:"required"`
	Price    float64 `json:"price" binding:"required"`
}

type ProductResponse struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Quantity int64     `json:"quantity"`
	Price    float64   `json:"price"`
}
