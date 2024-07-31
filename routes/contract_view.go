package routes

import (
	"cryptocatalysts/core/db"
	"net/http"
	"text/template"
)

func ContractView(w http.ResponseWriter, r *http.Request) {
	address := r.FormValue("address")
	contract := db.ConRetrieve(address)

	tmpl := template.Must(template.ParseFiles(("./templates/contract_display.html")))
	tmpl.Execute(w, contract)
}
