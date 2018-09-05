package handlers

import (
	"html/template"
	"net/http"
)

func parseTemplate(body []byte) *appTemplate {
	t := template.Must(template.ParseFiles("templates/base.html"))
	template.Must(t.New("body").Parse(string(body)))
	return &appTemplate{t.Lookup("base.html")}
}

type appTemplate struct {
	template *template.Template
}

func (t *appTemplate) Execute(w http.ResponseWriter, r *http.Request, data interface{}) error {
	d := struct {
		Data interface{}
	}{
		Data: data,
	}
	return t.template.Execute(w, d)
}
