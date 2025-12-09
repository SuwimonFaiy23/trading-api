package service

import (
	"context"
	"errors"
	"fmt"
	"trading-api/internal/domain/entity"
	"trading-api/internal/domain/repository"
	"trading-api/internal/dto"

	log "github.com/sirupsen/logrus"
	uuid "github.com/tentone/mssql-uuid"
	"gorm.io/gorm"
)

type AffiliateService struct {
	affiliateRepository repository.AffiliateRepository
	userRepository      repository.UserRepository
}

type AffiliateServiceDependencies struct {
	AffiliateRepository repository.AffiliateRepository
	UserRepository      repository.UserRepository
}

func NewAffiliateService(d AffiliateServiceDependencies) *AffiliateService {
	return &AffiliateService{
		affiliateRepository: d.AffiliateRepository,
		userRepository:      d.UserRepository,
	}
}

func (s *AffiliateService) CreateAffiliate(ctx context.Context, reqBody dto.AffiliateRequest) (*dto.AffiliateResponse, error) {
	user, err := s.userRepository.FindByID(ctx, reqBody.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found: %v", reqBody.UserID.String())
		}
		log.WithFields(log.Fields{
			"layer":    "service",
			"function": "CreateAffiliate",
			"userId":   reqBody.UserID,
		}).WithError(err).Error("unable to get user by id")
		return nil, err
	}

	affiliate := &entity.Affiliate{
		ID:              uuid.NewV4(),
		Name:            reqBody.Name,
		MasterAffiliate: reqBody.MasterID,
		Balance:         0,
	}
	if err := s.affiliateRepository.Create(ctx, affiliate); err != nil {
		log.WithFields(log.Fields{
			"layer":    "service",
			"function": "CreateAffiliate",
			"reqBody":  reqBody,
		}).WithError(err).Error("unable to create affiliate")
		return nil, err
	}

	user.AffiliateID = &affiliate.ID
	_, err = s.userRepository.Update(ctx, user)
	if err != nil {
		log.WithFields(log.Fields{
			"layer":    "service",
			"function": "CreateAffiliate",
			"entity":   user,
		}).WithError(err).Error("unable to update user")
		return nil, err
	}

	result := dto.AffiliateResponse{
		ID:       affiliate.ID,
		Name:     affiliate.Name,
		MasterID: affiliate.MasterAffiliate,
		Balance:  affiliate.Balance,
	}

	return &result, nil
}

func (s *AffiliateService) GetAffiliateByID(ctx context.Context, affiliateId uuid.UUID) (*dto.AffiliateResponse, error) {
	var result dto.AffiliateResponse
	affiliate, err := s.affiliateRepository.FindByID(ctx, affiliateId)
	if err != nil {
		log.WithFields(log.Fields{
			"layer":       "service",
			"function":    "GetAffiliateByID",
			"affiliateId": affiliateId,
		}).WithError(err).Error("unable to get affiliate by id")
		return nil, err
	}

	if affiliate != nil {
		result = dto.AffiliateResponse{
			ID:       affiliate.ID,
			Name:     affiliate.Name,
			MasterID: affiliate.MasterAffiliate,
			Balance:  affiliate.Balance,
		}
	}

	return &result, nil
}

func (s *AffiliateService) GetListAffiliate(ctx context.Context) ([]dto.AffiliateResponse, error) {
	var result []dto.AffiliateResponse
	affiliates, err := s.affiliateRepository.FindAll(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"layer":    "service",
			"function": "GetListAffiliate",
		}).WithError(err).Error("unable to get all affiliate")
		return nil, err
	}

	if len(affiliates) > 0 {
		for _, affiliate := range affiliates {
			result = append(result, dto.AffiliateResponse{
				ID:       affiliate.ID,
				Name:     affiliate.Name,
				MasterID: affiliate.MasterAffiliate,
				Balance:  affiliate.Balance,
			})
		}
	}

	return result, nil
}
