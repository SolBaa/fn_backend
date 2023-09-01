package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/SolBaa/recipes-backend/internal/service"
	"github.com/go-chi/chi/v5"
)

type User struct {
	userService service.Service
}

func NewUser(c service.Service) *User {
	return &User{
		userService: c,
	}
}

func (u *User) GetUser(w http.ResponseWriter, r *http.Request) {

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

	user, err := u.userService.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	usr, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(usr)

}

// Nueva función para configurar rutas del controlador
func (h *User) Routes() http.Handler {
	r := chi.NewRouter()
	r.Get("/{id}", h.GetUser)
	// Agregar más rutas relacionadas con usuarios aquí
	return r
}
