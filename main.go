package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"git.sr.ht/~rehandaphedar/frendds/pkg/relations"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: frendds-api [port]")
	}
	port := os.Args[1]

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/{root}", getRelations)

	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}

func getRelations(w http.ResponseWriter, r *http.Request) {
	root := chi.URLParam(r, "root")
	rootRelations := relations.GetRelations(root)
	DOTString := ""

	for _, relation := range rootRelations {
		DOTString += fmt.Sprintf("\"%s\" -> \"%s\"\n", relation.Source, relation.Target)
	}

	w.Write([]byte(DOTString))
}
