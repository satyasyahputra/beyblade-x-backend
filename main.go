package main

import (
	"encoding/json"
	"log"
	"net/http"
	"satyasyahputra/beyblade-x/store"

	"github.com/gorilla/mux"
)

func getBeyblades(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(store.LoadBeyblade())
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/beyblades", getBeyblades).Methods("GET")

	port := ":8000"
	log.Printf("Server Beyblade X sedang berjalan di port %s", port)
	log.Fatal(http.ListenAndServe(port, r))
}
