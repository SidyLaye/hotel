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

func bodyToFourniture(r *http.Request, f *hotel.Fourniture) error {
	if r.Body == nil {
		return errors.New("request body is emty")
	}
	if f == nil {
		return errors.New("a user is required")
	}
	bd, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bd, f)
}

func fournituresGetAll(w http.ResponseWriter, r *http.Request) {
	fournitures, err := hotel.All_fournitures()
	if err != nil {
		postError(w, http.StatusInternalServerError)
		return
	}
	if r.Method == http.MethodHead {
		postBodyResponse(w, http.StatusOK, jsonResponse{})
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"fournitures": fournitures})
}

func fournituresPostOne(w http.ResponseWriter, r *http.Request) {
	c := new(hotel.Fourniture)
	err := bodyToFourniture(r, c)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	c.Id_fourniture = uuid.New()
	err = c.Save_fourniture()
	if err != nil {
		if err == hotel.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Location", "/fournitures/"+c.Id_fourniture.String())
	w.WriteHeader(http.StatusCreated)
}

func fournituresPatchOne(w http.ResponseWriter, r *http.Request, id_fourniture uuid.UUID) {
	f, err := hotel.One_fourniture(id_fourniture)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			postError(w, http.StatusNotFound)
			return
		}
		postError(w, http.StatusInternalServerError)
		return
	}
	err = bodyToFourniture(r, f)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	f.Id_fourniture = uuid.New()
	err = f.Save_fourniture()
	if err != nil {
		if err == hotel.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"Fourniture": f})
}

func fournituresDeleteOne(w http.ResponseWriter, _ *http.Request, id_fourniture uuid.UUID) {
	err := hotel.Delete_fourniture(id_fourniture)
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
