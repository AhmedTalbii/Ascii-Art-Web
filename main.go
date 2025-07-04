package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	asciiart "ascii-art/asciiArt"
)

var (
	tampl   = template.Must(template.ParseFiles("./templates/index.html"))
	lastOut string
)

func HandleRendering(w http.ResponseWriter, r *http.Request) {
	tampl.Execute(w, lastOut)
}

func HandlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	text := r.FormValue("inputText")
	file := r.FormValue("dropDown")

	out, errA := asciiart.AsciiArt(text, file)
	if errA != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	lastOut = out
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		HandleRendering(w, r)
	})
	fmt.Println("server running on http://localhost:3000/")
	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatalf("HTTP server failed: %v", err)
	}
}
