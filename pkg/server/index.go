package server

import (
	"net/http"
	"strings"
)

type RouterType struct {
	item []rrr
}

type rrr struct {
	Path    string
	Method  string
	Handler HandlerFunc
}

type Params = map[string]any

type HandlerFunc = func(http.ResponseWriter, *http.Request, Params) 

var Server *http.Server
var Routers RouterType

var routes = make(map[string]RouterType)

type countHandler struct{}

func Listen(port int, handler func(error)) {

	err := http.ListenAndServe(":33333", &countHandler{})
	if err != nil && handler != nil {

		handler(err)
	}
}

func (s *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	clientPath := r.URL.Path
	for path, route := range routes {

		for _, value := range route.item {
			var params = make(map[string]any)

			var pathItems = strings.Split((path+value.Path), "/")
			var pathClientItems = strings.Split(clientPath, "/")
			var isEqualStrings = true

			if len(pathItems) == len(pathClientItems) {
				for indexPathItem, pathItem := range pathItems {
					if strings.Contains(pathItem, ":") {
						params[strings.Trim(pathItem, ":")] = pathClientItems[indexPathItem]
					} else if(isEqualStrings) {
						isEqualStrings = pathItem == pathClientItems[indexPathItem]
					}
				}
			} else {
				isEqualStrings = false
			}
			

			if isEqualStrings && r.Method == value.Method {
				value.Handler(w, r, params)
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Page not found!"))
	}
}

func Use(path string, handler RouterType) {
	routes[path] = handler
}

func (r *RouterType) Get(Path string, Handler HandlerFunc) {
	r.item = append(r.item, rrr{Path: Path, Method: http.MethodGet, Handler: Handler})
}

func (r *RouterType) Post(Path string, Handler HandlerFunc) {
	r.item = append(r.item, rrr{Path: Path, Method: http.MethodPost, Handler: Handler})
}

func (r *RouterType) Delete(Path string, Handler HandlerFunc) {
	r.item = append(r.item, rrr{Path: Path, Method: http.MethodDelete, Handler: Handler})
}

func (r *RouterType) Patch(Path string, Handler HandlerFunc) {
	r.item = append(r.item, rrr{Path: Path, Method: http.MethodPatch, Handler: Handler})
}

func BadRequestAnswer(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}
