package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	//"os"
	"path/filepath"
	"text/template"
)

var functions = template.FuncMap{} 

// renderTemplate is called in the handler function to parse template using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string) {

    /*
     *_, err := RenderTemplateTest(w)
     *if err != nil {
     *    fmt.Println("Error getting template cahce", err)
     *}
     */
    tc, err := CreateTemplateCache()
    if err != nil {
        log.Fatal(err)
    }
     
    t, ok := tc[tmpl]
    if !ok {
        log.Fatal(err)
    }

    buf := new(bytes.Buffer)

    _ = t.Execute(buf, nil)

    _, err = buf.WriteTo(w)
    if err != nil{
        fmt.Println("Error writing template to buffer", err)
    }

    /*
     *parseTemplate, _ := template.ParseFiles("./templates/" + tmpl)
     *err := parseTemplate.Execute(w, nil)
     *if err != nil {
     *    fmt.Println("error parsing template:", err)
     *    return
     *}
     */
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
