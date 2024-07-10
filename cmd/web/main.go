package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aleph-design/hello/pkg/config"
	"github.com/aleph-design/hello/pkg/handlers"
	"github.com/aleph-design/hello/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const PORT = ":8080"

// declare var's to make them available to entire main package
var app config.AppConfig
var session *scs.SessionManager

func main() {
	app.Production = false // Reset in production mode S3-L35-14:50
	app.SetSecure = false

	// initiate sessions
	session = scs.New() // now it's the same as above
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.SetSecure // TRUE for https ============

	// make session app wide available
	app.Session = session

	// create template cache
	tmpCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tmpCache

	// make app config.AppConfig available to handlers package
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// make app config.AppConfig available to render package
	render.NewRender(&app)

	// start server
	fmt.Println("Start application on port: ", PORT)
	srv := &http.Server{
		Addr: PORT,
		// Handler: routes(&app),	// pass data to routes
		Handler: routes(),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
