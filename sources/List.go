package sources

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Gaardsholt/go-gitguardian/client"
	"github.com/Gaardsholt/go-gitguardian/types"
)

type SourcesListType string

const (
	Bitbucket SourcesListType = "bitbucket"
	GitHub    SourcesListType = "github"
	GitLab    SourcesListType = "gitlab"
)

type ListOptions struct {
	Cursor  string          `json:"-" url:"cursor"`
	PerPage *int            `json:"-" url:"per_page"`
	Search  string          `json:"-" url:"search"`
	Type    SourcesListType `json:"-" url:"type"`
}

func (c *SourcesClient) List(lo ListOptions) (*SourcesListResult, *client.PaginationMeta, error) {
	ep := types.Endpoints["SourcesList"]

	req, err := c.client.NewRequest(ep.Operation, ep.Path, lo)
	if err != nil {
		return nil, nil, err
	}

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
