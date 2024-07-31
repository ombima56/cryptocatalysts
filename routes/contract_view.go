package routes

import (
	"cryptocatalysts/core/db"
	"cryptocatalysts/core/entities"
	"net/http"
	"text/template"
	"time"
)

func ContractView(w http.ResponseWriter, r *http.Request) {
	address := r.FormValue("address")
	contract := db.ConRetrieve(address)

	var finalize bool
	current := time.Now()
	left := contract.End.Sub(current)
	if left <= time.Second && !contract.Closed {
		finalize = true
	}

	data := struct {
		Finalize bool
		Contract entities.Contract
	}{
		Finalize: finalize,
		Contract: contract,
	}
	tmpl := template.Must(template.ParseFiles(("./templates/contract_display.html")))
	tmpl.Execute(w, data)
}
