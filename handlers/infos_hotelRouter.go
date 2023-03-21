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

var conf_info_hotel, _ = hotel.Conf()

// infos_hotelRouter handles the infos_hotel route
func Infos_hotelRouter(w http.ResponseWriter, r *http.Request) {

	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_info_hotel.DBUser, conf_info_hotel.DBPass, conf_info_hotel.DBHost, conf_info_hotel.DBPort, conf_info_hotel.DBName)), &gorm.Config{})
	if err != nil {
		return
	}
	db.AutoMigrate(&hotel.Info_hotel{})

	fmt.Println(r.URL.Path)
	path := strings.TrimSuffix(r.URL.Path, "/")

	if path == "/infos_hotel" {
		switch r.Method {
		case http.MethodGet:
			infos_hotelGetAll(w, r)
			return
		case http.MethodPost:
			infos_hotelPostOne(w, r)
			return
		case http.MethodHead:
			infos_hotelGetAll(w, r)
			return
		case http.MethodOptions:
			postOptionsResponse(w, []string{http.MethodGet, http.MethodPost, http.MethodHead, http.MethodOptions}, nil)
			return
		default:
			postError(w, http.StatusMethodNotAllowed)
			return
		}
	}

	path = strings.TrimPrefix(path, "/infos_hotel/")

	id, err := uuid.Parse(path)

	if err != nil {
		postError(w, http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodPatch:
		infos_hotelPatchOne(w, r, id)
		return
	case http.MethodDelete:
		infos_hotelDeleteOne(w, r, id)
		return
	case http.MethodOptions:
		postOptionsResponse(w, []string{http.MethodGet, http.MethodPatch, http.MethodPut, http.MethodDelete, http.MethodHead, http.MethodOptions}, nil)
		return
	default:
		postError(w, http.StatusMethodNotAllowed)
		return
	}
}
