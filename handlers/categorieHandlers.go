package handlers

import (
	"GoAPIREST/hotel"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"gorm.io/gorm"
)

func bodyToCategorie(r *http.Request, c *hotel.Categorie) error {
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

func categoriesGetAll(w http.ResponseWriter, r *http.Request) {
	categories, err := hotel.All_categories()
	if err != nil {
		postError(w, http.StatusInternalServerError)
		return
	}
	if r.Method == http.MethodHead {
		postBodyResponse(w, http.StatusOK, jsonResponse{})
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"categories": categories})
}

func categoriesGetOne(w http.ResponseWriter, r *http.Request, nom hotel.Nom) {
	c, err := hotel.One_categorie(nom)
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
	postBodyResponse(w, http.StatusOK, jsonResponse{"Categorie": c})
}

func categoriesPostOne(w http.ResponseWriter, r *http.Request) {
	c := new(hotel.Categorie)
	err := bodyToCategorie(r, c)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	err = c.Save_categorie()
	if err != nil {
		if err == hotel.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Location", "/categories/"+string(c.Nom))
	w.WriteHeader(http.StatusCreated)
}

func categoriesPutOne(w http.ResponseWriter, r *http.Request, nom hotel.Nom) {
	c := new(hotel.Categorie)
	err := bodyToCategorie(r, c)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	err = c.Save_categorie()
	if err != nil {
		if err == hotel.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"Categorie": c})
}

func categoriesPatchOne(w http.ResponseWriter, r *http.Request, nom hotel.Nom) {
	c, err := hotel.One_categorie(nom)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			postError(w, http.StatusNotFound)
			return
		}
		postError(w, http.StatusInternalServerError)
		return
	}
	err = bodyToCategorie(r, c)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	err = c.Save_categorie()
	if err != nil {
		if err == hotel.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"Categorie": c})
}

func categoriesDeleteOne(w http.ResponseWriter, _ *http.Request, nom hotel.Nom) {
	err := hotel.Delete_categorie(nom)
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
