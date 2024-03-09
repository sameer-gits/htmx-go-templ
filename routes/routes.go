package routes

import (
	"html/template"
	"net/http"
)

func Rcontent(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Contenttwo string
		Content string
	}{
		Contenttwo: "This is content 2",
	}

	contentTmpl, err := template.ParseFiles("dynamic/content2.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = contentTmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
