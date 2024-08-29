package handlers

import (
	"io"
	"net/http"
)

type Handler struct{}

func (h Handler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Hello, world!")
}
