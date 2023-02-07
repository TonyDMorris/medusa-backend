package main

import (
	"encoding/json"
	"fmt"

	"github.com/TonyDMorris/scraper/pkg/medusa/client"
)

func main() {
	client := client.NewClientWithDefaultTransport("http://localhost:9000/admin", "admin@medusa-test.com", "supersecret")
	err := client.Login()
	if err != nil {
		panic(err)
	}

	cr, err := client.ListCollections(200, 0)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Total collections: %d", len(cr.Collections))
	for _, c := range cr.Collections {
		col, err := client.GetCollection(*c.ID)
		if err != nil {
			panic(err)
		}
		bytes, err := json.Marshal(col.Collection)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(bytes))

	}

}
