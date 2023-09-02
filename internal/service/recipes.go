package service

import (
	"fmt"

	"github.com/SolBaa/recipes-backend/internal/models"
	"github.com/SolBaa/recipes-backend/internal/repository"
)

type recipeService struct {
	repository repository.RecipeRepository
}

type RecipeService interface {
	// CreateUser() (int, error)
	GetRecipe(id int) (models.Recipe, error)
	// GetCarryCount(ctx context.Context, id string) ([]models.User, error)
}

func NewRecipeService(r repository.RecipeRepository) RecipeService {
	return &recipeService{repository: r}
}

func (s *recipeService) GetRecipe(id int) (models.Recipe, error) {
	recipe, err := s.repository.GetRecipeByID(id)
	fmt.Printf("%+v\n", recipe)
	if err != nil {
		return models.Recipe{}, ErrNotFound
	}
	return recipe, nil
}
