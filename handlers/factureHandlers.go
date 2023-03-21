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

func bodyToFacture(r *http.Request, f *hotel.Facture) error {
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

func facturesGetAll(w http.ResponseWriter, r *http.Request) {
	factures, err := hotel.All_factures()
	if err != nil {
		postError(w, http.StatusInternalServerError)
		return
	}
	if r.Method == http.MethodHead {
		postBodyResponse(w, http.StatusOK, jsonResponse{})
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"factures": factures})
}

func facturesGetOne(w http.ResponseWriter, r *http.Request, id_facture uuid.UUID) {
	f, err := hotel.One_facture(id_facture)
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
	postBodyResponse(w, http.StatusOK, jsonResponse{"Facture": f})
}

func facturesPostOne(w http.ResponseWriter, r *http.Request) {
	f := new(hotel.Facture)
	err := bodyToFacture(r, f)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	f.Id_facture = uuid.New()
	err = f.Save_facture()
	if err != nil {
		if err == hotel.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Location", "/factures/"+f.Id_facture.String())
	w.WriteHeader(http.StatusCreated)
}

func facturesPutOne(w http.ResponseWriter, r *http.Request, id_facture uuid.UUID) {
	f := new(hotel.Facture)
	err := bodyToFacture(r, f)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	f.Id_facture = uuid.New()
	err = f.Save_facture()
	if err != nil {
		if err == hotel.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"Facture": f})
}

func facturesPatchOne(w http.ResponseWriter, r *http.Request, id_facture uuid.UUID) {
	f, err := hotel.One_facture(id_facture)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			postError(w, http.StatusNotFound)
			return
		}
		postError(w, http.StatusInternalServerError)
		return
	}
	err = bodyToFacture(r, f)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	f.Id_facture = uuid.New()
	err = f.Save_facture()
	if err != nil {
		if err == hotel.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"Facture": f})
}

func facturesDeleteOne(w http.ResponseWriter, _ *http.Request, id_facture uuid.UUID) {
	err := hotel.Delete_facture(id_facture)
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
