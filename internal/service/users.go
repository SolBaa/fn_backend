package service

import (
	"errors"
	"fmt"

	"github.com/SolBaa/recipes-backend/internal/models"
	"github.com/SolBaa/recipes-backend/internal/repository"
)

// Errors
var (
	ErrNotFound           = errors.New("seller not found")
	ErrSave               = errors.New("seller not saved")
	ErrCIDExists          = errors.New("cid already exists")
	ErrLocalityIDNotFound = errors.New("locality not found")
	ErrSellerExists       = errors.New("seller already exists")
)

type service struct {
	repository repository.UserRepository
}

type Service interface {
	// CreateUser() (int, error)
	GetUser(id int) (models.User, error)
	// GetCarryCount(ctx context.Context, id string) ([]models.User, error)
}

func NewService(r repository.UserRepository) Service {
	return &service{repository: r}
}

func (s *service) GetUser(id int) (models.User, error) {
	user, err := s.repository.GetUserByID(id)
	fmt.Printf("%+v\n", user)
	if err != nil {
		return models.User{}, ErrNotFound
	}
	return user, nil
}
