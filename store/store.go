package store

import (
	"encoding/json"
	"log"
	"os"
	"satyasyahputra/beyblade-x/beyblade"
)

const DB_BEYBLADE_FILE = "beybladex.json"
const DB_BIT_FILE = "beybladex-bit.json"

func LoadBeyblade() []beyblade.Beyblade {
	fileData, err := os.ReadFile(DB_BEYBLADE_FILE)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("File %s tidak ditemukan. Memulai dengan data kosong.", DB_BEYBLADE_FILE)
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
	log.Printf("Berhasil memuat %d data Beyblade dari %s", len(beyblades), DB_BEYBLADE_FILE)

	return beyblades
}

func LoadBit() map[string]beyblade.Bit {
	fileData, err := os.ReadFile(DB_BIT_FILE)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("File %s tidak ditemukan. Memulai dengan data kosong.", DB_BIT_FILE)
			return map[string]beyblade.Bit{}
		}
		log.Fatalf("Gagal membaca file data: %v", err)
	}

	if len(fileData) == 0 {
		return map[string]beyblade.Bit{}
	}

	var bits map[string]beyblade.Bit
	err = json.Unmarshal(fileData, &bits)
	if err != nil {
		log.Fatalf("Gagal mengubah data dari JSON: %v", err)
	}
	log.Printf("Berhasil memuat %d data Bit dari %s", len(bits), DB_BIT_FILE)

	return bits
}
