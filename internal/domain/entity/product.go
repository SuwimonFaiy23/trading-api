package entity

import uuid "github.com/tentone/mssql-uuid"

const TableNameProduct = "product"

type Product struct {
	ID       uuid.UUID `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name     string    `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Quantity int64     `gorm:"column:quantity;default:0" json:"quantity"`
	Price    float64   `gorm:"column:price;type:double precision;default:0" json:"price"`
	BaseModel
}

// TableName TableNameProduct's table name
func (*Product) TableName() string {
	return TableNameProduct
}
