package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func mySql() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/golang")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
