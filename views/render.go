package views

import (
	"html/template"
	"net/http"
)

func RenderAuthResponse(w http.ResponseWriter, templateName string, data interface{}) {
	t := template.Must(template.ParseFiles("views/templates/" + templateName + ".template"))
	t.Execute(w, data)
}
