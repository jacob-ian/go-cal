package main

import (
	"embed"
	"html/template"
	"io"
)

//go:embed views/*
var views embed.FS

type Template struct {
}

// Renders a template matching a name
func (t *Template) RenderPage(w io.Writer, name string, data interface{}) error {
	tmpl, err := template.ParseFS(views, "views/index.html", name)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, data)
}

// Renders a template snippet matching a name
func (t *Template) RenderSnippet(w io.Writer, file string, name string, data interface{}) error {
	tmpl, err := template.ParseFS(views, "views/index.html", file)
	if err != nil {
		return err
	}
	return tmpl.ExecuteTemplate(w, name, data)
}

// Create a new template renderer
func NewTemplateRenderer() *Template {
	return &Template{}
}
