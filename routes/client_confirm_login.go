package routes

import (
	"cryptocatalysts/core/db"
	"fmt"
	"net/http"
)

func CConfirmLogin(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	cname := r.FormValue("cname")

	client, retrieved := db.CRetrieve(cname)
	if !retrieved {
		fmt.Fprintf(w, "Failed to login the client")
		return
	}

	if email != client.Email {
		fmt.Fprintf(w, "Failed to login the client")
		return
	}
	http.Redirect(w, r, "/?login=ok", http.StatusFound)
}
