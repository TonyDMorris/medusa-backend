package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/TonyDMorris/scraper/pkg/medusa/models"
)

type CreateProductResponse struct {
	Product models.Product `json:"product"`
}

func (c *HTTPClient) CreateProduct(p *models.Product) (*models.Product, error) {

	payloadBytes, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/admin/products", c.baseURL), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		bytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unsuccessful request with response : %s", string(bytes))
	}
	productResp := CreateProductResponse{}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &productResp)
	if err != nil {
		return nil, err
	}
	return &productResp.Product, nil
}
