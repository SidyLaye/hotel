package handlers

import (
	"GoAPIREST/hotel"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/asdine/storm"
	"gopkg.in/mgo.v2/bson"
)

func bodyToClient(r *http.Request, c *hotel.Client) error {
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

func clientsGetAll(w http.ResponseWriter, r *http.Request) {
	clients, err := hotel.All()
	if err != nil {
		postError(w, http.StatusInternalServerError)
		return
	}
	if r.Method == http.MethodHead {
		postBodyResponse(w, http.StatusOK, jsonResponse{})
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"clients": clients})
}

func clientsGetOne(w http.ResponseWriter, r *http.Request, id bson.ObjectId) {
	c, err := hotel.One(id)
	if err != nil {
		if err == storm.ErrNotFound {
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
	postBodyResponse(w, http.StatusOK, jsonResponse{"client": c})
}

func clientsPostOne(w http.ResponseWriter, r *http.Request) {
	c := new(hotel.Client)
	err := bodyToClient(r, c)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	c.ID = bson.NewObjectId()
	err = c.Save()
	if err != nil {
		if err == hotel.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Location", "/clients/"+c.ID.Hex())
	w.WriteHeader(http.StatusCreated)
}

func clientsPutOne(w http.ResponseWriter, r *http.Request, id bson.ObjectId) {
	c := new(hotel.Client)
	err := bodyToClient(r, c)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	c.ID = id
	err = c.Save()
	if err != nil {
		if err == hotel.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"client": c})
}

func clientsPatchOne(w http.ResponseWriter, r *http.Request, id bson.ObjectId) {
	c, err := hotel.One(id)
	if err != nil {
		if err == storm.ErrNotFound {
			postError(w, http.StatusNotFound)
			return
		}
		postError(w, http.StatusInternalServerError)
		return
	}
	err = bodyToClient(r, c)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	c.ID = id
	err = c.Save()
	if err != nil {
		if err == hotel.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"client": c})
}

func clientsDeleteOne(w http.ResponseWriter, _ *http.Request, id bson.ObjectId) {
	err := hotel.Delete(id)
	if err != nil {
		if err == storm.ErrNotFound {
			postError(w, http.StatusNotFound)
			return
		}
		postError(w, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
