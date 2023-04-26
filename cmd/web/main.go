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

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting application on port %s\n", port)
	_ = http.ListenAndServe(port, nil)
}
