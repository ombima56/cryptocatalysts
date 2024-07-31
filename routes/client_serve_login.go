package routes

import (
	"net/http"
	"text/template"
)

func CServeLoginn(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/client_login_form.html"))
	tmpl.Execute(w, nil)
}
