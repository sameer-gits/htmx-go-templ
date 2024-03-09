package main

import (
	"html/template"
	"net/http"

	"github.com/sameer-gits/htmx-go-templ/routes"
)

func Index(w http.ResponseWriter, r *http.Request) {
	indexdata := struct {
		Title   string
		Heading string
		Content string
	}{
		Title:   "This is Title",
		Heading: "This is Heading",
		Content: "This is Content 1",
	}

	contentTmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = contentTmpl.Execute(w, indexdata)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/rcontent", routes.Rcontent)
	http.HandleFunc("/", Index)
	http.ListenAndServe(":3000", nil)
}
