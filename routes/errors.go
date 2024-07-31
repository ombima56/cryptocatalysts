package routes

import (
	"net/http"
	"text/template"
)

func Error404(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/error_404.html"))
	tmpl.Execute(w, nil)
}

func Error500(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/error_500.html"))
	tmpl.Execute(w, nil)
}
