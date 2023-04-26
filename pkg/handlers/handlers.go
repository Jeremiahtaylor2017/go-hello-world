package handlers

import (
	"net/http"

	"github.com/jeremiahtaylor2017/go-hello-world/pkg/config"
	"github.com/jeremiahtaylor2017/go-hello-world/pkg/models"
	"github.com/jeremiahtaylor2017/go-hello-world/pkg/render"
)

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is the repositry type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some business logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	// send the data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
