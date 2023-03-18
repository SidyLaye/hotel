package handlers

import (
	"GoAPIREST/hotel"
	"fmt"
	"net/http"
)

var conf, _ = hotel.Conf()

// RootHandler handles the root route
func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Asset not found\n"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.DBUser, conf.DBPass, conf.DBHost, conf.DBPort, conf.DBName)))
}
