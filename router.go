package main

import (
	"github.com/gorilla/mux"
)

func router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/user", listUsers).Methods("GET")
	router.HandleFunc("/user/{lastName}", listUsers).Methods("GET")
	router.HandleFunc("/user", addUser).Methods("POST")
	router.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")
	router.HandleFunc("/user", deleteUser).Methods("DELETE")
	router.HandleFunc("/user/{id}", updateUser).Methods("PUT")

	return router
}
