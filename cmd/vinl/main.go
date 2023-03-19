package main

import (
	"os"

	"vinl/internal/server"
)

func main() {
	s := server.Server{}
	s.Initialize(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	s.Run(":8080")
}
