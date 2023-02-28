package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/TonyDMorris/scraper/pkg/medusa/models"
)

func (c *HTTPClient) UploadFile(fileKey string, image []byte) (*models.FileUploadResponse, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("files", fileKey)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, bytes.NewReader(image))
	if err != nil {
		return nil, err
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/uploads", c.baseURL), bytes.NewReader(body.Bytes()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "image/jpeg")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiToken))

	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	_ = os.WriteFile("body.txt", bodyBytes, 0644)
	req.Body = io.NopCloser(bytes.NewReader(bodyBytes))

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
