package main

import (
	"net/http"
	"html/template"
	"go/constant"
)

type person struct {
	First string
	Last  string
}

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/mystruct", mystruct)
	http.HandleFunc("/mapkv", mymapkv)
	http.HandleFunc("/map", mymap)
	http.HandleFunc("/slice", slice)
	http.HandleFunc("/bool", other)
	http.HandleFunc("/string", stuff)
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w, "index.html", 42)
}

func stuff(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w, "string.html", "Hello string")
}

func other(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w, "bool.html", constant.Bool)
}

func slice(w http.ResponseWriter, r *http.Request){
	p1 := person{"vince",
		"bailey",
	}
	tpl.ExecuteTemplate(w, "slice.html", p1)
}

func mymap(w http.ResponseWriter, r *http.Request){
	m := map[string]int{"vince": 32}
	tpl.ExecuteTemplate(w, "map.html", m)
}

func mymapkv(w http.ResponseWriter, r *http.Request){
	m := map[string]int{"vince": 32}
	tpl.ExecuteTemplate(w, "mapkv.html", m)
}


func mystruct(w http.ResponseWriter, r *http.Request){
	p1 := person{"vince",
		"bailey",
	}
	tpl.ExecuteTemplate(w, "mystruct.html", p1)
}