package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	w.Write([]byte("Howdy, gimme a house to rent please."))
}

func registerPropertyForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Displaying a form to register a new property..."))
}

func registerPropertyPost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Registering a new property..."))
}

func viewProperty(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Viewing property with ID %d", id)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /property/register", registerPropertyForm)
	mux.HandleFunc("POST /property/register", registerPropertyPost)
	mux.HandleFunc("GET /property/view/{id}", viewProperty)

	log.Print("starting server on :4001")

	err := http.ListenAndServe(":4001", mux)
	log.Fatal(err)
}
