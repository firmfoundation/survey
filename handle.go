package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

const tempDir = "templates"

func Index(w http.ResponseWriter, r *http.Request) error {
	ex, err := os.Executable()

	if err != nil {
		log.Fatal(err)
	}
	exPath := filepath.Dir(ex)

	if r.Method != http.MethodGet {
		return CustomeError(nil, 405, "Error: Method not allowed.")
	}

	t, _ := template.ParseFiles(exPath + "/templates/index.html")
	t.Execute(w, "")
	return nil
}
