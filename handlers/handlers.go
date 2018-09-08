// Package handlers provides and handles all HTTP handlers.
package handlers

import (
	"io/ioutil"
	"net/http"
	"path"

	"github.com/gorilla/mux"
	"github.com/russross/blackfriday"
)

// Register registers all nescessary handlers and return a router.
func Register() *mux.Router {
	r := mux.NewRouter()

	r.Path("/").Methods("GET").HandlerFunc(serveMarkdown("plan.md"))
	r.Path("/plan").Methods("GET").HandlerFunc(serveMarkdown("plan.md"))
	r.Path("/readinglist").Methods("GET").HandlerFunc(serveMarkdown("reading.md"))
	r.Path("/styles").Methods("GET").HandlerFunc(serveMarkdown("style.md"))
	r.Path("/people").Methods("GET").HandlerFunc(serveMarkdown("people.md"))

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	return r
}

func serveMarkdown(name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadFile(path.Join("content", name))
		if err != nil {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		body := blackfriday.MarkdownCommon(b)
		t := parseTemplate(body)
		t.Execute(w, r, "")
	}
}
