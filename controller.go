package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "home")
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	var arrPerson []person
	var response response

	db := connect()
	defer db.Close()

	query := "SELECT * FROM person"

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}

	defer rows.Close()

	for rows.Next() {
		var el = person{}
		err := rows.Scan(&el.Id, &el.FirstName, &el.LastName)
		if err != nil {
			log.Print(err)
		}

		arrPerson = append(arrPerson, el)
	}

	response.Status = 1
	response.Message = "List of person"
	response.Data = arrPerson

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	var p person
	var response response

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Print(err)
	}

	db := connect()
	defer db.Close()

	query := "INSERT INTO golang.person (first_name, last_name) VALUE (?,?)"

	_, err = db.Exec(query, p.FirstName, p.LastName)
	if err != nil {
		log.Print(err)
	}

	response.Status = 1
	response.Message = "Data added successfully"

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(response)
}

func removeUser(w http.ResponseWriter, r *http.Request) {

}

func updateUser(w http.ResponseWriter, r *http.Request) {

}

func findUser(w http.ResponseWriter, r *http.Request) {

}