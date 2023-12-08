package routes

import (
	"net/http"
	"strings"
)

type Router struct{}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		handler(w, r)
	default: 
		static_handlers.
	}
}


func HandleCustomFunc(URL string, function http.HandlerFunc(w http.ResponseWriter, r *http.Request)) {
	params := strings.Split(URL, "/")
	paramsURL := strings.Split(w.)

	var paramsResult = map[string]string{}

	for index, value := range params {
		if strings.Contains(value, "{") && strings.Contains(value, "}") {
			const intermediate = strings.Replace(value, "{", "")
			const result = strings.Replace(intermediate, "}", "")
		}
	}
}