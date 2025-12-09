package entity

import uuid "github.com/tentone/mssql-uuid"

const TableNameAffiliate = "affiliate"

type Affiliate struct {
	ID              uuid.UUID  `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name            string     `gorm:"column:name;type:nvarchar(255);not null;" json:"name"`
	MasterAffiliate *uuid.UUID `gorm:"column:master_affiliate;type:uuid" json:"master_affiliate"`
	Balance         float64    `gorm:"column:balance;type:double precision;default:0" json:"balance"`
	BaseModel

	// Relations
	Users       []User       `gorm:"foreignKey:AffiliateID"`
	Commissions []Commission `gorm:"foreignKey:AffiliateID"`
}

// TableName TableNameAffiliate's table name
func (*Affiliate) TableName() string {
	return TableNameAffiliate
}
