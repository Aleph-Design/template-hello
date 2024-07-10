package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// This paclage must be available from anywhere in the application
// This package must NOT import packages from other parts of the app

// AppConfig holds global application data ~ the infoFile.
type AppConfig struct {
	Production    bool
	SetSecure     bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	Session       *scs.SessionManager
}
