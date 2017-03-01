package main

import (
	"net/http"
	"io"
)

func main() {
	http.HandleFunc("/pic5", pic5)
	http.HandleFunc("/pic4", pic4)
	http.HandleFunc("/pic3", pic3)
	http.HandleFunc("/pic2", pic2)
	http.HandleFunc("/", pic1)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func pic1 (w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/resources/pic1.jpg">`)
}

func pic2 (w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/resources/pic2.jpg">`)
}

func pic3 (w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/resources/pic3.jpg">`)
}

func pic4 (w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/resources/pic4.jpg">`)
}

func pic5 (w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/resources/pic5.jpg">`)
}