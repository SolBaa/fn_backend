package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Ingredient struct {
	ingredientService service.IngredientService
}

func NewIngredient(c service.IngredientService) *Ingredient {
	return &Ingredient{
		ingredientService: c,
	}
}

func (i *Ingredient) GetIngredient(w http.ResponseWriter, r *http.Request) {

	if chi.URLParam(r, "id") == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ingredient, err := i.ingredientService.GetIngredient(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	ing, err := json.Marshal(ingredient)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(ing)

}
