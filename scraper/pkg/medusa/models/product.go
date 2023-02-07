package models

type Product struct {
	Title        *string   `json:"title,omitempty"`
	Subtitle     *string   `json:"subtitle,omitempty"`
	Description  *string   `json:"description,omitempty"`
	IsGiftcard   bool      `json:"is_giftcard,omitempty"`
	Discountable bool      `json:"discountable,omitempty"`
	Images       []*string `json:"images,omitempty"`
	Thumbnail    *string   `json:"thumbnail,omitempty"`
	Handle       string    `json:"handle,omitempty"`
	Status       *string   `json:"status,omitempty"`
	Type         Type      `json:"type,omitempty"`
	CollectionID *string   `json:"collection_id,omitempty"`
	Tags         *[]Tag    `json:"tags,omitempty"`
	// SalesChannels []SalesChannel `json:"sales_channels,omitempty"`
	Options       *[]Option        `json:"options,omitempty"`
	Variants      []ProductVariant `json:"variants,omitempty"`
	Weight        float64          `json:"weight,omitempty"`
	Length        float64          `json:"length,omitempty"`
	Height        float64          `json:"height,omitempty"`
	Width         float64          `json:"width,omitempty"`
	HsCode        *string          `json:"hs_code,omitempty"`
	OriginCountry *string          `json:"origin_country,omitempty"`
	MidCode       *string          `json:"mid_code,omitempty"`
	Material      *string          `json:"material,omitempty"`
	Metadata      Metadata         `json:"metadata,omitempty"`
}
type Type struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}
type Tag struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}
type SalesChannel struct {
	ID string `json:"id"`
}

type Metadata map[string]interface{}
