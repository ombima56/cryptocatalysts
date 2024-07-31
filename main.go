package main

import (
	"fmt"
	"net/http"

	"cryptocatalysts/routes"
	"cryptocatalysts/core/db"
)

func init() {
	db.Init()
}

func stripPrefixHandler(prefix string, h http.Handler) http.Handler {
	return http.StripPrefix(prefix, h)
}

func main() {
	fmt.Println("Server running on 'http://localhost:9000'")

	// static files
	staticFileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", stripPrefixHandler("/static/", staticFileServer))

	//routes
	http.HandleFunc("/", routes.Index)
	http.HandleFunc("/cserveregistration", routes.CServeRegistration)
	http.HandleFunc("/csaveregistration", routes.CSaveRegistration)
	http.HandleFunc("/cservelogin", routes.CServeLoginn)
	http.HandleFunc("/cconfirmlogin", routes.CConfirmLogin)
	http.HandleFunc("/contractserve", routes.ContractServeForm)
	http.HandleFunc("/contractconfirmation", routes.ContractConfirmation)
	http.HandleFunc("/contractcommit", routes.ContractCommit)
	http.HandleFunc("/contractview", routes.ContractView)
	http.HandleFunc("/contractfinalize", routes.ContractFinalize)
	http.HandleFunc("/contractlist", routes.ContractList)

	http.HandleFunc("/searches", routes.Searches)

	http.HandleFunc("/500", routes.Error404)
	http.HandleFunc("/400", routes.Error500)


	// server
	http.ListenAndServe(":9000", nil)
}
