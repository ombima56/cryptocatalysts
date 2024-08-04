package routes

import (
	"log"
	"net/http"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404\nPage Not Found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		log.Printf("Error parsing templates: %s", err)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Unxpected Error occured try again later", http.StatusInternalServerError)
		log.Printf("Error occured while executing index.html: %s\n", err)
		return
	}
}
