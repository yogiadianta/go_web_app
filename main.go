package main

import (
	//"errors"
	"fmt"
	"net/http"
	"text/template"
)

const portNumber = ":8080"

// About is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "about.page.tmpl")
}

// renderTemplate is called in the handler function to parse template
func renderTemplate(w http.ResponseWriter, tmpl string) {
    parseTemplate, _ := template.ParseFiles("./templates/" + tmpl)
    err := parseTemplate.Execute(w, nil)
    if err != nil {
        fmt.Println("error parsing template:", err)
        return
    }
}

// main is the main application function
func main() {
    http.HandleFunc("/", Home)
    http.HandleFunc("/about", About)

    fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
    _ = http.ListenAndServe(portNumber, nil)
}
