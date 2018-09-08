package main

import (
	"net/http"
	"os"

	"github.com/golovers/go14/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	if err := http.ListenAndServe(":"+port, handlers.Register()); err != nil {
		panic(err)
	}
}
