package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	asciiart "ascii-art/asciiArt"
)

type dataRender struct {
  Input    string
  AsciiArt string
}

var (
	tmpl    = template.Must(template.ParseFiles("./templates/index.html"))
	lastOut dataRender
)


func StartServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", renderPage)
	mux.HandleFunc("/ascii-art", handlePost)
	serv := &http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	fmt.Println("Server running at http://localhost:3000/")
	log.Fatal(serv.ListenAndServe())
}

func renderPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		tmpErr := template.Must(template.ParseFiles("./templates/error.html"))
		w.WriteHeader(http.StatusNotFound)
		tmpErr.Execute(w, nil)
		return
	}
	if err := tmpl.Execute(w, lastOut); err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
	}
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed: only POST is accepted", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("inputText")
	font := r.FormValue("dropDown")

	
	out, err := asciiart.AsciiArt(text, font)
	nD := &dataRender {
		Input: text,
		AsciiArt: out,
	}
	if err != nil {
		http.Error(w, "AsciiArt error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	lastOut = *nD
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
