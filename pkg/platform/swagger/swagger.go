package swagger

import (
	"embed"
	"encoding/json"
	"log"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	"go.yaml.in/yaml/v4"
)

//go:embed swagger-ui/*
var SwaggerUI embed.FS

func SetupRoutes(r chi.Router, openapi *openapi3.T) {
	r.Get("/swagger/doc.json", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(openapi)
		if err != nil {
			log.Printf("failed to write swagger spec: %v", err)
		}
	})

	r.Get("/swagger/doc.yaml", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		err := yaml.NewEncoder(w).Encode(openapi)
		if err != nil {
			log.Printf("failed to write swagger spec: %v", err)
		}
	})

	r.Get("/swagger-ui/*", func(w http.ResponseWriter, r *http.Request) {
		http.FileServerFS(SwaggerUI).ServeHTTP(w, r)
	})
}
