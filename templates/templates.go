package templates

import (
	"net/http"
	"text/template"

	"github.com/mboldt/assignments/student"
)

type IndexData struct {
	Students []*student.Student
}

func Render(w http.ResponseWriter, name string, data interface{}) {
	// TODO optimize so we don't read file every time.
	// Good for now to iterate quickly.
	t, err := template.ParseGlob("./*.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	err = t.ExecuteTemplate(w, name, data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
