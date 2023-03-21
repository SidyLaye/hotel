package handlers

import (
	"GoAPIREST/hotel"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var conf_fourniture, _ = hotel.Conf()

// fournituresRouter handles the fournitures route
func FournituresRouter(w http.ResponseWriter, r *http.Request) {

	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_fourniture.DBUser, conf_fourniture.DBPass, conf_fourniture.DBHost, conf_fourniture.DBPort, conf_fourniture.DBName)), &gorm.Config{})
	if err != nil {
		return
	}
	db.AutoMigrate(&hotel.Fourniture{})

	fmt.Println(r.URL.Path)
	path := strings.TrimSuffix(r.URL.Path, "/")

	if path == "/fournitures" {
		switch r.Method {
		case http.MethodGet:
			fournituresGetAll(w, r)
			return
		case http.MethodPost:
			fournituresPostOne(w, r)
			return
		case http.MethodHead:
			fournituresGetAll(w, r)
			return
		case http.MethodOptions:
			postOptionsResponse(w, []string{http.MethodGet, http.MethodPost, http.MethodHead, http.MethodOptions}, nil)
			return
		default:
			postError(w, http.StatusMethodNotAllowed)
			return
		}
	}

	path = strings.TrimPrefix(path, "/fournitures/")

	id, err := uuid.Parse(path)

	if err != nil {
		postError(w, http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodPatch:
		fournituresPatchOne(w, r, id)
		return
	case http.MethodDelete:
		fournituresDeleteOne(w, r, id)
		return
	case http.MethodOptions:
		postOptionsResponse(w, []string{http.MethodGet, http.MethodPatch, http.MethodPut, http.MethodDelete, http.MethodHead, http.MethodOptions}, nil)
		return
	default:
		postError(w, http.StatusMethodNotAllowed)
		return
	}
}
