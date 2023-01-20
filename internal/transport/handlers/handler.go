package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/timvkim/notes/internal/repository"
	"github.com/timvkim/notes/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h Handler) InitRouters() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/note/save", h.createNoteHandler)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("dislplaying home page ...")
		io.WriteString(w, "What's the matter?\n")
	})

	mux.HandleFunc("/note/1", func(w http.ResponseWriter, r *http.Request) {
		log.Println("dislplaying the 1st note ...")
		io.WriteString(w, "This is the first note\n")
	})

	return mux
}

// `SendNote` implements POST request
func (h Handler) createNoteHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	log.Println("creating a new note ...")
	var note repository.Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		log.Println("provided json is invalid")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(note)
}
