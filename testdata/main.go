package main

import (
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/ichiban/assets"
)

func main() {
	l, err := assets.New()
	if err != nil {
		log.Fatalf("assets.New() failed: %v", err)
	}
	defer l.Close()

	log.Printf("assets: %s", l.Path)

	t := template.Must(template.ParseGlob(filepath.Join(l.Path, "templates", "*")))
	if err := t.ExecuteTemplate(os.Stdout, "hello.tmpl", "World"); err != nil {
		log.Fatalf("t.ExecuteTemplate() failed: %v", err)
	}
}
