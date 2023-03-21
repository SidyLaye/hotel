package main

import (
	"GoAPIREST/handlers"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/factures", handlers.FacturesRouter)
	http.HandleFunc("/factures/", handlers.FacturesRouter)
	http.HandleFunc("/services_supp", handlers.Services_suppRouter)
	http.HandleFunc("/services_supp/", handlers.Services_suppRouter)
	http.HandleFunc("/categories", handlers.CategoriesRouter)
	http.HandleFunc("/categories/", handlers.CategoriesRouter)
	http.HandleFunc("/chambres", handlers.ChambresRouter)
	http.HandleFunc("/chambres/", handlers.ChambresRouter)
	http.HandleFunc("/infos_hotel", handlers.Infos_hotelRouter)
	http.HandleFunc("/infos_hotel/", handlers.Infos_hotelRouter)
	http.HandleFunc("/fournitures", handlers.FournituresRouter)
	http.HandleFunc("/fournitures/", handlers.FournituresRouter)
	http.HandleFunc("/users", handlers.UsersRouter)
	http.HandleFunc("/users/", handlers.UsersRouter)
	http.HandleFunc("/reservations", handlers.ReservationsRouter)
	http.HandleFunc("/reservations/", handlers.ReservationsRouter)
	http.HandleFunc("/", handlers.RootHandler)
	err := http.ListenAndServe(":11111", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
