package main

import (
	"html/template"
	"log"
	"net/http"
)

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("").ParseFiles(templatePath+"base.html", templatePath+"contact/index.html")
	if err != nil {
		log.Fatal(err)
	}
	content := getContent(contentPath + "contact.json")
	err = tmpl.ExecuteTemplate(w, "base", content)
	if err != nil {
		log.Fatal(err)
	}
}
