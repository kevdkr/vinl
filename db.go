package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host 		= "localhost"
	port		= 5432
	user		= "postgres"
	password 	= "postgres"
	dbname 		= "postgres"
)

func connectDB() (*sql.DB) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal(err)
	}

	//defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Connected to %s", host)
	return db
}
