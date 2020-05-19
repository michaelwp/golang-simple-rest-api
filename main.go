package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	port := "3000"
	router := mux.NewRouter()
	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/user", allUsers).Methods("GET")
	router.HandleFunc("/user", addUser).Methods("POST")
	http.Handle("/", router)

	fmt.Println("Server listening on port ", port)
	log.Fatal(http.ListenAndServe(":" + port, router))
}