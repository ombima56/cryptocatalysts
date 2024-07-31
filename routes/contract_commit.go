package routes

import (
	"cryptocatalysts/core/db"
	"cryptocatalysts/core/entities"
	"net/http"
	"strconv"
	"text/template"
)

func ContractCommit(w http.ResponseWriter, r *http.Request) {
	client := r.FormValue("client")
	amount, _ := strconv.Atoi(r.FormValue("amount"))
	recipient := r.FormValue("recipient")
	duration, _ := strconv.Atoi(r.FormValue("duration"))

	_, found := db.CRetrieve(recipient)
	if !found {
		http.Redirect(w, r, "/500", http.StatusFound)
		return
	}

	contract, created := entities.NewContract(client, recipient, amount, duration)
	if !created {
		http.Redirect(w, r, "/500", http.StatusFound)
		return
	}

	db.ConSave(contract.Address, contract)
	tmpl := template.Must(template.ParseFiles("./templates/contract_commit_details.html"))
	tmpl.Execute(w, contract)
}
