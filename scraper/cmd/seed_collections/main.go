package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/TonyDMorris/scraper/pkg/medusa/client"
	"github.com/TonyDMorris/scraper/pkg/medusa/models"
)

func main() {

	// get dir names in scraped/artists

	// for each dir name, create a collection
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
	for _, dir := range dirs {
		// create collection
		bytes, err := ioutil.ReadFile("scraped/artists/" + dir.Name() + "/profile/bio.json")
		if err != nil {
			panic(err)
		}
		bio := make(map[string]interface{})
		err = json.Unmarshal(bytes, &bio)
		if err != nil {
			panic(err)
		}
		split := strings.Split(dir.Name(), "-")

		for i, s := range split {
			split[i] = strings.Title(s)
		}

		name := strings.Join(split, " ")
		collection := models.Collection{
			Title:    dir.Name(),
			Handle:   name,
			Metadata: bio,
		}
		col, err := client.CreateCollection(&collection)
		if err != nil {
			for _, c := range cr.Collections {
				if c.Title != dir.Name() {
					continue
				}
				collection.ID = c.ID
			}
			if collection.ID == nil {
				panic(err)
			}
			col, err = client.CreateCollection(&collection)
			if err != nil {
				panic(err)
			}

		}
		bytes, err = json.Marshal(col)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(bytes))
	}

}
