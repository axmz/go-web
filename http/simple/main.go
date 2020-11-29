package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

type server int

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Panic(err)
	}

	data := struct {
		Method     string
		URL        *url.URL
		Proto      string // "HTTP/1.0"
		ProtoMajor int    // 1
		ProtoMinor int    // 0
		Header     http.Header
	}{
		r.Method,
		r.URL,
		r.Proto,
		r.ProtoMajor,
		r.ProtoMinor,
		r.Header,
	}

	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func main() {
	var s server
	http.ListenAndServe(":8080", s)
}
