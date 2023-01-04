package main

import "github.com/gorilla/mux"

// func main() {
// 	s := &Server{connectDB(), mux.NewRouter()}
// 	defer s.db.Close()
// 	s.routes()
// 	//s := server{connectDB(), mux.NewRouter()}
// 	//s := newServer()
// 	log.Fatalln(http.ListenAndServe(":8080", s.router))
// }


func main() {
	s := Server{connectDB(), mux.NewRouter()}
	defer s.db.Close()
	//s.Initialize(
		//os.Getenv("DB_USERNAME"),
		//os.Getenv("DB_PASSWORD"),
		//os.Getenv("DB_NAME"),
	//	"localhost",
	//	5432,
	//	"postgres",
	//	"postgres",
	//	"postgres",
	//	"disable",
	//)
	s.routes()
	s.Run(":8080")
}
