package converter

import (
	"encoding/json"
	"log"
	"os"
	"satyasyahputra/beyblade-x/store"
)

func Convert() {
	beyblades := store.LoadBeyblade()
	for i, bey := range beyblades {
		if bey.ImageUrls == nil {
			bey.ImageUrls = []string{}
		}
		beyblades[i] = bey
	}

	jsonData, err := json.MarshalIndent(beyblades, "", "  ")
	if err != nil {
		log.Fatalf("Gagal melakukan marshal JSON: %v", err)
	}

	err = os.WriteFile(store.DBFile, jsonData, 0644)
	if err != nil {
		log.Fatalf("Gagal menulis file: %v", err)
	}

	log.Printf("Data berhasil ditulis ke %s", store.DBFile)
}
