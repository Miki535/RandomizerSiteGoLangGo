package main

import (
	"html/template"
	"math/rand"
	"net/http"
)

type random struct {
	Random int
}

func htmlPasre(w http.ResponseWriter, r *http.Request) {
	min := 1
	max := 10
	result := random{Random: rand.Intn(max-min) + min}
	tmpl, _ := template.ParseFiles("index.html")
	tmpl.Execute(w, result)
}

func main() {
	http.HandleFunc("/", htmlPasre)
	http.ListenAndServe(":3333", nil) //start server on localhost 3333
}
