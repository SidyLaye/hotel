package handlers

import (
	"GoAPIREST/hotel"
	"fmt"
	"net/http"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var conf_service_supp, _ = hotel.Conf()

// services_suppRouter handles the services_supp route
func Services_suppRouter(w http.ResponseWriter, r *http.Request) {

	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_service_supp.DBUser, conf_service_supp.DBPass, conf_service_supp.DBHost, conf_service_supp.DBPort, conf_service_supp.DBName)), &gorm.Config{})
	if err != nil {
		return
	}
	db.AutoMigrate(&hotel.Service_supp{})

	fmt.Println(r.URL.Path)
	path := strings.TrimSuffix(r.URL.Path, "/")

	if path == "/services_supp" {
		switch r.Method {
		case http.MethodGet:
			services_suppGetAll(w, r)
			return
		case http.MethodPost:
			services_suppPostOne(w, r)
			return
		case http.MethodHead:
			services_suppGetAll(w, r)
			return
		case http.MethodOptions:
			postOptionsResponse(w, []string{http.MethodGet, http.MethodPost, http.MethodHead, http.MethodOptions}, nil)
			return
		default:
			postError(w, http.StatusMethodNotAllowed)
			return
		}
	}

	path = strings.TrimPrefix(path, "/services_supp/")

	nomss := hotel.Nomss(path)

	switch r.Method {
	case http.MethodGet:
		services_suppGetOne(w, r, nomss)
		return
	case http.MethodPatch:
		services_suppPatchOne(w, r, nomss)
		return
	case http.MethodDelete:
		services_suppDeleteOne(w, r, nomss)
		return
	case http.MethodHead:
		services_suppGetOne(w, r, nomss)
	case http.MethodOptions:
		postOptionsResponse(w, []string{http.MethodGet, http.MethodPatch, http.MethodPut, http.MethodDelete, http.MethodHead, http.MethodOptions}, nil)
		return
	default:
		postError(w, http.StatusMethodNotAllowed)
		return
	}
}
