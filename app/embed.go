package main

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"
)

//go:embed templates
var templatesFS embed.FS

//go:embed assets/static
var staticFS embed.FS

func embeddedTemplates() *template.Template {
	return template.Must(template.New("").ParseFS(templatesFS, "templates/**/*"))
}

func embeddedStatic() http.FileSystem {
	sub, err := fs.Sub(staticFS, "assets/static")
	if err != nil {
		panic("failed to sub static FS: " + err.Error())
	}
	return http.FS(sub)
}
