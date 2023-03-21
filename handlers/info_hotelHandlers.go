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

func bodyToInfo_hotel(r *http.Request, i *hotel.Info_hotel) error {
	if r.Body == nil {
		return errors.New("request body is emty")
	}
	if i == nil {
		return errors.New("a user is required")
	}
	bd, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bd, i)
}

func infos_hotelGetAll(w http.ResponseWriter, r *http.Request) {
	infos_hotel, err := hotel.All_info_hotel()
	if err != nil {
		postError(w, http.StatusInternalServerError)
		return
	}
	if r.Method == http.MethodHead {
		postBodyResponse(w, http.StatusOK, jsonResponse{})
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"infos_hotel": infos_hotel})
}

func infos_hotelPostOne(w http.ResponseWriter, r *http.Request) {
	var conf_info_hotel, _ = hotel.Conf()
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_info_hotel.DBUser, conf_info_hotel.DBPass, conf_info_hotel.DBHost, conf_info_hotel.DBPort, conf_info_hotel.DBName)), &gorm.Config{})
	if err != nil {
		return
	}
	i := new(hotel.Info_hotel)
	id, _ := uuid.Parse("752bf718-ba40-4e59-8891-50a315e53a82")
	err = db.Where("id_info_hotel = ?", id).First(i).Error
	if err == nil {
		postError(w, http.StatusPreconditionFailed)
		return
	}
	_ = bodyToInfo_hotel(r, i)
	i.Id_info_hotel = id
	err = i.Save_info_hotel()
	if err != nil {
		if err == hotel.ErrRecordInvalid_info_hotel {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Location", "/infos_hotel/"+i.Id_info_hotel.String())
	w.WriteHeader(http.StatusCreated)
}

func infos_hotelPatchOne(w http.ResponseWriter, r *http.Request, id_info_hotel uuid.UUID) {
	var conf_info_hotel, _ = hotel.Conf()
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_info_hotel.DBUser, conf_info_hotel.DBPass, conf_info_hotel.DBHost, conf_info_hotel.DBPort, conf_info_hotel.DBName)), &gorm.Config{})
	if err != nil {
		return
	}
	i := new(hotel.Info_hotel)
	err = db.First(i, id_info_hotel).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			postError(w, http.StatusNotFound)
			return
		}
		postError(w, http.StatusInternalServerError)
		return
	}
	err = bodyToInfo_hotel(r, i)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	err = i.Save_info_hotel()
	if err != nil {
		if err == hotel.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"Info_hotel": i})
}

func infos_hotelDeleteOne(w http.ResponseWriter, _ *http.Request, id_info_hotel uuid.UUID) {
	err := hotel.Delete_info_hotel(id_info_hotel)
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
