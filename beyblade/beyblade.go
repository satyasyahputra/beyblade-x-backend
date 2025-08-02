package beyblade

// Beyblade struct tetap sama
type Beyblade struct {
	ID           string   `json:"id"`
	Series       string   `json:"series"`
	ProductCodes []string `json:"product_codes"`
	Blade        string   `json:"blade"`
	Assist       string   `json:"assist"`
	Ratchet      string   `json:"ratchet"`
	Bit          string   `json:"bit"`
	ImageUrls    []string `json:"image_urls"`
}
