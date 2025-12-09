package repository

import (
	"context"
	"trading-api/internal/domain/entity"

	uuid "github.com/tentone/mssql-uuid"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

type OrderDependencies struct {
	Database *gorm.DB
}

func NewOrderRepository(d OrderDependencies) *OrderRepository {
	f := &OrderRepository{
		db: d.Database,
	}
	return f
}

func (r *OrderRepository) Create(ctx context.Context, ent *entity.Order) error {
	if err := r.db.WithContext(ctx).Create(ent).Error; err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) Update(ctx context.Context, ent *entity.Order) (*entity.Order, error) {
	tx := r.db.WithContext(ctx).Model(&entity.Order{}).Where("id = ?", ent.ID).Updates(ent)
	if tx.Error != nil {
		return ent, tx.Error
	}
	if tx.RowsAffected == 0 {
		return ent, gorm.ErrRecordNotFound
	}
	return ent, nil
}

func (r *OrderRepository) FindByID(ctx context.Context, orderId uuid.UUID) (*entity.Order, error) {
	var order *entity.Order
	err := r.db.WithContext(ctx).Preload("User").Where("id = ?", orderId).First(&order).Error
	return order, err
}

func (r *OrderRepository) DB() *gorm.DB {
    return r.db
}