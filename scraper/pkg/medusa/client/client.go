package client

import (
	"net/http"
	"net/http/cookiejar"
	"time"
)

type HTTPClient struct {
	baseURL  string
	apiToken string
	email    string
	password string
	client   *http.Client
}

func NewClient(baseURL, email, password string, client *http.Client) *HTTPClient {
	return &HTTPClient{
		baseURL:  baseURL,
		client:   client,
		email:    email,
		password: password,
	}
}

func NewClientWithDefaultTransport(baseURL, email, password string) *HTTPClient {
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}

	client := &http.Client{Timeout: 10 * time.Second, Jar: jar}
	return NewClient(baseURL, email, password, client)
}
