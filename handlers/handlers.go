package handlers

import (
	"net/http"

	"github.com/mboldt/assignments/templates"
)

type Handler struct{}

func (h Handler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	renderer := templates.NewRenderer("templates")
	renderer.Render(w, "index.html", nil)
}
