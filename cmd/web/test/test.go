package test

import (
	"log"
	"net/http"
)

func TestTabHandler(w http.ResponseWriter, r *http.Request) {
	component := Test()
	err := component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error rendering in HelloWebHandler: %e", err)
	}
}
