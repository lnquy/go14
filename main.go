package main

import (
	"net/http"
	"os"

	"github.com/golovers/go14/handlers"
)

func main() {
	address := os.Getenv("GO14_HTTP_ADDRESS")
	if address == "" {
		address = ":80"
	}
	if err := http.ListenAndServe(address, handlers.Register()); err != nil {
		panic(err)
	}
}
