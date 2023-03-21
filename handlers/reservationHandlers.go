package handlers

import (
	"GoAPIREST/hotel"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
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
	res.Id_reservation = uuid.New()
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

func reservationsPutOne(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	res := new(hotel.Reservation)
	err := bodyToReservation(r, res)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	res.Id_reservation = uuid.New()
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
