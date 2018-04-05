package controllers

import (
	"database/sql"
	"log"
)

type Context struct {
	DB *sql.DB
}

func DBInitial(*sql.DB) *sql.DB {
	var err error
	db, err := sql.Open("mysql", "root:root@/DBSelling")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
