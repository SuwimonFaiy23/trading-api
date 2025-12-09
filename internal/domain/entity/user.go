package entity

import uuid "github.com/tentone/mssql-uuid"

const TableNameUser = "user"

type User struct {
	ID          uuid.UUID  `gorm:"column:id;type:uuid;default:uuid_generate_v4()" json:"id"`
	Username    string     `gorm:"column:username;type:nvarchar(50);not null;" json:"username"`
	Balance     float64    `gorm:"column:balance;type:double;default:0" json:"balance"`
	AffiliateID *uuid.UUID `gorm:"column:affiliate_id;type:uuid" json:"affiliate_id"`
	BaseModel

	// Relations
	Affiliate *Affiliate `gorm:"foreignKey:AffiliateID"`
	Orders    []Order    `gorm:"foreignKey:UserID"`
}

// TableName TableNameUser's table name
func (*User) TableName() string {
	return TableNameUser
}
