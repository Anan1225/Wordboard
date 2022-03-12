package service

import (
	"context"

	"github.com/Anan1225/wordboard/account/model"
	"github.com/google/uuid"
)

// UserService acts as a struct for injecting an implementation of UserRepository
// for use in service methods
type UserService struct {
	UserRepository model.UserRepository
}

// USConfig will hold repositories taht will eventually be injected into
// this service layer
type USConfig struct {
	UserRepository model.UserRepository
}

// NewUserService is a factory function for
// initializing a UserService with its repository layer dependence
func NewUserService(c *USConfig) model.UserService {
	return &UserService{
		UserRepository: c.UserRepository,
	}
}

// Get retrieves a user based on their uuid
func (s *UserService) Get(ctx context.Context, uid uuid.UUID) (*model.User, error) {
	u, err := s.UserRepository.FindByID(ctx, uid)

	return u, err
}
