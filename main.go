package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := router()
	port := "3000"

	http.Handle("/", router)
	fmt.Println("Server listening on port ", port)
	log.Fatal(http.ListenAndServe(":" + port, router))
}