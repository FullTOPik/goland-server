package views

import (
	"fmt"
	"net/http"
	"text/template"
)

func Parse(filepath string) (Template, error) {
	t, err := template.ParseFiles(filepath)
		if err != nil {
			return Template{}, fmt.Errorf("Parse error : %w", err)
		}
	
	return Template{
		htmlTml: t,
	}, nil
}

type Template struct {
	htmlTml *template.Template
}

func(t Template) Execute(w  http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.htmlTml.Execute(w, data)
		if err != nil {
			http.Error(w, "Incorrect server data!", http.StatusInternalServerError)
		}
}
