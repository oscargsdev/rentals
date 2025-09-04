package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Howdy, gimme a house to rent please."))
}

func registerProperty(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Registering a new property..."))
}

func viewProperty(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	msg := fmt.Sprintf("Viewing property with ID %d", id)
	w.Write([]byte(msg))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/property/register", registerProperty)
	mux.HandleFunc("/property/view/{id}", viewProperty)

	log.Print("starting server on :4001")

	err := http.ListenAndServe(":4001", mux)
	log.Fatal(err)
}
