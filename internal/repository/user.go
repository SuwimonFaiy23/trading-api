package repository

import (
	"context"
	"errors"
	"trading-api/internal/domain/entity"

	uuid "github.com/tentone/mssql-uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository struct {
	db *gorm.DB
}

type UserDependencies struct {
	Database *gorm.DB
}

func NewUserRepository(d UserDependencies) *UserRepository {
	f := &UserRepository{
		db: d.Database,
	}
	return f
}

func (r *UserRepository) Create(ctx context.Context, ent *entity.User) error {
	// 1) check duplicate username
	var count int64
	if err := r.db.WithContext(ctx).
		Model(&entity.User{}).
		Where("username = ?", ent.Username).
		Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return errors.New("username already exists")
	}

	// 2) create new user
	if err := r.db.WithContext(ctx).Create(ent).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Update(ctx context.Context, ent *entity.User) (*entity.User, error) {
	tx := r.db.WithContext(ctx).Model(&entity.User{}).Where("id = ?", ent.ID).Updates(ent)
	if tx.Error != nil {
		return ent, tx.Error
	}
	if tx.RowsAffected == 0 {
		return ent, gorm.ErrRecordNotFound
	}
	return ent, nil
}

func (r *UserRepository) FindByID(ctx context.Context, userId uuid.UUID) (*entity.User, error) {
	var user *entity.User
	err := r.db.WithContext(ctx).Preload("Affiliate").Where("id = ?", userId).First(&user).Error
	return user, err
}

func (r *UserRepository) DeductBalanceTx(ctx context.Context, userID uuid.UUID, amount float64) (*entity.User, error) {
	var user entity.User

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 1) lock row
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&user, "id = ?", userID).Error; err != nil {
			return err
		}

		// 2) check balance
		if user.Balance < amount {
			return errors.New("insufficient balance")
		}

		// 3) deduct
		user.Balance -= amount

		// 4) update
		if err := tx.Save(&user).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &user, nil
}


func (r *UserRepository) AddBalanceTx(ctx context.Context, userID uuid.UUID, amount float64) (*entity.User, error) {
    if amount <= 0 {
        return nil, errors.New("amount must be positive")
    }

    var user entity.User

    err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
        // 1) Lock row
        if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
            First(&user, "id = ?", userID).Error; err != nil {
            return err
        }

        // 2) เพิ่ม balance
        user.Balance += amount

        // 3) Update DB
        if err := tx.Save(&user).Error; err != nil {
            return err
        }

        return nil
    })

    if err != nil {
        return nil, err
    }

    return &user, nil
}

func (r *UserRepository) GetAllUsersByPagination(ctx context.Context, limit, offset int) ([]entity.User, int64, error) {
    var users []entity.User
    var total int64

    // นับทั้งหมด
    if err := r.db.WithContext(ctx).Model(&entity.User{}).Count(&total).Error; err != nil {
        return nil, 0, err
    }

    // ดึงข้อมูลตาม limit + offset
    if err := r.db.WithContext(ctx).
        Limit(limit).
        Offset(offset).
        Order("created_at DESC").
        Find(&users).Error; err != nil {
        return nil, 0, err
    }

    return users, total, nil
}