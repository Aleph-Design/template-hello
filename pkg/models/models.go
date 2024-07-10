package models

// Holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
