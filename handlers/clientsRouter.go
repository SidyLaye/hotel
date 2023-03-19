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

// ClientsRouter handles the clients route
func ClientsRouter(w http.ResponseWriter, r *http.Request) {
	db, _ := gorm.Open(mysql.Open(hotel.DbPath), &gorm.Config{})
	db.AutoMigrate(&hotel.Client{})
	fmt.Println(r.URL.Path)
	path := strings.TrimSuffix(r.URL.Path, "/")

	if path == "/clients" {
		switch r.Method {
		case http.MethodGet:
			clientsGetAll(w, r)
			return
		case http.MethodPost:
			clientsPostOne(w, r)
			return
		case http.MethodHead:
			clientsGetAll(w, r)
			return
		case http.MethodOptions:
			postOptionsResponse(w, []string{http.MethodGet, http.MethodPost, http.MethodHead, http.MethodOptions}, nil)
			return
		default:
			postError(w, http.StatusMethodNotAllowed)
			return
		}
	}

	path = strings.TrimPrefix(path, "/clients/")

	id, err := uuid.Parse(path)

	if err != nil {
		postError(w, http.StatusNotFound)
		return
	}

	tel := id.String()

	switch r.Method {
	case http.MethodGet:
		clientsGetOne(w, r, tel)
		return
	case http.MethodPatch:
		clientsPatchOne(w, r, tel)
		return
	case http.MethodPut:
		clientsPutOne(w, r, tel)
		return
	case http.MethodDelete:
		clientsDeleteOne(w, r, tel)
		return
	case http.MethodHead:
		clientsGetOne(w, r, tel)
	case http.MethodOptions:
		postOptionsResponse(w, []string{http.MethodGet, http.MethodPatch, http.MethodPut, http.MethodDelete, http.MethodHead, http.MethodOptions}, nil)
		return
	default:
		postError(w, http.StatusMethodNotAllowed)
		return
	}
}
