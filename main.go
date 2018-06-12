package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/appengine"
)

const templatePath string = "./templates/"
const contentPath string = "./content/"

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("").ParseFiles(templatePath+"base.html", templatePath+"404.html")
	if err != nil {
		log.Fatal(err)
	}
	content := getContent(contentPath + "404.json")
	err = tmpl.ExecuteTemplate(w, "base", content)
	if err != nil {
		log.Fatal(err)
	}
}

func getContent(jsonPath string) interface{} {

	content, _ := ioutil.ReadFile(jsonPath)
	var data interface{}
	err := json.Unmarshal(content, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/portfolio", portfolioHandler)
	r.HandleFunc("/about", aboutHandler)
	r.HandleFunc("/contact", contactHandler)
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	http.Handle("/", r)
}

func main() {
	appengine.Main() // Starts the server to receive requests
}
