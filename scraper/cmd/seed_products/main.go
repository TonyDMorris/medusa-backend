package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/TonyDMorris/scraper/pkg/castle"
	"github.com/TonyDMorris/scraper/pkg/castle/mapper"
	"github.com/TonyDMorris/scraper/pkg/medusa/client"
)

func main() {

	client := client.NewClientWithDefaultTransport("http://localhost:9000/admin", "admin@medusa-test.com", "supersecret")
	err := client.Login()
	if err != nil {
		panic(err)
	}
	dirs, err := ioutil.ReadDir("scraped/artists")
	if err != nil {
		panic(err)
	}
	cr, err := client.ListCollections(200, 0)
	if err != nil {
		panic(err)
	}
	idsMap := make(map[string]string)
	for _, c := range cr.Collections {
		idsMap[c.Handle] = *c.ID
	}
	for _, dir := range dirs {
		files, err := ioutil.ReadDir("scraped/artists/" + dir.Name() + "/pieces")
		if err != nil {
			panic(err)
		}

		for _, file := range files {
			if len(file.Name()) > 5 {
				bytes, err := ioutil.ReadFile("scraped/artists/" + dir.Name() + "/pieces/" + file.Name())
				if err != nil {
					panic(err)
				}
				// create product
				// add product to collection
				var hit castle.Hits
				err = json.Unmarshal(bytes, &hit)
				if err != nil {
					panic(err)
				}
				if hit.Thumbnail == "" {
					hit.Thumbnail = hit.SquareThumbnail
				}
				if hit.Thumbnail == "" {
					hit.Thumbnail = hit.HoverThumbnail
				}
				if hit.Thumbnail != "" {
					err := uploadFileImage(hit.Thumbnail, client)
					if err != nil {
						panic(err)
					}
				}
				product, err := mapper.ProductAlgoliaToMedusa(&hit, idsMap[hit.ArtistTitle])
				if err != nil {
					panic(err)
				}

				product, err = client.CreateProduct(product)
				if err != nil {
					panic(err)
				}
				bytes, err = json.Marshal(product)
				if err != nil {
					panic(err)
				}
				println(string(bytes))

			}
		}
	}

}

func uploadFileImage(url string, client *client.HTTPClient) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	_, err = client.UploadFile(filepath.Base(url), body)
	if err != nil {
		return err
	}

	return nil
}
