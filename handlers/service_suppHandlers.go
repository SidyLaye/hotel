package handlers

import (
	"GoAPIREST/hotel"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"gorm.io/gorm"
)

func bodyToService_supp(r *http.Request, c *hotel.Service_supp) error {
	if r.Body == nil {
		return errors.New("request body is emty")
	}
	if c == nil {
		return errors.New("a user is required")
	}
	bd, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bd, c)
}

func services_suppGetAll(w http.ResponseWriter, r *http.Request) {
	services_supp, err := hotel.All_service_supp()
	if err != nil {
		postError(w, http.StatusInternalServerError)
		return
	}
	if r.Method == http.MethodHead {
		postBodyResponse(w, http.StatusOK, jsonResponse{})
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"services_supp": services_supp})
}

func services_suppGetOne(w http.ResponseWriter, r *http.Request, nomss hotel.Nomss) {
	c, err := hotel.One_service_supp(nomss)
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
	postBodyResponse(w, http.StatusOK, jsonResponse{"Service_supp": c})
}

func services_suppPostOne(w http.ResponseWriter, r *http.Request) {
	s := new(hotel.Service_supp)
	err := bodyToService_supp(r, s)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	err = s.Save_service_supp()
	if err != nil {
		if err == hotel.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Location", "/services_supp/"+string(s.Nom))
	w.WriteHeader(http.StatusCreated)
}

func services_suppPatchOne(w http.ResponseWriter, r *http.Request, nomss hotel.Nomss) {
	s, err := hotel.One_service_supp(nomss)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			postError(w, http.StatusNotFound)
			return
		}
		postError(w, http.StatusInternalServerError)
		return
	}
	err = bodyToService_supp(r, s)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	err = s.Save_service_supp()
	if err != nil {
		if err == hotel.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"Service_supp": s})
}

func services_suppDeleteOne(w http.ResponseWriter, _ *http.Request, nomss hotel.Nomss) {
	err := hotel.Delete_service_supp(nomss)
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
