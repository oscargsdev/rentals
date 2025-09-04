package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Howdy, gimme a house to rent please."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Print("starting server on :4001")

	err := http.ListenAndServe(":4001", mux)
	log.Fatal(err)
}
