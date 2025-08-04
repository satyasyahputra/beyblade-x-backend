package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"satyasyahputra/beyblade-x/converter"
	"satyasyahputra/beyblade-x/download"
	"satyasyahputra/beyblade-x/store"

	"github.com/gorilla/mux"
)

func getBeyblades(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(store.LoadBeyblade())
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Sub-perintah dibutuhkan: 'hello' atau 'goodbye'")
		return
	}

	subcommand := os.Args[1]

	switch subcommand {
	case "router":
		router()
	case "convert":
		converter.Convert()
	case "convert-bits":
		converter.ConvertBits()
	case "download":
		download.Download()
	default:
		fmt.Printf("Perintah tidak dikenal: %s\n", subcommand)
	}
}

func router() {
	r := mux.NewRouter()

	r.HandleFunc("/beyblades", getBeyblades).Methods("GET")

	port := ":8000"
	log.Printf("Server Beyblade X sedang berjalan di port %s", port)
	log.Fatal(http.ListenAndServe(port, r))
}
