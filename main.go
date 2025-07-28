// main.go
package main

import (
	"encoding/json" // Paket untuk encoding dan decoding JSON
	"log"           // Paket untuk logging error
	"net/http"      // Paket untuk membangun server HTTP

	// Paket untuk data waktu (opsional, bisa digunakan nanti)
	"github.com/google/uuid" // Untuk menghasilkan ID unik (UUID)
	"github.com/gorilla/mux" // Router HTTP yang kuat
)

// ========= 1. STRUKTUR DATA (MODEL) =========
// Beyblade struct akan menjadi cetak biru untuk data kita.
// Tag `json:"..."` digunakan untuk memberi tahu Go bagaimana cara
// mengubah field struct ini menjadi format JSON dan sebaliknya.
type Beyblade struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	SubTitle    string   `json:"sub_title"`
	ImageUrls   []string `json:"image_urls"`
	Description string   `json:"description"`
	Bit         string   `json:"bit"`
	Ratchet     string   `json:"ratchet"`
	Blade       string   `json:"blade"`
}

// ========= 2. DATABASE DALAM MEMORI =========
// Untuk memulai, kita akan menggunakan slice (array dinamis) Go sebagai database.
// Nantinya, ini bisa diganti dengan koneksi ke database sungguhan.
var beyblades []Beyblade

// ========= 3. HANDLER FUNCTIONS =========
// Handler bertanggung jawab atas logika untuk setiap endpoint.

// getBeyblades akan mengembalikan semua data beyblade sebagai JSON.
func getBeyblades(w http.ResponseWriter, r *http.Request) {
	// Set header Content-Type agar client tahu bahwa responsnya adalah JSON.
	w.Header().Set("Content-Type", "application/json")

	// Encode slice beyblades kita ke format JSON dan kirim sebagai respons.
	json.NewEncoder(w).Encode(beyblades)
}

// createBeyblade akan menambahkan beyblade baru dari request body.
func createBeyblade(w http.ResponseWriter, r *http.Request) {
	// Set header Content-Type ke JSON.
	w.Header().Set("Content-Type", "application/json")

	// Deklarasikan variabel baru dengan tipe Beyblade.
	var beyblade Beyblade

	// Decode data JSON dari body request ke dalam variabel beyblade.
	err := json.NewDecoder(r.Body).Decode(&beyblade)
	if err != nil {
		// Jika ada error saat decoding (misalnya, format JSON salah),
		// kirim respons error.
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Buat ID unik baru untuk beyblade yang baru dibuat.
	beyblade.ID = uuid.New().String()

	// Tambahkan beyblade baru ke slice database kita.
	beyblades = append(beyblades, beyblade)

	// Atur status code respons ke 201 Created.
	w.WriteHeader(http.StatusCreated)

	// Kirim kembali data beyblade yang baru dibuat sebagai konfirmasi.
	json.NewEncoder(w).Encode(beyblade)
}

// ========= 4. FUNGSI UTAMA (MAIN) =========
// Ini adalah titik masuk aplikasi kita.
func main() {
	// Inisialisasi router baru dari Gorilla Mux.
	r := mux.NewRouter()

	// (Opsional) Tambahkan beberapa data awal untuk testing.
	beyblades = append(beyblades, Beyblade{
		ID:          uuid.New().String(),
		Name:        "Dran Sword",
		SubTitle:    "3-60F",
		ImageUrls:   []string{"https://example.com/dran_sword.jpg"},
		Description: "A powerful attack-type Beyblade.",
		Blade:       "Dran Sword",
		Ratchet:     "3-60",
		Bit:         "Flat",
	})

	// Definisikan Routes (Endpoint) kita.
	// Endpoint untuk mendapatkan semua beyblade (Method: GET)
	r.HandleFunc("/beyblades", getBeyblades).Methods("GET")
	// Endpoint untuk membuat beyblade baru (Method: POST)
	r.HandleFunc("/beyblades", createBeyblade).Methods("POST")

	// Jalankan server.
	port := ":8000"
	log.Printf("Server Beyblade X sedang berjalan di port %s", port)
	// http.ListenAndServe akan memulai server dan akan berhenti jika ada error.
	log.Fatal(http.ListenAndServe(port, r))
}
