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

func GetAddWordFormHandler(w http.ResponseWriter, r *http.Request) {
	component := EdwardForm()
	err := component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error rendering in GetAddWordFormHandler: %e", err)
	}
}

func DeleteWord(w http.ResponseWriter, r *http.Request) {
	word := r.FormValue("word")
	store.DeleteWord(word)
}

func AddWordHandler(w http.ResponseWriter, r *http.Request) {
	word := r.FormValue("word")
	example := r.FormValue("example")
	store.AddWord(word, example)
	list := store.GetDictionary()
	component := ListContent(list)
	err := component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error rendering in AddWordHandler: %e", err)
	}
}
