package web

import (
	"html/template"
	"log"
	"net/http"
	"regexp"
)

// cache templates
var templates = template.Must(template.ParseGlob("static/*.html"))
var validPath = regexp.MustCompile("^/(index|generate)/$")

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil && r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		fn(w, r)
	}
}

func RunWeb() {
    http.HandleFunc("/", makeHandler(indexHandler))
    http.HandleFunc("/index/", makeHandler(indexHandler))
    http.HandleFunc("/generate/", makeHandler(generateHandler))
	log.Fatal(http.ListenAndServe("172.28.73.221:8080", nil))
}

