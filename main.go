package main

import (
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Input  string
	Output string
}

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func handleIndex(w http.ResponseWriter, r *http.Request) {
	data := PageData{}

	if r.Method == http.MethodPost {
		inputText := r.FormValue("text")
		data.Input = inputText

		result := collapseSpaces(inputText)
		result = fixQuotes(result)
		result = fixPunctuationSpacing(result)
		result = capitalizeSentences(result)

		data.Output = result
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handleIndex)
	log.Println("Server running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}