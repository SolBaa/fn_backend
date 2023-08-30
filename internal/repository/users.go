package repository

import (
	"github.com/SolBaa/finances-backend/internal/models"
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	GetUserByID(id int) (*models.User, error)
	CreateUser(user *models.User) error
}

// Path: internal/repository/users.go

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (ur *UserRepositoryImpl) GetUserByID(id int) (*models.User, error) {
	// Implementación para obtener un usuario por ID
	return nil, nil

}

func (ur *UserRepositoryImpl) CreateUser(user *models.User) error {
	// Implementación para crear un nuevo usuario
	return nil
}
