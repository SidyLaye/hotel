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

var conf_reservation, _ = hotel.Conf()

// reservationsRouter handles the reservations route
func ReservationsRouter(w http.ResponseWriter, r *http.Request) {

	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_reservation.DBUser, conf_reservation.DBPass, conf_reservation.DBHost, conf_reservation.DBPort, conf_reservation.DBName)), &gorm.Config{})
	if err != nil {
		return
	}
	db.AutoMigrate(&hotel.Reservation{})

	fmt.Println(r.URL.Path)
	path := strings.TrimSuffix(r.URL.Path, "/")

	if path == "/reservations" {
		switch r.Method {
		case http.MethodGet:
			reservationsGetAll(w, r)
			return
		case http.MethodPost:
			reservationsPostOne(w, r)
			return
		case http.MethodHead:
			reservationsGetAll(w, r)
			return
		case http.MethodOptions:
			postOptionsResponse(w, []string{http.MethodGet, http.MethodPost, http.MethodHead, http.MethodOptions}, nil)
			return
		default:
			postError(w, http.StatusMethodNotAllowed)
			return
		}
	}

	path = strings.TrimPrefix(path, "/reservations/")

	id, err := uuid.Parse(path)

	if err != nil {
		postError(w, http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		reservationsGetOne(w, r, id)
		return
	case http.MethodPatch:
		reservationsPatchOne(w, r, id)
		return
	case http.MethodDelete:
		reservationsDeleteOne(w, r, id)
		return
	case http.MethodHead:
		reservationsGetOne(w, r, id)
	case http.MethodOptions:
		postOptionsResponse(w, []string{http.MethodGet, http.MethodPatch, http.MethodPut, http.MethodDelete, http.MethodHead, http.MethodOptions}, nil)
		return
	default:
		postError(w, http.StatusMethodNotAllowed)
		return
	}
}
