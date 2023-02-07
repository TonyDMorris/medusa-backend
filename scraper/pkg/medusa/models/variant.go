package models

type ProductVariant struct {
	Title             string   `json:"title"`
	Sku               string   `json:"sku"`
	Ean               string   `json:"ean"`
	Upc               string   `json:"upc"`
	Barcode           string   `json:"barcode"`
	HsCode            string   `json:"hs_code"`
	InventoryQuantity int      `json:"inventory_quantity"`
	AllowBackorder    bool     `json:"allow_backorder"`
	ManageInventory   bool     `json:"manage_inventory"`
	Weight            float64  `json:"weight"`
	Length            float64  `json:"length"`
	Height            float64  `json:"height"`
	Width             float64  `json:"width"`
	OriginCountry     string   `json:"origin_country"`
	MidCode           string   `json:"mid_code"`
	Material          string   `json:"material"`
	Metadata          Metadata `json:"metadata"`
	Prices            []Price  `json:"prices"`
	Options           []Option `json:"options"`
}

type Price struct {
	ID           string  `json:"id"`
	RegionID     string  `json:"region_id"`
	CurrencyCode string  `json:"currency_code"`
	Amount       float64 `json:"amount"`
	MinQuantity  int     `json:"min_quantity"`
	MaxQuantity  int     `json:"max_quantity"`
}
type Option struct {
	OptionID string `json:"option_id"`
	Value    string `json:"value"`
}
