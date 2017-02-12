package main

import (
	"net/http"
	"html/template"
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("layouts/*.html"))
}

func main() {
	http.ListenAndServe(":8080", nil)
	http.HandleFunc("/", index)
}

func index(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w, "index.html", nil)
}
