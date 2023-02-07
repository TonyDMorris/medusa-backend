package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/TonyDMorris/scraper/pkg/medusa/models"
)

func (c *HTTPClient) CreateCollection(collection *models.Collection) (*models.CollectionResponse, error) {
	var id string
	if collection.ID != nil {
		id = fmt.Sprintf("/%s", *collection.ID)
		collection.ID = nil
	}
	payloadBytes, err := json.Marshal(collection)
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/collections%s", c.baseURL, id), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		bytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unsuccessful request with response : %s and status code : %v", string(bytes), resp.StatusCode)
	}
	var collectionResponse models.CollectionResponse
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &collectionResponse)
	if err != nil {
		return nil, err
	}
	return &collectionResponse, nil
}
func (c *HTTPClient) GetCollection(id string) (*models.CollectionResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/collections/%s", c.baseURL, id), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		bytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unsuccessful request with response : %s and status code : %v", string(bytes), resp.StatusCode)
	}
	var collectionResponse models.CollectionResponse
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &collectionResponse)
	if err != nil {
		return nil, err
	}
	return &collectionResponse, nil
}

func (c *HTTPClient) ListCollections(limit, offset int) (*models.ListCollectionsResponse, error) {
	if limit == 0 {
		limit = 10
	}
	url := fmt.Sprintf("%s/collections?limit=%d&offset=%d", c.baseURL, limit, offset)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		bytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unsuccessful request with response : %s and status code : %v", string(bytes), resp.StatusCode)
	}
	var listCollectionsResponse models.ListCollectionsResponse
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &listCollectionsResponse)
	if err != nil {
		return nil, err
	}
	return &listCollectionsResponse, nil
}
