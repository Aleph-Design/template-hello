package handlers

import (
	"net/http"

	"github.com/aleph-design/hello/pkg/config"
	"github.com/aleph-design/hello/pkg/models"
	"github.com/aleph-design/hello/pkg/render"
)

// Repository used by handlers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// Create a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Set handlers repository
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	m.App.Session.Put(r.Context(), "rem_key", r.RemoteAddr)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// some initial logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Data from handlers"

	stringMap["rem_key"] =  m.App.Session.GetString(r.Context(), "rem_key")

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
