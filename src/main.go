package main

import (
	"fmt"
	"net/http"
	"github.com/waytix-l/HangmanWeb/config"
	"github.com/waytix-l/HangmanWeb/internal/handlers"
)

func main() {
	var appConfig config.Config

	templateCache, err := handlers.CreateTemplateCache()

	if err != nil {
		panic(err)
	}

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	appConfig.TemplateCache = templateCache
	appConfig.Port = ":8877"

	handlers.CreateTemplates(&appConfig)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/contact", handlers.Contact)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("corruscant", handlers.Corruscant)

	fmt.Println("(htpp://localhost:8877) - Server started on port ", appConfig.Port)
	http.ListenAndServe(appConfig.Port, nil)
}
