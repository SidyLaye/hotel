package handlers

import (
	"GoAPIREST/hotel"
	"fmt"
	"net/http"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var conf_categorie, _ = hotel.Conf()

// categoriesRouter handles the categories route
func CategoriesRouter(w http.ResponseWriter, r *http.Request) {

	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_categorie.DBUser, conf_categorie.DBPass, conf_categorie.DBHost, conf_categorie.DBPort, conf_categorie.DBName)), &gorm.Config{})
	if err != nil {
		return
	}
	db.AutoMigrate(&hotel.Categorie{})

	fmt.Println(r.URL.Path)
	path := strings.TrimSuffix(r.URL.Path, "/")

	if path == "/categories" {
		switch r.Method {
		case http.MethodGet:
			categoriesGetAll(w, r)
			return
		case http.MethodPost:
			categoriesPostOne(w, r)
			return
		case http.MethodHead:
			categoriesGetAll(w, r)
			return
		case http.MethodOptions:
			postOptionsResponse(w, []string{http.MethodGet, http.MethodPost, http.MethodHead, http.MethodOptions}, nil)
			return
		default:
			postError(w, http.StatusMethodNotAllowed)
			return
		}
	}

	path = strings.TrimPrefix(path, "/categories/")
	nom := hotel.Nom(path)

	if err != nil {
		postError(w, http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		categoriesGetOne(w, r, nom)
		return
	case http.MethodPatch:
		categoriesPatchOne(w, r, nom)
		return
	case http.MethodPut:
		categoriesPutOne(w, r, nom)
		return
	case http.MethodDelete:
		categoriesDeleteOne(w, r, nom)
		return
	case http.MethodHead:
		categoriesGetOne(w, r, nom)
	case http.MethodOptions:
		postOptionsResponse(w, []string{http.MethodGet, http.MethodPatch, http.MethodPut, http.MethodDelete, http.MethodHead, http.MethodOptions}, nil)
		return
	default:
		postError(w, http.StatusMethodNotAllowed)
		return
	}
}
