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

type Bit struct {
	Description string   `json:"description"`
	Weight      string   `json:"weight"`
	Design      []string `json:"design"`
	Strengths   []string `json:"strengths"`
	Weakness    []string `json:"weakness"`
	Gameplan    []string `json:"gameplan"`
	ImageUrl    string   `json:"image_url"`
}

type BitMap map[string]Bit
