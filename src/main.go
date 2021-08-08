package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mrsujitsah/bookings/pkg/config"
	"github.com/mrsujitsah/bookings/pkg/rander"

	handler "github.com/mrsujitsah/bookings/pkg/Handler"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8085"

var app config.AppConfig

var session *scs.SessionManager

func main() {

	//change this to true when in production
	app.InProduction = false
	//initialize session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	//store session
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	tc, err := rander.CreateTemplateCache()

	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handler.NewRepo(&app)

	handler.NewHandlers(repo)

	rander.NewTemplate(&app)

	// http.HandleFunc("/", handler.Repo.Home)
	// http.HandleFunc("/about", handler.Repo.About)
	fmt.Println(fmt.Sprintf("starting applicationon port %s\n", portNumber))
	// _ = http.ListenAndServe(portNumber, nil)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
