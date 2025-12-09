package repository

import (
	"context"
	"errors"
	"trading-api/internal/domain/entity"

	uuid "github.com/tentone/mssql-uuid"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

type ProductDependencies struct {
	Database *gorm.DB
}

func NewProductRepository(d ProductDependencies) *ProductRepository {
	f := &ProductRepository{
		db: d.Database,
	}
	return f
}

func (r *ProductRepository) Create(ctx context.Context, ent *entity.Product) error {
	// 1) check duplicate username
	var count int64
	if err := r.db.WithContext(ctx).
		Model(&entity.Product{}).
		Where("name = ?", ent.Name).
		Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return errors.New("name product already exists")
	}

	// 2) create new product
	if err := r.db.WithContext(ctx).Create(ent).Error; err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) Update(ctx context.Context, ent *entity.Product) (*entity.Product, error) {
	tx := r.db.WithContext(ctx).Model(&entity.Product{}).Where("id = ?", ent.ID).Updates(ent)
	if tx.Error != nil {
		return ent, tx.Error
	}
	if tx.RowsAffected == 0 {
		return ent, gorm.ErrRecordNotFound
	}
	return ent, nil
}

func (r *ProductRepository) FindByID(ctx context.Context, productId uuid.UUID) (*entity.Product, error) {
	var product *entity.Product
	err := r.db.WithContext(ctx).Where("id = ?", productId).First(&product).Error
	return product, err
}

func (r *ProductRepository) FindAll(ctx context.Context) ([]entity.Product, error) {
	var product []entity.Product
	err := r.db.WithContext(ctx).Find(&product).Error
	return product, err
}
