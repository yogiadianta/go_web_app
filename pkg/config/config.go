package config

import "text/template"

// AppConfig Hold the application config
type AppConfig struct {
    TemplateCache map[string]*template.Template
}
