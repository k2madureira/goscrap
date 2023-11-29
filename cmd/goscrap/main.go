package main

import (
	"fmt"
	"net/http"

	"github.comm/k2madureira/goscrap/internal/routes"
)

func main() {
	router := routes.NewRouter()

	port := 3000
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server listening on http://localhost%s\n", addr)
	err := http.ListenAndServe(addr, router)

	if err != nil {
		panic(err)
	}
}
