package repository

import (
	"github.com/SolBaa/recipes-backend/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(id int) (models.User, error)
	// CreateUser(user *models.User) error
}

// Path: internal/repository/users.go

type repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &repository{db: db}
}

func (r *repository) GetUserByID(id int) (models.User, error) {
	var user models.User
	result := r.db.First(&user, id)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

// func (r *repository) CreateUser(user *models.User) error {
// 	// Implementación para crear un nuevo usuario
// 	return nil
// }
