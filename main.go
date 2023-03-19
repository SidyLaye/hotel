package main

import (
	"GoAPIREST/handlers"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/clients", handlers.ClientsRouter)
	http.HandleFunc("/clients/", handlers.ClientsRouter)
	http.HandleFunc("/", handlers.RootHandler)
	err := http.ListenAndServe("localhost:11111", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
