package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// About is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "This is the home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
    sum := addvalues(2, 2)
    _, _ = fmt.Fprintf(w, fmt.Sprintf("This is at the about page and 2 + 2 is %d", sum))
}


// addValues  add two integers and return the sum
func addvalues(x, y int) int {
     return x+y
}

// main is the main application function
func main() {
    http.HandleFunc("/", Home)
    http.HandleFunc("/about", About)

    fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
    _ = http.ListenAndServe(portNumber, nil)
}
