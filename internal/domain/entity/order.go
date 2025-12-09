package entity

import (
	"time"

	uuid "github.com/tentone/mssql-uuid"
	"gorm.io/datatypes"
)

const TableNameOrder = "order"

type Order struct {
	ID              uuid.UUID      `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID          uuid.UUID      `gorm:"column:user_id;type:uuid;not null" json:"user_id"`
	ProductDetail   datatypes.JSON `gorm:"column:product_detail;type:jsonb;not null" json:"product_detail"`
	TotalAmount     float64        `gorm:"column:total_amount;type:double precision;default:0" json:"total_amount"`
	TotalCommission float64        `gorm:"column:total_commission;type:double precision;default:0" json:"total_commission"`
	CreatedAt       *time.Time     `gorm:"column:created_at;autoCreateTime" json:"created_at"`

	// Relations
	User        User         `gorm:"foreignKey:UserID"`
	Commissions []Commission `gorm:"foreignKey:OrderID"`
}

// TableName TableNameOrder's table name
func (*Order) TableName() string {
	return TableNameOrder
}
