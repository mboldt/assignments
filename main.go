package main

import (
	"log"
	"net/http"

	"github.com/mboldt/assignments/handlers"
)

func init() {
	http.Handle("/", handlers.Handler{})
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}
