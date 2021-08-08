package handler

import (
	"net/http"

	"github.com/mrsujitsah/bookings/pkg/config"
	"github.com/mrsujitsah/bookings/pkg/models"
	"github.com/mrsujitsah/bookings/pkg/rander"
)

// Repo is the repository  type
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo create a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	//grab the ip of visitor
	remoteIP := r.RemoteAddr                              //<- return ip4 or ip6 as string
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP) //setting session

	rander.RanderTemplate(w, "home.page.html", &models.TemplateData{})
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perform some logic
	stringMap := make(map[string]string)

	stringMap["test"] = "Hello again. " //send the some data to template
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	rander.RanderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
