package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/TonyDMorris/scraper/pkg/medusa/models"
)

func (c *HTTPClient) UploadFile(fileKey string, image []byte) (*models.FileUploadResponse, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", fileKey)
	if err != nil {
		return nil, err
	}
	_, err = part.Write(image)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/uploads", c.baseURL), body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiToken))

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unsuccessful request with status code : %v", resp.StatusCode)
	}
	var fileUploadResponse models.FileUploadResponse
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &fileUploadResponse)
	if err != nil {
		return nil, err
	}
	return &fileUploadResponse, nil

}
