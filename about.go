package main

import (
	"html/template"
	"log"
	"net/http"
)

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("").ParseFiles(
		templatePath+"base.html",
		templatePath+"about/index.html",
	)
	if err != nil {
		log.Fatal(err)
	}

	content := getContent(contentPath + "about.json")

	c := map[string]interface{}{
		"Content": content,
	}

	err = tmpl.ExecuteTemplate(w, "base", c)
	if err != nil {
		log.Fatal(err)
	}
}
