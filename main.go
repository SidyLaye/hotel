package main

import (
	"GoAPIREST/handlers"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/fournitures", handlers.FournituresRouter)
	http.HandleFunc("/fournitures/", handlers.FournituresRouter)
	http.HandleFunc("/users", handlers.UsersRouter)
	http.HandleFunc("/users/", handlers.UsersRouter)
	http.HandleFunc("/reservations", handlers.ReservationsRouter)
	http.HandleFunc("/reservations/", handlers.ReservationsRouter)
	http.HandleFunc("/", handlers.RootHandler)
	err := http.ListenAndServe("localhost:11111", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
