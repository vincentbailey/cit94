package main

import (
	"net/http"
	"io"
)

func main() {
	http.HandleFunc("/pic2", pic2)
	http.HandleFunc("/pic", pic1)
	http.HandleFunc("/",  pix)
	http.ListenAndServe(":8080", nil)
}

func pix (w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "assets/")
}

func pic1 (w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="pic1.jpg">`)
	io.WriteString(w, `<img src="pic2.jpg">`)
}

func pic2 (w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="pic2.jpg">`)
}

