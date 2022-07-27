package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	//"os"
	"path/filepath"
	"text/template"

	"github.com/yogiadianta/go_web_app/pkg/config"
	//"github.com/yogiadianta/go_web_app/pkg/handlers"
	"github.com/yogiadianta/go_web_app/pkg/models"
)

var functions = template.FuncMap{} 

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig){
    app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData{
    return td
}

// renderTemplate is called in the handler function to parse template using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData ) {
    var tc map[string]*template.Template

    if app.UseCache {
        // get the template cache  from the  app config
        tc = app.TemplateCache
    } else {
        tc, _ = CreateTemplateCache()
    }
     
    t, ok := tc[tmpl]
    if !ok {
        log.Fatal("Could not get template from template cache")
    }

    buf := new(bytes.Buffer)

    td = AddDefaultData(td)

    _ = t.Execute(buf, td)

    _, err := buf.WriteTo(w)
    if err != nil{
        fmt.Println("Error writing template to buffer", err)
    }
}

// CreateTemplateCache create a template as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
    myCache := map[string]*template.Template{} 

    pages, err := filepath.Glob("./templates/*.page.tmpl")
    if err != nil {
        return myCache, err
    }
    
    for _, page := range pages {
        name := filepath.Base(page)

        ts, err := template.New(name).Funcs(functions).ParseFiles(page)
        if err != nil {
            return myCache, err
        }

        matches, err := filepath.Glob("./templates/*.layout.tmpl")
        if err != nil {
            return myCache, err
        }

        if len(matches) > 0 {
            ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
            if err != nil {
                return myCache, err
            }
        }

       myCache[name] = ts
    }
    return myCache, nil
}
