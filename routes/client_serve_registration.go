package routes

import (
	"net/http"
	"text/template"
)

func CServeRegistration(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/client_registration_form.html"))
	tmpl.Execute(w, nil)
}
