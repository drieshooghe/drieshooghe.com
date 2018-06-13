package main

import (
	"html/template"
	"log"
	"net/http"
)

func disclaimerHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("").ParseFiles(templatePath+"base.html", templatePath+"disclaimer/index.html")
	if err != nil {
		log.Fatal(err)
	}
	content := getContent(contentPath + "disclaimer.json")
	err = tmpl.ExecuteTemplate(w, "base", content)
	if err != nil {
		log.Fatal(err)
	}
}
