package store

import (
	"encoding/json"
	"log"
	"os"
	"satyasyahputra/beyblade-x/beyblade"
)

const dbFile = "beybladex.json"

func LoadBeyblade() []beyblade.Beyblade {
	fileData, err := os.ReadFile(dbFile)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("File %s tidak ditemukan. Memulai dengan data kosong.", dbFile)
			return []beyblade.Beyblade{}
		}
		log.Fatalf("Gagal membaca file data: %v", err)
	}

	if len(fileData) == 0 {
		return []beyblade.Beyblade{}
	}

	var beyblades []beyblade.Beyblade
	err = json.Unmarshal(fileData, &beyblades)
	if err != nil {
		log.Fatalf("Gagal mengubah data dari JSON: %v", err)
	}
	log.Printf("Berhasil memuat %d data Beyblade dari %s", len(beyblades), dbFile)

	return beyblades
}
