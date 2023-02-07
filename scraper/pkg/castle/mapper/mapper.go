package mapper

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/TonyDMorris/scraper/pkg/castle"
	"github.com/TonyDMorris/scraper/pkg/medusa/models"
)

func ProductAlgoliaToMedusa(h *castle.Hits, collectionID string) (*models.Product, error) {
	var thumbnail *string
	if h.Thumbnail != "" {

		req, err := http.NewRequest("GET", h.Thumbnail, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		image := string(bytes)
		thumbnail = &image
		fmt.Print("image: ", image)
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
