package main

type person struct {
	Id        int `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []person
}