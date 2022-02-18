package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/Gaardsholt/go-gitguardian/types"
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

func (c *Client) newContentScanRequest(payload types.ContentScanPayload) (*http.Request, error) {
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(payload)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(c.Server)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("/v1/scan")

	queryURL, err := serverURL.Parse(path)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), b)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Token "+c.ApiKey)

	return req, nil
}

func (c *Client) ContentScan(payload types.ContentScanPayload) (*types.ContentScanResult, error) {
	req, err := c.newContentScanRequest(payload)
	if err != nil {
		return nil, err
	}

	r, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		var target types.Error
		decode := json.NewDecoder(r.Body)
		err = decode.Decode(&target)
		if err != nil {
			return nil, err
		}
		return &types.ContentScanResult{Error: &target}, fmt.Errorf("%s", target.Detail)
	}

	var target types.ContentScanResponse
	decode := json.NewDecoder(r.Body)
	err = decode.Decode(&target)
	if err != nil {
		return nil, err
	}

	return &types.ContentScanResult{Result: &target}, nil
}
