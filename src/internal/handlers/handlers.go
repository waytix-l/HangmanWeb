package handlers

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contact")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about")
}

func renderTemplate(w http.ResponseWriter, tmplName string) {
	templateCache, err := createTemplateCache()
	fmt.Println(templateCache)

	if err != nil {
		panic(err)
	}

	tmpl, ok := templateCache[tmplName + ".page.tmpl"]
	fmt.Println(tmpl)

	if !ok {
		http.Error(w, "Le template n'existe pas", http.StatusInternalServerError)
		return
	}

	buf := new(bytes.Buffer)
	tmpl.Execute(buf, nil)
	buf.WriteTo(w)
}

func createTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("../../templates_HTML/*.page.tmpl")
	fmt.Println(pages)

	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		tmpl := template.Must(template.ParseFiles(page))

		layouts, err := filepath.Glob("../../templates_HTML/layouts/*.layout.tmpl")

		if err != nil {
			return cache, err
		}

		if len(layouts) > 0 {
			tmpl.ParseGlob("../../templates_HTML/layouts/*.layout.tmpl")
		}

		cache[name] = tmpl
	}

	return cache, nil
}