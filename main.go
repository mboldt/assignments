package main

import (
	"log"
	"net/http"

	"github.com/mboldt/assignments/datastore"
	"github.com/mboldt/assignments/handlers"
)

type app struct {
	ds *datastore.Datastore
}

func middleware(f func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("path: %s", r.URL.Path)
		f(w, r)
	})
}

func (a *app) init() {
	handler := handlers.New(a.ds)
	http.Handle("/{$}", middleware(handler.Index))
	http.Handle("GET /add-student", middleware(handler.AddStudentGet))
	http.Handle("POST /add-student", middleware(handler.AddStudentPost))
}

func main() {
	app := &app{
		ds: datastore.New(),
	}
	app.init()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
