package router

import (
	"net/http"

	"github.com/SolBaa/recipes-backend/cmd/api/handler"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type App struct {
	db          *gorm.DB
	userHandler *handler.User
}

func NewRouter(db *gorm.DB, userHandler *handler.User) *App {
	return &App{db: db, userHandler: userHandler}
}

func (a *App) InitializeRoutes() http.Handler {
	r := chi.NewRouter()

	// Rutas relacionadas con usuarios
	r.Mount("/users", a.userHandler.Routes())

	// Agregar más rutas aquí

	return r
}
