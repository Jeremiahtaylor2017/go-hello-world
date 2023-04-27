package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jeremiahtaylor2017/go-hello-world/pkg/config"
	"github.com/jeremiahtaylor2017/go-hello-world/pkg/handlers"
	"github.com/jeremiahtaylor2017/go-hello-world/pkg/render"
)

const port = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = templateCache
	// Useful for development mode so it reads from disk
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("Starting application on port %s\n", port)

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
