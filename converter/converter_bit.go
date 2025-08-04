package converter

import (
	"encoding/json"
	"log"
	"os"
	"satyasyahputra/beyblade-x/store"
)

func ConvertBits() {
	bits := store.LoadBit()

	jsonData, err := json.MarshalIndent(bits, "", "  ")
	if err != nil {
		log.Fatalf("Gagal melakukan marshal JSON: %v", err)
	}

	err = os.WriteFile(store.DB_BIT_FILE, jsonData, 0644)
	if err != nil {
		log.Fatalf("Gagal menulis file: %v", err)
	}

	log.Printf("Data berhasil ditulis ke %s", store.DB_BIT_FILE)
}
