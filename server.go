package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	db		*sql.DB
	router  *mux.Router
}

func (s *Server) Initialize(host string, port int, user string, password string, dbname string, sslmode string) {
	//s := &server{}
	//s.routes()
	//return s

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

	s.router = mux.NewRouter()
}

func (s *Server) Run(addr string) {
	log.Fatalln(http.ListenAndServe(addr, s.router))
}
