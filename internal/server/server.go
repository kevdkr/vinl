package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

type Server struct {
	db		*sql.DB
	router  *mux.Router
}

func (s *Server) Initialize(host string, port string, user string, password string, dbname string, sslmode string) {

	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var err error
	s.db, err = sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal(err)
	}

	err = s.db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Connected to db %s", host)

	s.router = mux.NewRouter()
	s.routes()
}

func (s *Server) Run(addr string) {
	defer s.db.Close()
	log.Fatalln(http.ListenAndServe(addr, s.router))
}
