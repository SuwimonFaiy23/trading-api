package repository

import (
	"context"
	"trading-api/internal/domain/entity"

	uuid "github.com/tentone/mssql-uuid"
	"gorm.io/gorm"
)

type AffiliateRepository struct {
	db *gorm.DB
}

type AffiliateDependencies struct {
	Database *gorm.DB
}

func NewAffiliateRepository(d AffiliateDependencies) *AffiliateRepository {
	f := &AffiliateRepository{
		db: d.Database,
	}
	return f
}

func (r *AffiliateRepository) Create(ctx context.Context, ent *entity.Affiliate) error {
	if err := r.db.WithContext(ctx).Create(ent).Error; err != nil {
		return err
	}
	return nil
}

func (r *AffiliateRepository) Update(ctx context.Context, ent *entity.Affiliate) (*entity.Affiliate, error) {
	tx := r.db.WithContext(ctx).Model(&entity.Affiliate{}).Where("id = ?", ent.ID).Updates(ent)
	if tx.Error != nil {
		return ent, tx.Error
	}
	if tx.RowsAffected == 0 {
		return ent, gorm.ErrRecordNotFound
	}
	return ent, nil
}

func (r *AffiliateRepository) FindByID(ctx context.Context, affiliateId uuid.UUID) (*entity.Affiliate, error) {
	var affiliate *entity.Affiliate
	err := r.db.WithContext(ctx).Where("id = ?", affiliateId).First(&affiliate).Error
	return affiliate, err
}

func (r *AffiliateRepository) FindAll(ctx context.Context) ([]entity.Affiliate, error) {
	var affiliate []entity.Affiliate
	err := r.db.WithContext(ctx).Find(&affiliate).Error
	return affiliate, err
}