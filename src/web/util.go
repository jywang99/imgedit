package web

import "net/http"

func renderTemplate(w http.ResponseWriter, tpl string, data any) {
	err := templates.ExecuteTemplate(w, tpl+".html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

