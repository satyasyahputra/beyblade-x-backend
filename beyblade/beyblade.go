package beyblade

// Beyblade struct tetap sama
type Beyblade struct {
	ID           string   `json:"id"`
	ProductCodes []string `json:"product_codes"`
	Series       string   `json:"series"`
	Bit          string   `json:"bit"`
	Ratchet      string   `json:"ratchet"`
	Blade        string   `json:"blade"`
}
