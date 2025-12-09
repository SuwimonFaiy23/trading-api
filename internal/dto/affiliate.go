package dto

import uuid "github.com/tentone/mssql-uuid"

type AffiliateRequest struct {
	UserID   uuid.UUID  `json:"user_id" binding:"required"`
	Name     string     `json:"name" binding:"required"`
	MasterID *uuid.UUID `json:"master_id"`
}

type AffiliateResponse struct {
	ID       uuid.UUID  `json:"id"`
	Name     string     `json:"name"`
	MasterID *uuid.UUID `json:"master_id"`
	Balance  float64    `json:"balance"`
}
