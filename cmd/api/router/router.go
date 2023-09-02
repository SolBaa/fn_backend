package router

import (
	"net/http"

	"github.com/SolBaa/recipes-backend/cmd/api/handler"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type App struct {
	db             *gorm.DB
	userHandler    *handler.User
	recipesHandler *handler.Recipe
}

func NewRouter(db *gorm.DB, uh *handler.User, rh *handler.Recipe) *App {
	return &App{db: db, userHandler: uh, recipesHandler: rh}
}

func (a *App) InitializeRoutes() http.Handler {
	r := chi.NewRouter()

	// Rutas relacionadas con usuarios
	r.Mount("/users", a.userHandler.Routes())
	r.Mount("/recipes", a.userHandler.Routes())

	// Agregar más rutas aquí

	return r
}
