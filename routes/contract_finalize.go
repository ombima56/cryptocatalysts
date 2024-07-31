package routes

import (
	"cryptocatalysts/core/db"
	"net/http"
)

func ContractFinalize(w http.ResponseWriter, r *http.Request) {
	address := r.FormValue("address")

	contract := db.ConRetrieve(address)
	if !contract.Closed {
		contract.Closed = true
	}
	db.ConSave(contract.Address, &contract)

	http.Redirect(w, r, "/", http.StatusFound)
}
