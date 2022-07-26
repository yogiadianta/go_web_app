package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yogiadianta/go_web_app/pkg/config"
	"github.com/yogiadianta/go_web_app/pkg/handlers"
	"github.com/yogiadianta/go_web_app/pkg/render"
)

const portNumber = ":8080"

// main is the main application function
func main() {
    var app config.AppConfig

    tc, err := render.CreateTemplateCache()
    if err != nil {
        log.Fatal("cannot create template cache")
    }
    
    app.TemplateCache = tc
    
    http.HandleFunc("/", handlers.Home)
    http.HandleFunc("/about", handlers.About)

    fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
    _ = http.ListenAndServe(portNumber, nil)
}
