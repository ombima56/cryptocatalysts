package routes

import (
	"cryptocatalysts/core/db"
	"cryptocatalysts/core/entities"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

func ContractCommit(w http.ResponseWriter, r *http.Request) {
	client := r.FormValue("client")
	amount, _ := strconv.Atoi(r.FormValue("amount"))
	recipient := r.FormValue("recipient")
	duration, _ := strconv.Atoi(r.FormValue("duration"))

	contractor, found := db.CRetrieve(recipient)
	if !found {
		fmt.Fprintf(w, "No such contractor %s", contractor)
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
