package main

import (
	"short_url/internal/server"
)

func main() {

	server := server.NewServer()

	println("starting server...")

	err := server.ListenAndServe()
	if err != nil {
		panic("cannot start server")
	}
}
