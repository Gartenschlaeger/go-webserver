package config

import "text/template"

// AppConfig holds the application configuration
type AppConfig struct {
	TemplateCache map[string]*template.Template
}
