package main

import (
	"github.com/k2madureira/goscrap/internal/server"
)

func main() {

	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic("Cannot start server")
	}
}
