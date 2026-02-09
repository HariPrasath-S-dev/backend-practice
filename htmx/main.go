package main

import (
	"html/template"
	"log"
	"net/http"
)
type Film struct{
	Title string
	Director string
}
func main() {
	h1 := func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, "Hello world\n")
		// io.WriteString(w, r.Method)

		// To read the HTML file
		tmpl:=template.Must(template.ParseFiles("index.html"))

		films:=map[string][]Film{
			"Films":{
			{Title:"Star Wars", Director:"George Lucas"},
			{Title:"Star Trek", Director:"James Cameron"},
			{Title: "Avatar", Director:"James Cameron"},
		},
	}
		// To execute the template
		tmpl.Execute(w, films)


	}
	http.HandleFunc("/", h1)
	log.Fatal(http.ListenAndServe(":8080", nil)) // Specify port number and handler

}
