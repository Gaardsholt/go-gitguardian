package client

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type Client struct {
	Server string // GITGUARDIAN_SERVER
	ApiKey string // GITGUARDIAN_API_KEY
	Client HttpRequest
}

type ClientOption func(*Client) error

type HttpRequest interface {
	Do(req *http.Request) (*http.Response, error)
}

func New(opts ...ClientOption) (*Client, error) {

	client := Client{}

	// Add all the provided options to the client
	for _, v := range opts {
		if err := v(&client); err != nil {
			return nil, err
		}
	}

	if client.Client == nil {
		client.Client = &http.Client{}
	}

	if client.Server == "" {
		client.Server = os.Getenv("GITGUARDIAN_SERVER")
	}
	if client.Server == "" {
		client.Server = "https://api.gitguardian.com"
	}

	if client.ApiKey == "" {
		client.ApiKey = os.Getenv("GITGUARDIAN_API_KEY")
	}
	if client.ApiKey == "" {
		return nil, fmt.Errorf("GITGUARDIAN_API_KEY is not set")
	}

	return &client, nil
}

func (c *Client) NewRequest(method string, path string, payload *bytes.Buffer) (*http.Request, error) {
	serverURL, err := url.Parse(c.Server)
	if err != nil {
		return nil, err
	}

	queryURL, err := serverURL.Parse(path)
	if err != nil {
		return nil, err
	}

	var req *http.Request
	if payload == nil {
		req, err = http.NewRequest(method, queryURL.String(), nil)
	} else {
		req, err = http.NewRequest(method, queryURL.String(), payload)
	}

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Token "+c.ApiKey)

	return req, nil
}

func WithHTTPClient(ht HttpRequest) ClientOption {
	return func(c *Client) error {
		c.Client = ht
		return nil
	}
}
func WithServer(server string) ClientOption {
	return func(c *Client) error {
		c.Server = server
		return nil
	}
}
func WithApiKey(apiKey string) ClientOption {
	return func(c *Client) error {
		c.ApiKey = apiKey
		return nil
	}
}
