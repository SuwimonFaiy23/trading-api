package service

import (
	"context"
	"errors"
	"math"
	"trading-api/internal/domain/entity"
	"trading-api/internal/domain/repository"
	"trading-api/internal/dto"

	log "github.com/sirupsen/logrus"
	uuid "github.com/tentone/mssql-uuid"
)

type UserService struct {
	userRepository repository.UserRepository
}

type UserServiceDependencies struct {
	UserRepository repository.UserRepository
}

func NewUserService(d UserServiceDependencies) *UserService {
	return &UserService{
		userRepository: d.UserRepository,
	}
}

func (s *UserService) CreateUser(ctx context.Context, reqBody dto.UserRequest) (*dto.UserResponse, error) {
	user := &entity.User{
		ID:       uuid.NewV4(),
		Username: reqBody.Username,
		Balance:  0,
	}
	if err := s.userRepository.Create(ctx, user); err != nil {
		log.WithFields(log.Fields{
			"layer":    "service",
			"function": "CreateUser",
			"reqBody":  reqBody,
		}).WithError(err).Error("unable to create user")
		return nil, err
	}
	result := dto.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Balance:  user.Balance,
	}
	return &result, nil
}

func (s *UserService) UpdateUser(ctx context.Context, userId uuid.UUID, reqBody dto.UserRequest) (*dto.UserResponse, error) {
	user, err := s.userRepository.FindByID(ctx, userId)
	if err != nil {
		log.WithFields(log.Fields{
			"layer":    "service",
			"function": "UpdateUser",
			"userId":   userId,
		}).WithError(err).Error("unable to get user by id")
		return nil, err
	}

	user.Username = reqBody.Username
	response, err := s.userRepository.Update(ctx, user)
	if err != nil {
		log.WithFields(log.Fields{
			"layer":    "service",
			"function": "UpdateUser",
			"reqBody":  reqBody,
			"userId":   userId,
		}).WithError(err).Error("unable to update user")
		return nil, err
	}

	result := dto.UserResponse{
		ID:       response.ID,
		Username: response.Username,
		Balance:  response.Balance,
	}

	return &result, nil
}

func (s *UserService) GetUserByID(ctx context.Context, userId uuid.UUID) (*dto.UserResponse, error) {
	user, err := s.userRepository.FindByID(ctx, userId)
	if err != nil {
		log.WithFields(log.Fields{
			"layer":    "service",
			"function": "GetUserByID",
			"userId":   userId,
		}).WithError(err).Error("unable to get user by id")
		return nil, err
	}

	result := dto.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Balance:  user.Balance,
	}

	return &result, nil
}

func (s *UserService) AddBalanceUser(ctx context.Context, userId uuid.UUID, reqBody dto.BalanceRequest) (*dto.BalanceResponse, error) {
	if reqBody.Amount <= 0 {
		return nil, errors.New("amount must be positive")
	}

	response, err := s.userRepository.AddBalanceTx(ctx, userId, reqBody.Amount)
	if err != nil {
		log.WithFields(log.Fields{
			"layer":    "service",
			"function": "AddBalanceUser",
			"reqBody":  reqBody,
			"userId":   userId,
		}).WithError(err).Error("unable to add balance user")
		return nil, err
	}

	result := dto.BalanceResponse{
		UserID:  response.ID,
		Balance: response.Balance,
	}

	return &result, nil
}

func (s *UserService) DeductBalanceUser(ctx context.Context, userId uuid.UUID, reqBody dto.BalanceRequest) (*dto.BalanceResponse, error) {
	if reqBody.Amount <= 0 {
		return nil, errors.New("amount must be positive")
	}
	response, err := s.userRepository.DeductBalanceTx(ctx, userId, reqBody.Amount)
	if err != nil {
		log.WithFields(log.Fields{
			"layer":    "service",
			"function": "DeductBalanceUser",
			"reqBody":  reqBody,
			"userId":   userId,
		}).WithError(err).Error("unable to deduct balance user")
		return nil, err
	}

	result := dto.BalanceResponse{
		UserID:  response.ID,
		Balance: response.Balance,
	}

	return &result, nil
}

func (s *UserService) GetUserAllByPagination(ctx context.Context, limit, page int) (*dto.UserListResponse, error) {
	if limit <= 0 {
		limit = 10
	}
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * limit

	users, total, err := s.userRepository.GetAllUsersByPagination(ctx, limit, offset)
	if err != nil {
		log.WithFields(log.Fields{
			"layer":    "service",
			"function": "GetUserAllByPagination",
			"limit":    limit,
			"page":     page,
		}).WithError(err).Error("unable to get users paginated")
		return nil, err
	}
	totalPage := int(math.Ceil(float64(total) / float64(limit)))
	count := len(users)
	var data []dto.UserResponse
	if count > 0 {
		for _, user := range users {
			user := dto.UserResponse{
				ID:       user.ID,
				Username: user.Username,
				Balance:  user.Balance,
			}
			data = append(data, user)
		}
	}

	response := dto.UserListResponse{
		Page:       page,
		TotalPage:  totalPage,
		Count:      count,
		TotalCount: total,
		Data:       data,
	}

	return &response, nil
}
