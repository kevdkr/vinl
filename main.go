package main

import "os"

// func main() {
// 	s := &Server{connectDB(), mux.NewRouter()}
// 	defer s.db.Close()
// 	s.routes()
// 	//s := server{connectDB(), mux.NewRouter()}
// 	//s := newServer()
// 	log.Fatalln(http.ListenAndServe(":8080", s.router))
// }


func main() {
	s := Server{}
	s.Initialize(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	defer s.db.Close()
	s.routes()
	s.Run(":8080")
}
