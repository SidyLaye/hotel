package handlers

import (
	"GoAPIREST/hotel"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var conf_chambre, _ = hotel.Conf()

// chambresRouter handles the chambres route
func ChambresRouter(w http.ResponseWriter, r *http.Request) {

	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_chambre.DBUser, conf_chambre.DBPass, conf_chambre.DBHost, conf_chambre.DBPort, conf_chambre.DBName)), &gorm.Config{})
	if err != nil {
		return
	}
	db.AutoMigrate(&hotel.Chambre{})

	fmt.Println(r.URL.Path)
	path := strings.TrimSuffix(r.URL.Path, "/")

	if path == "/chambres" {
		switch r.Method {
		case http.MethodGet:
			chambresGetAll(w, r)
			return
		case http.MethodPost:
			chambresPostOne(w, r)
			return
		case http.MethodHead:
			chambresGetAll(w, r)
			return
		case http.MethodOptions:
			postOptionsResponse(w, []string{http.MethodGet, http.MethodPost, http.MethodHead, http.MethodOptions}, nil)
			return
		default:
			postError(w, http.StatusMethodNotAllowed)
			return
		}
	}

	path = strings.TrimPrefix(path, "/chambres/")

	id, err := strconv.Atoi(path)

	if err != nil {
		postError(w, http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		chambresGetOne(w, r, id)
		return
	case http.MethodPatch:
		chambresPatchOne(w, r, id)
		return
	case http.MethodDelete:
		chambresDeleteOne(w, r, id)
		return
	case http.MethodHead:
		chambresGetOne(w, r, id)
	case http.MethodOptions:
		postOptionsResponse(w, []string{http.MethodGet, http.MethodPatch, http.MethodPut, http.MethodDelete, http.MethodHead, http.MethodOptions}, nil)
		return
	default:
		postError(w, http.StatusMethodNotAllowed)
		return
	}
}
