package handler

import "github.com/SolBaa/recipes-backend/internal/service"

type Recipe struct {
	recipeService service.RecipeService
}

func NewRecipe(c service.RecipeService) *Recipe {
	return &Recipe{
		recipeService: c,
	}
}
