package handler

import (
	"net/http"

	"github.com/jinzhu/gorm"
)

func GetUsers(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	}
}
