package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// creating `Note` struct with relevant fields
type Note struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// `SendNote` implements POST request
func sendNote(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	log.Println("creating a new note ...")
	var note Note
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

func main() {
	log.Println("starting server ...")

	http.HandleFunc("/note/save", sendNote)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("dislplaying home page ...")
		io.WriteString(w, "What's the matter?\n")
	})

	http.HandleFunc("/note/1", func(w http.ResponseWriter, r *http.Request) {
		log.Println("dislplaying the 1st note ...")
		io.WriteString(w, "This is the first note\n")
	})

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Printf("error starting server: %s\n", err)
	}
}
