package handlers

import (
	"GoAPIREST/hotel"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

func bodyToChambre(r *http.Request, c *hotel.Chambre) error {
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

func chambresGetAll(w http.ResponseWriter, r *http.Request) {
	chambres, err := hotel.All_chambres()
	if err != nil {
		postError(w, http.StatusInternalServerError)
		return
	}
	if r.Method == http.MethodHead {
		postBodyResponse(w, http.StatusOK, jsonResponse{})
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"chambres": chambres})
}

func chambresGetOne(w http.ResponseWriter, r *http.Request, num_chambre int) {
	c, err := hotel.One_chambre(num_chambre)
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
	postBodyResponse(w, http.StatusOK, jsonResponse{"Chambre": c})
}

func chambresPostOne(w http.ResponseWriter, r *http.Request) {
	c := new(hotel.Chambre)
	err := bodyToChambre(r, c)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	err = c.Save_chambre()
	if err != nil {
		if err == hotel.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Location", "/chambres/"+strconv.Itoa(c.Num_chambre))
	w.WriteHeader(http.StatusCreated)
}

func chambresPatchOne(w http.ResponseWriter, r *http.Request, num_chambre int) {
	c, err := hotel.One_chambre(num_chambre)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			postError(w, http.StatusNotFound)
			return
		}
		postError(w, http.StatusInternalServerError)
		return
	}
	err = bodyToChambre(r, c)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	err = c.Save_chambre()
	if err != nil {
		if err == hotel.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"Chambre": c})
}

func chambresDeleteOne(w http.ResponseWriter, _ *http.Request, num_chambre int) {
	err := hotel.Delete_chambre(num_chambre)
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
