package download

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"net/http"
	"os"
	"satyasyahputra/beyblade-x/store"
	"strings"
)

// downloadToTempFile downloads an image and saves it to a temporary file.
func downloadToTempFile(url string) (*os.File, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat permintaan GET: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("kode status respons buruk: %s", resp.Status)
	}

	tempFile, err := os.CreateTemp("", "downloaded-image-*")
	if err != nil {
		return nil, fmt.Errorf("gagal membuat file sementara: %w", err)
	}

	_, err = io.Copy(tempFile, resp.Body)
	if err != nil {
		tempFile.Close()           // Pastikan file ditutup jika terjadi error
		os.Remove(tempFile.Name()) // Hapus file sementara
		return nil, fmt.Errorf("gagal menyalin data gambar: %w", err)
	}

	convertToJPG(tempFile.Name(), tempFile.Name())

	return tempFile, nil
}

// getExtensionFromContent reads the start of a file and determines its extension.
func getExtensionFromContent(tempFile *os.File) (string, error) {
	// Pindahkan pointer file kembali ke awal agar bisa dibaca dari awal
	_, err := tempFile.Seek(0, io.SeekStart)
	if err != nil {
		return "", fmt.Errorf("gagal memindahkan pointer file: %w", err)
	}

	buffer := make([]byte, 512)
	_, err = tempFile.Read(buffer)
	if err != nil && err != io.EOF {
		return "", fmt.Errorf("gagal membaca file header: %w", err)
	}

	contentType := http.DetectContentType(buffer)

	switch contentType {
	case "image/jpeg":
		return ".jpg", nil
	case "image/png":
		return ".png", nil
	case "image/gif":
		return ".gif", nil
	case "image/webp":
		return ".webp", nil
	default:
		if strings.HasPrefix(contentType, "image/") {
			return "." + strings.TrimPrefix(contentType, "image/"), nil
		}
		return "", fmt.Errorf("tipe konten tidak didukung: %s", contentType)
	}
}

func doDownload(imageUrl string, filePath string) {
	imageURL := imageUrl
	baseFileName := filePath

	// 1. Unduh gambar ke file sementara
	tempFile, err := downloadToTempFile(imageURL)
	if err != nil {
		fmt.Printf("Error saat mengunduh gambar: %v\n", err)
		return
	}
	defer os.Remove(tempFile.Name()) // Pastikan file sementara dihapus

	// 2. Dapatkan ekstensi dari konten file
	extension, err := getExtensionFromContent(tempFile)
	if err != nil {
		fmt.Printf("Error saat mendeteksi ekstensi: %v\n", err)
		return
	}

	// 3. Buat nama file akhir
	finalFileName := baseFileName + extension

	// 4. Tutup file sementara sebelum mengganti nama
	tempFile.Close()

	// 5. Ganti nama file sementara ke file akhir
	err = os.Rename(tempFile.Name(), finalFileName)
	if err != nil {
		fmt.Printf("Error saat mengganti nama file: %v\n", err)
		return
	}

	fmt.Printf("Gambar berhasil disimpan sebagai: %s\n", finalFileName)
}

// convertToJPG decodes an input image and encodes it as a JPEG file.
// It automatically detects the format of the input image.
func convertToJPG(inputPath, outputPath string) error {
	// 1. Open the input image file.
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open input file: %w", err)
	}
	defer inputFile.Close()

	// 2. Decode the input image.
	// The image.Decode function automatically detects the file format (PNG, GIF, etc.).
	img, _, err := image.Decode(inputFile)
	if err != nil {
		return fmt.Errorf("failed to decode image: %w", err)
	}

	// 3. Create the output JPEG file.
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outputFile.Close()

	// 4. Encode the image as a JPEG into the output file.
	// You can specify jpeg.Options to set the quality (0-100). nil uses the default.
	err = jpeg.Encode(outputFile, img, nil)
	if err != nil {
		return fmt.Errorf("failed to encode image to JPEG: %w", err)
	}

	return nil
}

func Download() {
	beyblades := store.LoadBeyblade()
	basePath := "/Users/satya.syahputra/Documents/playground/beyblade-x-backend/"

	for _, bey := range beyblades {
		if len(bey.ImageUrls) > 0 {
			name := bey.Blade + bey.Ratchet + bey.Bit
			doDownload(bey.ImageUrls[0], basePath+"beyblade-images/"+strings.ReplaceAll(name, " ", ""))
		}
	}
}
