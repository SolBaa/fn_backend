package repository

import (
	"fmt"

	"github.com/SolBaa/recipes-backend/internal/models"
	"gorm.io/gorm"
)

type RecipeRepository interface {
	GetRecipeByID(id int) (models.Recipe, error)
	GetRecipes() ([]models.Recipe, error)
}

type recipeRepository struct {
	db *gorm.DB
}

func NewRecipeRepository(db *gorm.DB) RecipeRepository {
	return &recipeRepository{db: db}
}

func (r *recipeRepository) GetRecipeByID(id int) (models.Recipe, error) {
	var recipe models.Recipe
	result := r.db.First(&recipe, id)
	if result.Error != nil {
		return models.Recipe{}, result.Error
	}
	return recipe, nil
}

func (r *recipeRepository) GetRecipes() ([]models.Recipe, error) {
	var recipe []models.Recipe
	result := r.db.Find(&recipe)
	if result.Error != nil {
		return []models.Recipe{}, result.Error
	}
	fmt.Println(recipe)
	return recipe, nil
}
