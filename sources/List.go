package sources

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Gaardsholt/go-gitguardian/client"
)

type SourcesListType string

const (
	Bitbucket SourcesListType = "bitbucket"
	GitHub    SourcesListType = "github"
	GitLab    SourcesListType = "gitlab"
)

type ListOptions struct {
	Cursor  string          `json:"cursor"`
	PerPage *int            `json:"per_page"`
	Search  string          `json:"search"`
	Type    SourcesListType `json:"type"`
}

func (c *SourcesClient) List(lo ListOptions) (*SourcesListResult, *client.PaginationMeta, error) {
	req, err := c.client.NewRequest("GET", "/v1/sources", nil)
	if err != nil {
		return nil, nil, err
	}

	// Add query parameters
	q := req.URL.Query()

	if lo.PerPage != nil {
		q.Add("per_page", strconv.Itoa(*lo.PerPage))
	}

	if lo.Cursor != "" {
		q.Add("cursor", lo.Cursor)
	}

	q.Add("search", lo.Search)
	q.Add("type", string(lo.Type))
	req.URL.RawQuery = q.Encode()

	r, err := c.client.Client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		var target Error
		decode := json.NewDecoder(r.Body)
		err = decode.Decode(&target)
		if err != nil {
			return nil, nil, err
		}
		return &SourcesListResult{Error: &target}, nil, fmt.Errorf("%s", target.Detail)
	}

	var target []SourcesResponse
	decode := json.NewDecoder(r.Body)
	err = decode.Decode(&target)
	if err != nil {
		return nil, nil, err
	}
	pagination, err := client.GetPaginationMeta(r)
	if err != nil {
		return nil, nil, err
	}

	return &SourcesListResult{Result: target}, pagination, nil
}
