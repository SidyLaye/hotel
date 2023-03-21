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

var conf_facture, _ = hotel.Conf()

// facturesRouter handles the factures route
func FacturesRouter(w http.ResponseWriter, r *http.Request) {

	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_facture.DBUser, conf_facture.DBPass, conf_facture.DBHost, conf_facture.DBPort, conf_facture.DBName)), &gorm.Config{})
	if err != nil {
		return
	}
	db.AutoMigrate(&hotel.Facture{})

	fmt.Println(r.URL.Path)
	path := strings.TrimSuffix(r.URL.Path, "/")

	if path == "/factures" {
		switch r.Method {
		case http.MethodGet:
			facturesGetAll(w, r)
			return
		case http.MethodPost:
			facturesPostOne(w, r)
			return
		case http.MethodHead:
			facturesGetAll(w, r)
			return
		case http.MethodOptions:
			postOptionsResponse(w, []string{http.MethodGet, http.MethodPost, http.MethodHead, http.MethodOptions}, nil)
			return
		default:
			postError(w, http.StatusMethodNotAllowed)
			return
		}
	}

	path = strings.TrimPrefix(path, "/factures/")

	id, err := uuid.Parse(path)

	if err != nil {
		postError(w, http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		facturesGetOne(w, r, id)
		return
	case http.MethodPatch:
		facturesPatchOne(w, r, id)
		return
	case http.MethodPut:
		facturesPutOne(w, r, id)
		return
	case http.MethodDelete:
		facturesDeleteOne(w, r, id)
		return
	case http.MethodHead:
		facturesGetOne(w, r, id)
	case http.MethodOptions:
		postOptionsResponse(w, []string{http.MethodGet, http.MethodPatch, http.MethodPut, http.MethodDelete, http.MethodHead, http.MethodOptions}, nil)
		return
	default:
		postError(w, http.StatusMethodNotAllowed)
		return
	}
}
