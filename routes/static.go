package routes

import (
	"html/template"
	"net/http"
	views "server/vuews"
)

func ShowPages(w http.ResponseWriter, r *http.Request) {
	temp, err := views.Parse(r.URL.Path.)
	if err != nil {
		http.Error(w, "Bad Request!", http.StatusBadRequest)
	}

	temp.Execute(w, nil)
}

func NotFoundPage(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("./static/NotFound.gohtml")
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "txt/html; charset=utf-8")
	html.Execute(w, map[string]string{
		"Path": r.URL.RequestURI(),
	})
}
