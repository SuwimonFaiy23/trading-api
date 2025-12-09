package entity

import uuid "github.com/tentone/mssql-uuid"

const TableNameCommission = "commission"

type Commission struct {
	ID          uuid.UUID  `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	OrderID     uuid.UUID  `gorm:"column:order_id;type:uuid;not null" json:"order_id"`
	AffiliateID *uuid.UUID `gorm:"column:affiliate_id;type:uuid; null" json:"affiliate_id"`
	Amount      float64    `gorm:"column:amount;type:double precision;default:0" json:"amount"`
	BaseModel

	// Relations
	Order     Order     `gorm:"foreignKey:OrderID"`
	Affiliate Affiliate `gorm:"foreignKey:AffiliateID"`
}

// TableName TableNameCommission's table name
func (*Commission) TableName() string {
	return TableNameCommission
}
