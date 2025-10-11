package web

import (
	"embed"
	"html/template"
	"io/fs"
	"log"
)

//go:embed templates templates/organizations
var templateFS embed.FS

func Parse() *template.Template {
	t := template.New("")

	patterns := []string{
		"templates/*.html",
		"templates/organizations/*.html",
		"templates/jobs/*.html",
		"templates/users/*.html",
	}

	templateCount := 0
	for _, pat := range patterns {
		matches, err := fs.Glob(templateFS, pat)
		if err != nil {
			log.Printf("Error globbing pattern %s: %v", pat, err)
			continue
		}
		if len(matches) == 0 {
			log.Printf("No matches found for pattern %s", pat)
			continue
		}

		log.Printf("Found %d templates matching pattern %s", len(matches), pat)
		for _, match := range matches {
			log.Printf("  - %s", match)
		}

		template.Must(t.ParseFS(templateFS, pat))
		templateCount += len(matches)
	}

	log.Printf("Total templates loaded: %d", templateCount)

	// Optional fallback to avoid panic if no files exist yet
	if t.Tree == nil || len(t.Templates()) == 0 {
		log.Println("No templates found, creating fallback template")
		template.Must(t.New("fallback.html").Parse(`<html><body><h1>Welcome to go-db-demo</h1><p>Templates not loaded yet.</p></body></html>`))
	}

	return t
}
