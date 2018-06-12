package main

import (
	"html/template"
	"log"
	"net/http"
)

func portfolioHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("").ParseFiles(
		templatePath+"base.html",
		templatePath+"portfolio/index.html",
	)
	if err != nil {
		log.Fatal(err)
	}

	content := getContent(contentPath + "portfolio.json")

	var gh = &github{}
	repositories := gh.getInfo(r)

	c := map[string]interface{}{
		"Content":      content,
		"Repositories": repositories,
	}

	err = tmpl.ExecuteTemplate(w, "base", c)
	if err != nil {
		log.Fatal(err)
	}
}
