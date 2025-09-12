package main

import (
	"fmt"
	"io"
	"net/http"
)

const (
	AppIDTickets = 97
	AppIDKBSC    = 96
)

type Client struct {
	APIKey  string
	BaseURL string
	Client  *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		APIKey:  apiKey,
		BaseURL: "https://support.stedwards.edu/TDWebApi/api/",
		Client:  &http.Client{},
	}
}

func (c *Client) DoRequest(method, path string) ([]byte, error) {
	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.BaseURL, path), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("API request failed: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
