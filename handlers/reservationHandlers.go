package handlers

import (
	"GoAPIREST/hotel"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func bodyToReservation(r *http.Request, res *hotel.Reservation) error {
	if r.Body == nil {
		return errors.New("request body is emty")
	}
	if res == nil {
		return errors.New("a reservation is required")
	}
	bd, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bd, res)
}

func reservationsGetAll(w http.ResponseWriter, r *http.Request) {
	reservations, err := hotel.All_reservations()
	if err != nil {
		postError(w, http.StatusInternalServerError)
		return
	}
	if r.Method == http.MethodHead {
		postBodyResponse(w, http.StatusOK, jsonResponse{})
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"reservations": reservations})
}

func reservationsGetOne(w http.ResponseWriter, r *http.Request, id_reservation uuid.UUID) {
	res, err := hotel.One_reservation(id_reservation)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			postError(w, http.StatusNotFound)
			return
		}
		postError(w, http.StatusInternalServerError)
		return
	}
	if r.Method == http.MethodHead {
		postBodyResponse(w, http.StatusOK, jsonResponse{})
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"Reservation": res})
}

func reservationsPostOne(w http.ResponseWriter, r *http.Request) {
	res := new(hotel.Reservation)
	err := bodyToReservation(r, res)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	ch := new(hotel.Chambre)
	var conf_chambre, _ = hotel.Conf()
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_chambre.DBUser, conf_chambre.DBPass, conf_chambre.DBHost, conf_chambre.DBPort, conf_chambre.DBName)), &gorm.Config{})
	if err != nil {
		return
	}
	err = db.Where("num_chambre = ?", res.Num_chambre).First(&ch).Error
	if err == nil {
		if ch.Etat != hotel.Libre {
			postError(w, http.StatusPreconditionFailed)
			return
		}
	}
	res.Id_reservation = uuid.New()
	entree, _ := hotel.ParseMysqlDate(res.Date_entree)
	sortie, _ := hotel.ParseMysqlDate(res.Date_sortie)
	res.Nuite = hotel.Nightsbeetween(sortie, entree)
	err = res.Save_reservation()
	if err != nil {
		if err == hotel.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Loresation", "/reservations/"+res.Id_reservation.String())
	w.WriteHeader(http.StatusCreated)
}

func reservationsPatchOne(w http.ResponseWriter, r *http.Request, id_reservation uuid.UUID) {
	res, err := hotel.One_reservation(id_reservation)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			postError(w, http.StatusNotFound)
			return
		}
		postError(w, http.StatusInternalServerError)
		return
	}
	err = bodyToReservation(r, res)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	res.Id_reservation = uuid.New()
	entree, _ := hotel.ParseMysqlDate(res.Date_entree)
	sortie, _ := hotel.ParseMysqlDate(res.Date_sortie)
	res.Nuite = hotel.Nightsbeetween(sortie, entree)
	err = res.Save_reservation()
	if err != nil {
		if err == hotel.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"Reservation": res})
}

func reservationsDeleteOne(w http.ResponseWriter, _ *http.Request, id_reservation uuid.UUID) {
	err := hotel.Delete_reservation(id_reservation)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			postError(w, http.StatusNotFound)
			return
		}
		postError(w, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
