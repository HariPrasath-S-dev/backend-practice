package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
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
	h2:=func(w http.ResponseWriter, r *http.Request){
		time.Sleep(1 * time.Second)
		title:=r.PostFormValue("title")
		director:=r.PostFormValue("director")
		
		tmpl:=template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "film-list-element",Film{Title: title,Director: director})

		// htmlStr := fmt.Sprintf("<li class=\"list-group-item bg-primary text-white\">%s - %s</li>", title, director)
		// tmpl,_:=template.New("t").Parse(htmlStr)

		// tmpl.Execute(w,nil)

		// log.Print("HTMX Request")

		// // HX - Request is a header set by HTMX
		// log.Print(r.Header.Get("HX - Request"))
	}
	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/",h2)

	log.Fatal(http.ListenAndServe(":8080", nil)) // Specify port number and handler

}
