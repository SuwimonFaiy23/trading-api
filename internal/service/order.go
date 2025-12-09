package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"trading-api/internal/domain/entity"
	"trading-api/internal/domain/repository"
	"trading-api/internal/dto"

	log "github.com/sirupsen/logrus"
	uuid "github.com/tentone/mssql-uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type OrderService struct {
	orderRepository   repository.OrderRepository
	userRepository    repository.UserRepository
	productRepository repository.ProductRepository
}

type OrderServiceDependencies struct {
	OrderRepository   repository.OrderRepository
	UserRepository    repository.UserRepository
	ProductRepository repository.ProductRepository
}

func NewOrderService(d OrderServiceDependencies) *OrderService {
	return &OrderService{
		orderRepository:   d.OrderRepository,
		userRepository:    d.UserRepository,
		productRepository: d.ProductRepository,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, reqBody dto.OrderRequest) (*dto.OrderResponse, error) {
	var result dto.OrderResponse
	// start transaction
	err := s.orderRepository.DB().Transaction(func(tx *gorm.DB) error {
		// 1. validate product
		if len(reqBody.Products) == 0 {
			return errors.New("at least one product is required")
		}

		for _, product := range reqBody.Products {
			_, err := s.productRepository.FindByID(ctx, product.ProductID)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return fmt.Errorf("product not found: %v", product.ProductID.String())
				}
				log.WithFields(log.Fields{
					"layer":     "service",
					"function":  "CreateOrder",
					"productId": product.ProductID,
				}).WithError(err).Error("unable to get product by id")
				return err
			}
		}

		// 2. convert struct → []byte
		jsonBytes, err := json.Marshal(reqBody.Products)
		if err != nil {
			log.WithFields(log.Fields{
				"layer":    "service",
				"function": "CreateOrder",
				"products": reqBody.Products,
			}).WithError(err).Error("unable to marshal product list")
			return errors.New("unable to process product list")
		}

		// 3. get user by user id
		user, err := s.userRepository.FindByID(ctx, reqBody.UserID)
		if err != nil {
			log.WithFields(log.Fields{
				"layer":    "service",
				"function": "CreateOrder",
				"userId":   reqBody.UserID,
			}).WithError(err).Error("unable to get user by user id")
			return err
		}

		// 4. calculate total
		var totalAmount float64 = 0.0
		for _, product := range reqBody.Products {
			totalAmount += (product.Price * float64(product.Quantity))
		}

		commissionRate := 0.10
		commissionAmount := totalAmount * commissionRate

		// no affiliate → no commission
		if user.AffiliateID == nil {
			commissionAmount = 0
		}

		// 5. create order
		order := &entity.Order{
			ID:              uuid.NewV4(),
			UserID:          reqBody.UserID,
			ProductDetail:   datatypes.JSON(jsonBytes),
			TotalAmount:     totalAmount,
			TotalCommission: commissionAmount,
		}
		if err := tx.Create(order).Error; err != nil {
			log.WithFields(log.Fields{
				"layer":    "service",
				"function": "CreateOrder",
				"reqBody":  reqBody,
			}).WithError(err).Error("unable to create order")
			return err
		}

		// 6. create commission only if affiliate exists
		if user.AffiliateID != nil {
			commission := &entity.Commission{
				ID:          uuid.NewV4(),
				OrderID:     order.ID,
				AffiliateID: user.AffiliateID,
				Amount:      commissionAmount,
			}

			if err := tx.Create(commission).Error; err != nil {
				log.WithFields(log.Fields{
					"layer":    "service",
					"function": "CreateOrder",
					"reqBody":  reqBody,
				}).WithError(err).Error("unable to create commission")
				return err
			}
		}

		//7. update product
		for _, v := range reqBody.Products {
			product, err := s.productRepository.FindByID(ctx, v.ProductID)
			if err != nil {
				log.WithFields(log.Fields{
					"layer":     "service",
					"function":  "CreateOrder",
					"productId": v.ProductID,
				}).WithError(err).Error("unable to get product by id")
				return err
			}
			product.Quantity -= v.Quantity
			if err := tx.Save(product).Error; err != nil {
				log.WithFields(log.Fields{
					"layer":    "service",
					"function": "CreateOrder",
					"product":  product,
				}).WithError(err).Error("unable to update product")
				return err
			}
		}

		// 8. update balance user
		user.Balance -= totalAmount
		if user.Balance < 0 {
			return errors.New("insufficient balance")
		}
		if err := tx.Save(user).Error; err != nil {
			log.WithFields(log.Fields{
				"layer":    "service",
				"function": "CreateOrder",
				"user":     user,
			}).WithError(err).Error("unable to update user")
			return err
		}

		// 9. response
		result = dto.OrderResponse{
			ID:              order.ID,
			TotalAmount:     order.TotalAmount,
			TotalCommission: order.TotalCommission,
			UserID:          order.UserID,
			Products:        reqBody.Products,
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *OrderService) GetOrderByID(ctx context.Context, orderId uuid.UUID) (*dto.OrderResponse, error) {
	var result dto.OrderResponse

	order, err := s.orderRepository.FindByID(ctx, orderId)
	if err != nil {
		log.WithFields(log.Fields{
			"layer":    "service",
			"function": "GetOrderByID",
			"orderId":  orderId,
		}).WithError(err).Error("unable to get order by id")
		return nil, err
	}

	if order != nil {
		var detail []dto.ProductDetail
		// order.ProductDetail is datatypes.JSON → []byte
		if err := json.Unmarshal(order.ProductDetail, &detail); err != nil {
			log.Println("unmarshal error:", err)
			return nil, err
		}

		result = dto.OrderResponse{
			ID:              order.ID,
			TotalAmount:     order.TotalAmount,
			TotalCommission: order.TotalCommission,
			UserID:          order.UserID,
			Products:        detail,
		}
	}
	return &result, nil
}
