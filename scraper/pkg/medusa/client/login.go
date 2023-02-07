package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/TonyDMorris/scraper/pkg/medusa/models"
)

// LoginResponse is the response from the login endpoint
type LoginResponse struct {
	Token string `json:"token"`
}

// Login logs in a user and returns the api token

func (c *HTTPClient) login(r models.LoginRequest) (*LoginResponse, error) {
	payloadBytes, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/auth", c.baseURL), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	c.client.Jar.SetCookies(req.URL, resp.Cookies())
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unsuccessful request with response : %s", resp.Status)
	}
	loginResponse := LoginResponse{}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &loginResponse)
	if err != nil {
		return nil, err
	}
	return &loginResponse, nil

}

func (c *HTTPClient) Login() error {
	loginResponse, err := c.login(models.LoginRequest{
		Email:    c.email,
		Password: c.password,
	})
	bytes, _ := json.Marshal(c.client.Jar)
	fmt.Println(string(bytes))
	if err != nil {
		return err
	}
	c.apiToken = loginResponse.Token

	return nil
}
