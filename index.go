package main

import (
	"html/template"
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("").ParseFiles(templatePath+"base.html", templatePath+"home/index.html")
	if err != nil {
		log.Fatal(err)
	}
	content := getContent(contentPath + "index.json")
	err = tmpl.ExecuteTemplate(w, "base", content)
	if err != nil {
		log.Fatal(err)
	}
}
