package dto

import uuid "github.com/tentone/mssql-uuid"

type CommissionResponse struct {
	ID        uuid.UUID        `json:"id"`
	Amount    float64          `json:"amount"`
	Order     OrderDetail      `json:"order"`
	Affiliate *AffiliateDetail `json:"affiliate"`
}

type OrderDetail struct {
	ID              uuid.UUID       `json:"id"`
	UserID          uuid.UUID       `json:"user_id"`
	Products        []ProductDetail `json:"products"`
	TotalAmount     float64         `json:"total_amount"`
	TotalCommission float64         `json:"total_commission"`
}

type AffiliateDetail struct {
	ID              uuid.UUID  `json:"id"`
	Balance         float64    `json:"balance"`
	MasterAffiliate *uuid.UUID `json:"master_affiliate"`
}
