package main

import (
	"embed"
	"mime"
	"net/http"
	"strings"
)

//go:embed assets/*
var assets embed.FS

// Serves files from the ./assets folder
func Assets() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		filepath := r.PathValue("path")
		split := strings.Split(filepath, ".")
		ext := "." + split[len(split)-1]
		if ext != "." {
			w.Header().Set("Content-Type", mime.TypeByExtension(ext))
		}
		http.ServeFileFS(w, r, assets, "/assets/"+filepath)
	})
}
