package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/waytix-l/HangmanWeb/config"
	"github.com/waytix-l/HangmanWeb/hangman"
)

var h hangman.HangmanData

func Home(w http.ResponseWriter, r *http.Request) {
	h.RandomWord()
	println(h.Word)
	renderTemplate(w, "home", hangman.HangmanData{Word: h.Word})
}

func Contact(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contact", nil)
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about", nil)
}

func Corruscant(w http.ResponseWriter, r *http.Request) {
	
}

var appConfig *config.Config

func CreateTemplates(app *config.Config) {
	appConfig = app
}

func renderTemplate(w http.ResponseWriter, tmplName string, data interface{}) {
	templateCache := appConfig.TemplateCache

	tmpl, ok := templateCache[tmplName+".page.html"]
	fmt.Println(tmpl)

	if !ok {
		http.Error(w, "Le template n'existe pas", http.StatusInternalServerError)
		return
	}

	buf := new(bytes.Buffer)
	tmpl.Execute(buf, data)
	buf.WriteTo(w)
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("assets/templates_HTML/*.page.html")
	fmt.Println(pages)

	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		tmpl := template.Must(template.ParseFiles(page))

		layouts, err := filepath.Glob("assets/templates_HTML/layouts/*.layout.html")

		if err != nil {
			return cache, err
		}

		if len(layouts) > 0 {
			tmpl.ParseGlob("assets/templates_HTML/layouts/*.layout.html")
		}

		cache[name] = tmpl
	}

	return cache, nil
}
