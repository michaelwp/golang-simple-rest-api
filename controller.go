package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

/*
	=====================================================
	Home [GET]
	http://localhost:3000
	=====================================================
*/
func home(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "home")
}

/*
	=====================================================
	List Users [GET]
	All users : http://localhost:3000/user
	spesific user : http://localhost:3000/user/{lastName}
	=====================================================
*/
func listUsers(w http.ResponseWriter, r *http.Request) {
	var arrPerson []person
	var response response
	vars := mux.Vars(r)
	var query string

	db := mySql()
	defer db.Close()

	if vars["lastName"] != "" {
		query = "SELECT * FROM person WHERE last_name like '%" +
			vars["lastName"] + "%' ORDER BY last_name"
	} else {
		query = "SELECT * FROM person ORDER BY last_name"
	}

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

/*
	=====================================================
	Add User [POST]
	http://localhost:3000/user
	request body : {
		"first_name":"Maria",
		"last_name":"Theresia"
	}
	=====================================================
*/
func addUser(w http.ResponseWriter, r *http.Request) {
	var p person
	var response response

	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Print(err)
	}

	if p.FirstName == "" || p.LastName == "" {
		response.Status = 0
		response.Message = "Data must not be empty"
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(response)
		return
	}

	db := mySql()
	defer db.Close()

	query := "INSERT INTO person (first_name, last_name) VALUE (?,?)"

	_, err = db.Exec(query, p.FirstName, p.LastName)
	if err != nil {
		log.Print(err)
	}

	response.Status = 1
	response.Message = "Data successfully added"

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(response)
}

/*
	=====================================================
	Delete User [DELETE]
	delete all users : http://localhost:3000/user
	delete spesific user : http://localhost:3000/user/{id}
	=====================================================
*/
func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var response response
	var query string

	w.Header().Set("Content-Type", "application/json")

	db := mySql()
	defer db.Close()

	if vars["id"] != "" {
		query = "DELETE FROM person WHERE id = " + vars["id"]
	} else {
		query = "DELETE FROM person"
	}

	res, err := db.Exec(query)
	if err != nil {
		log.Print(err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Print(err)
	}

	if rowsAffected > 0 {
		response.Status = 1
		response.Message = strconv.Itoa(int(rowsAffected)) + " data successfully deleted"
		w.WriteHeader(http.StatusOK)
	} else {
		response.Status = 0
		response.Message = "Data failed to delete"
		w.WriteHeader(http.StatusBadRequest)
	}

	_ = json.NewEncoder(w).Encode(response)
}

/*
	=====================================================
	Update User [PUT]
	http://localhost:3000/user/{id}
	request body : {
		"first_name":"Maria",
		"last_name":"Theresia"
	}
	=====================================================
*/
func updateUser(w http.ResponseWriter, r *http.Request) {
	var p person
	var response response
	vars := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Print(err)
	}

	if p.FirstName == "" || p.LastName == "" || vars["id"] == "" {
		response.Status = 0
		response.Message = "Data/id must not be empty"
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(response)
		return
	}

	db := mySql()
	defer db.Close()

	query := "UPDATE person SET first_name = ?, last_name = ? WHERE id = ?"

	res, err := db.Exec(query, p.FirstName, p.LastName, vars["id"])
	if err != nil {
		log.Print(err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Print(err)
	}

	if rowsAffected > 0 {
		response.Status = 1
		response.Message = strconv.Itoa(int(rowsAffected)) + " data successfully updated"
		w.WriteHeader(http.StatusOK)
	} else {
		response.Status = 0
		response.Message = "Data failed to updated"
		w.WriteHeader(http.StatusBadRequest)
	}

	_ = json.NewEncoder(w).Encode(response)
}