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
	http.HandleFunc("/pic5.jpg",  pix5)
	http.HandleFunc("/pic4.jpg",  pix4)
	http.HandleFunc("/pic3.jpg",  pix3)
	http.HandleFunc("/pic2.jpg",  pix2)
	http.HandleFunc("/pic1.jpg",  pix1)
	http.ListenAndServe(":8080", nil)
}

func pic1 (w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="pic1.jpg">`)
}
func pix1 (w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "assets/pic1.jpg")
}

func pic2 (w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="pic2.jpg">`)
}
func pix2 (w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "assets/pic2.jpg")
}

func pic3 (w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="pic3.jpg">`)
}
func pix3 (w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "assets/pic3.jpg")
}

func pic4 (w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="pic4.jpg">`)
}
func pix4 (w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "assets/pic4.jpg")
}

func pic5 (w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="pic5.jpg">`)
}
func pix5 (w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "assets/pic5.jpg")
}