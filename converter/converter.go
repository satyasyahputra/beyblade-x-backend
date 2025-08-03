package converter

import (
	"encoding/json"
	"log"
	"os"
	"satyasyahputra/beyblade-x/store"

	"github.com/google/uuid"
)

const BASE_IMAGE_URL = "https://satyasyahputra.github.io/beyblade-x-backend/beyblade-images/"

func Convert() {
	beyblades := store.LoadBeyblade()
	for i, bey := range beyblades {
		if bey.ImageUrls == nil {
			bey.ImageUrls = []string{}
		}
		bey.ID = uuid.New().String() // re-generate ID
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
