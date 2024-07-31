package routes

import (
	"net/http"
	"text/template"
)

func Searches(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/searches.html"))
	tmpl.Execute(w,nil)
}
