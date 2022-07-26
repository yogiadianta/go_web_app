package config

import "text/template"

// AppConfig Hold the application config
type AppConfig struct {
    UseCache bool
    TemplateCache map[string]*template.Template
}
