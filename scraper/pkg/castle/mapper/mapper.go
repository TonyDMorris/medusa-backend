package mapper

import (
	"strings"

	"github.com/TonyDMorris/scraper/pkg/castle"
	"github.com/TonyDMorris/scraper/pkg/medusa/models"
)

func ProductAlgoliaToMedusa(h *castle.Hits, collectionID string) (*models.Product, error) {
	var thumbnail *string
	if h.Thumbnail != "" {
		thumbnail = &h.Thumbnail
	}

	return &models.Product{
		Title:        &h.Title,
		Subtitle:     &h.MetaTitle,
		Description:  &h.MetaDescription,
		IsGiftcard:   false,
		Discountable: false,
		Thumbnail:    thumbnail,
		Handle:       getHandle(h.URL),
		Variants: []models.ProductVariant{{
			Title:             h.Title,
			Sku:               h.Sku,
			InventoryQuantity: 1,
			AllowBackorder:    false,
			ManageInventory:   true,
			Height:            h.Height,
			Width:             h.Width,
			Prices:            []models.Price{},
			Options:           []models.Option{},
		}},
		Height:       h.Height,
		Width:        h.Width,
		CollectionID: &collectionID,
	}, nil

}

func getHandle(url string) string {
	split := strings.Split(url, "/")
	return split[len(split)-1]
}
