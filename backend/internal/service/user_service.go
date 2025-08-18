package service

import (
	"context"
	"go-socket-scrum-poker/internal/domain"
)

type (
	UserService struct {
		*Service
	}

	UserServiceInterface interface {
		CreateUser(ctx context.Context, req *domain.CreateUserRequest) (*domain.CreateUserResponse, error)
	}
)

func NewUserService(s *Service) *UserService {
	return &UserService{
		Service: s,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *domain.CreateUserRequest) (*domain.CreateUserResponse, error) {

	return &domain.CreateUserResponse{}, nil
}
