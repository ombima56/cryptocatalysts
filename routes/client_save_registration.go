package routes

import (
	"cryptocatalysts/core/db"
	"cryptocatalysts/core/entities"
	"fmt"
	"net/http"
	"text/template"
)

func CSaveRegistration(w http.ResponseWriter, r *http.Request) {
	cname := r.FormValue("name")
	email := r.FormValue("email")
	bio := r.FormValue("bio")
	location := r.FormValue("location")
	account := r.FormValue("account")

	client, created := entities.NewClient(cname, email, location, bio, account)
	if !created {
		fmt.Fprintf(w, "Failed to register client")
		return
	}
	db.CSave(client.Cname, client)

	tmpl := template.Must(template.ParseFiles("./templates/client_registration_details.html"))
	tmpl.Execute(w, client)
}
