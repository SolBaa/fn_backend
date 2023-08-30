package router

import (
	"github.com/SolBaa/finances-backend/cmd/api/handler"
	"github.com/go-chi/chi/v5"
	"github.com/jinzhu/gorm"
)

type Router interface {
	MapRoutes()
}

type router struct {
	r  chi.Router
	db *gorm.DB
}

func NewRouter(r chi.Router, db *gorm.DB) Router {
	return &router{r: r, db: db}
}

func (r *router) MapRoutes() {
	r.UsersRoutes()
	// r.buildProductRoutes()
	// r.buildSectionRoutes()
	// r.buildWarehouseRoutes()
	// r.buildEmployeeRoutes()
	// r.buildBuyerRoutes()
}

func (r *router) UsersRoutes() {
	r.r.Get("/users", handler.GetUsers(r.db))
}
