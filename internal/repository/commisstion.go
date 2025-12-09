package repository

import (
	"context"
	"trading-api/internal/domain/entity"

	uuid "github.com/tentone/mssql-uuid"
	"gorm.io/gorm"
)

type CommissionRepository struct {
	db *gorm.DB
}

type CommissionDependencies struct {
	Database *gorm.DB
}

func NewCommissionRepository(d CommissionDependencies) *CommissionRepository {
	f := &CommissionRepository{
		db: d.Database,
	}
	return f
}

func (r *CommissionRepository) Create(ctx context.Context, ent *entity.Commission) error {
	if err := r.db.WithContext(ctx).Create(ent).Error; err != nil {
		return err
	}
	return nil
}

func (r *CommissionRepository) Update(ctx context.Context, ent *entity.Commission) (*entity.Commission, error) {
	tx := r.db.WithContext(ctx).Model(&entity.Commission{}).Where("id = ?", ent.ID).Updates(ent)
	if tx.Error != nil {
		return ent, tx.Error
	}
	if tx.RowsAffected == 0 {
		return ent, gorm.ErrRecordNotFound
	}
	return ent, nil
}

func (r *CommissionRepository) FindByID(ctx context.Context, commissionId uuid.UUID) (*entity.Commission, error) {
	var commission *entity.Commission
	err := r.db.WithContext(ctx).Preload("Order").Preload("Affiliate").Where("id = ?", commissionId).First(&commission).Error
	return commission, err
}

func (r *CommissionRepository) FindAll(ctx context.Context) ([]entity.Commission, error) {
	var commission []entity.Commission
	err := r.db.WithContext(ctx).Preload("Order").Preload("Affiliate").Find(&commission).Error
	return commission, err
}
