package handlers

import (
	"GoAPIREST/hotel"
	"encoding/binary"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func bodyToUser(r *http.Request, u *hotel.User) error {
	if r.Body == nil {
		return errors.New("request body is emty")
	}
	if u == nil {
		return errors.New("a User is required")
	}
	bd, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bd, u)
}

func usersGetAll(w http.ResponseWriter, r *http.Request) {
	users, err := hotel.All_users()
	if err != nil {
		postError(w, http.StatusInternalServerError)
		return
	}
	if r.Method == http.MethodHead {
		postBodyResponse(w, http.StatusOK, jsonResponse{})
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"users": users})
}

func usersGetOne(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	u, err := hotel.One_user(id)
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
	postBodyResponse(w, http.StatusOK, jsonResponse{"user": u})
}

func usersPostOne(w http.ResponseWriter, r *http.Request) {
	u := new(hotel.User)
	err := bodyToUser(r, u)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		postError(w, http.StatusInternalServerError)
		return
	}
	u.Id = uuid.New()
	u.Password = string(hashedPassword)
	err = u.Save_user()
	if err != nil {
		if err == hotel.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Loresation", "/Users/"+u.Id.String())
	w.WriteHeader(http.StatusCreated)
}

func usersPatchOne(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	u, err := hotel.One_user(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			postError(w, http.StatusNotFound)
			return
		}
		postError(w, http.StatusInternalServerError)
		return
	}
	err = bodyToUser(r, u)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		postError(w, http.StatusInternalServerError)
		return
	}
	u.Id = uuid.New()
	u.Password = string(hashedPassword)
	err = u.Save_user()
	if err != nil {
		if err == hotel.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"User": u})
}

func usersDeleteOne(w http.ResponseWriter, _ *http.Request, id uuid.UUID) {
	err := hotel.Delete_user(id)
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

func authenticateUser(w http.ResponseWriter, r *http.Request, name string, password string) {
	u, err := hotel.One_user_auth(name)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			postError(w, http.StatusNotFound)
			return
		}
		postError(w, http.StatusInternalServerError)
		return
	}
	err = bodyToUser(r, u)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}

	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		postError(w, http.StatusUnauthorized)
		return
	}

	b := []byte(u.Id.String())
	i := binary.BigEndian.Uint64(b[:8])
	id := int64(i)

	// Generate JWT token
	tokenString, err := hotel.GenerateToken(id)
	if err != nil {
		postError(w, http.StatusInternalServerError)
		return
	}

	// Send response
	postBodyResponse(w, http.StatusOK, jsonResponse{"token": tokenString})
}
