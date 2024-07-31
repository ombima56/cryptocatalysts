package routes

import (
	"net/http"
	"text/template"
)

func ContractServeForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/contract_form.html"))
	tmpl.Execute(w, nil)
}
