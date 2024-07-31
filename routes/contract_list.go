package routes

import (
	"cryptocatalysts/core/db"
	"net/http"
	"text/template"
)

func ContractList(w http.ResponseWriter, r *http.Request) {
	contracts, listed := db.ConList()
	if !listed {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	tmpl := template.Must(template.ParseFiles("./templates/contract_list.html"))
	tmpl.Execute(w, contracts)
}
