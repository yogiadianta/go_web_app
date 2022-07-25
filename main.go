package main

import (
	"errors"
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

func Divide(w http.ResponseWriter, r *http.Request) {
    f, err := divideValues(100.0, 10.0)
    if err != nil {
        fmt.Fprintf(w, "cannot divide by 0")
        return
    }

    fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 10.0, f))
}

// addValues  add two integers and return the sum
func addvalues(x, y int) int {
     return x+y
}

func divideValues(x, y float32) (float32, error) {
    if y <= 0{
        err := errors.New("cannot divide by 0")
        return 0, err
    }
    result := x/y
    return result, nil
}

// main is the main application function
func main() {
    http.HandleFunc("/", Home)
    http.HandleFunc("/about", About)
    http.HandleFunc("/divide", Divide)

    fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
    _ = http.ListenAndServe(portNumber, nil)
}
