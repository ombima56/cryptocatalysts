package routes

import (
	"net/http"
	"strconv"
	"text/template"
)

func ContractConfirmation(w http.ResponseWriter, r *http.Request) {
	client := r.FormValue("client")
	amount, _ := strconv.Atoi(r.FormValue("amount"))
	recipient := r.FormValue("recipient")
	duration := r.FormValue("duration")

	contract := struct {
		Client    string
		Recipient string
		Duration  string
		Amount    int
	}{
		Client:    client,
		Duration:  duration,
		Recipient: recipient,
		Amount:    amount,
	}

	tmpl := template.Must(template.ParseFiles(("./templates/contract_confirmation.html")))
	tmpl.Execute(w, contract)
}
