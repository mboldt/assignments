package handlers

import (
	"log"
	"net/http"

	"github.com/mboldt/assignments/datastore"
	"github.com/mboldt/assignments/student"
	"github.com/mboldt/assignments/templates"
)

type Handler struct {
	ds *datastore.Datastore
}

func New(ds *datastore.Datastore) *Handler {
	return &Handler{ds: ds}
}

func (h *Handler) Index(w http.ResponseWriter, _ *http.Request) {
	renderer := templates.NewRenderer("templates")
	renderer.Render(w, "index.html", h.ds)
}

func (h *Handler) AddStudentGet(w http.ResponseWriter, _ *http.Request) {
	renderer := templates.NewRenderer("templates")
	renderer.Render(w, "add-student.html", nil)
}

func (h *Handler) AddStudentPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Printf("add student post error: %s", err)
		http.Error(w, "error reading form", 400)
	}
	log.Printf("received form data %v", r.Form)
	log.Printf("received body %v", r.Body)
	h.ds.AddStudent(student.New(r.Form.Get("name")))
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
