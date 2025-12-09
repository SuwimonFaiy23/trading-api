package service

import (
	"context"
	"encoding/json"
	"trading-api/internal/domain/repository"
	"trading-api/internal/dto"

	log "github.com/sirupsen/logrus"
	uuid "github.com/tentone/mssql-uuid"
)

type CommissionService struct {
	commissionRepository repository.CommissionRepository
}

type CommissionServiceDependencies struct {
	CommissionRepository repository.CommissionRepository
}

func NewCommissionService(d CommissionServiceDependencies) *CommissionService {
	return &CommissionService{
		commissionRepository: d.CommissionRepository,
	}
}

func (s *CommissionService) GetCommissionByID(ctx context.Context, commissionId uuid.UUID) (*dto.CommissionResponse, error) {
	var result dto.CommissionResponse
	commission, err := s.commissionRepository.FindByID(ctx, commissionId)
	if err != nil {
		log.WithFields(log.Fields{
			"layer":        "service",
			"function":     "GetCommissionByID",
			"commissionId": commissionId,
		}).WithError(err).Error("unable to get commission by id")
		return nil, err
	}

	if commission != nil {
		var detail []dto.ProductDetail
		// order.ProductDetail is datatypes.JSON → []byte
		if err := json.Unmarshal(commission.Order.ProductDetail, &detail); err != nil {
			log.Println("unmarshal error:", err)
			return nil, err
		}
		order := dto.OrderDetail{
			ID:              commission.Order.ID,
			UserID:          commission.Order.UserID,
			TotalAmount:     commission.Order.TotalAmount,
			TotalCommission: commission.Order.TotalCommission,
			Products:        detail,
		}
		affiliate := dto.AffiliateDetail{
			ID:              commission.Affiliate.ID,
			MasterAffiliate: commission.Affiliate.MasterAffiliate,
			Balance:         commission.Affiliate.Balance,
		}

		result = dto.CommissionResponse{
			ID:        commission.ID,
			Amount:    commission.Amount,
			Order:     order,
			Affiliate: &affiliate,
		}
	}

	return &result, nil
}

func (s *CommissionService) GetListCommission(ctx context.Context) ([]dto.CommissionResponse, error) {
	var result []dto.CommissionResponse
	commissions, err := s.commissionRepository.FindAll(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"layer":    "service",
			"function": "GetListCommission",
		}).WithError(err).Error("unable to get all commission")
		return nil, err
	}

	if len(commissions) > 0 {
		for _, commission := range commissions {
			var detail []dto.ProductDetail
			// order.ProductDetail is datatypes.JSON → []byte
			if err := json.Unmarshal(commission.Order.ProductDetail, &detail); err != nil {
				log.Println("unmarshal error:", err)
				return nil, err
			}
			order := dto.OrderDetail{
				ID:              commission.Order.ID,
				UserID:          commission.Order.UserID,
				TotalAmount:     commission.Order.TotalAmount,
				TotalCommission: commission.Order.TotalCommission,
				Products:        detail,
			}
			affiliate := dto.AffiliateDetail{
				ID:              commission.Affiliate.ID,
				MasterAffiliate: commission.Affiliate.MasterAffiliate,
				Balance:         commission.Affiliate.Balance,
			}

			result = append(result, dto.CommissionResponse{
				ID:        commission.ID,
				Amount:    commission.Amount,
				Order:     order,
				Affiliate: &affiliate,
			})
		}
	}

	return result, nil
}
