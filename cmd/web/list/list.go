package list

import (
	"dictionary-htmx-demo/internal/store"
	"log"
	"net/http"
)

func ListPageHandler(w http.ResponseWriter, r *http.Request) {
	list := store.GetDictionary()
	component := ListPage(list)
	err := component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error rendering in ListPageHandler: %e", err)
	}
}

func ListHandler(w http.ResponseWriter, r *http.Request) {
	list := store.GetDictionary()
	component := List(list)
	err := component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error rendering in ListHandler: %e", err)
	}
}

func GetListContentHandler(w http.ResponseWriter, r *http.Request) {
	list := store.GetDictionary()
	component := ListContent(list)
	err := component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error rendering in ListHandler: %e", err)
	}
}
