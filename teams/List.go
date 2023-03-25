package teams

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Gaardsholt/go-gitguardian/client"
	"github.com/Gaardsholt/go-gitguardian/types"
)

type ListOptions struct {
	Cursor   string `json:"-" url:"cursor"`    // Pagination cursor.
	PerPage  *int   `json:"-" url:"per_page"`  // [ 1 .. 100 ]
	IsGlobal bool   `json:"-" url:"is_global"` // Filter on/exclude the "All-incidents" team.
	Search   string `json:"-" url:"search"`    // Search teams based on their name.
}

func (c *TeamsClient) List(lo ListOptions) (*TeamsResult, *client.PaginationMeta, error) {
	ep := types.Endpoints["TeamsList"]

	req, err := c.client.NewRequest(ep.Operation, ep.Path, nil)
	if err != nil {
		return nil, nil, err
	}

	// Validate query parameters
	if lo.PerPage != nil {
		if !(*lo.PerPage >= 1 && *lo.PerPage <= 100) {
			return nil, nil, fmt.Errorf("PerPage must be between 1 and 100")
		}
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
		return &TeamsResult{Error: &target}, nil, fmt.Errorf("%s", target.Detail)
	}

	var target []TeamsResponse
	decode := json.NewDecoder(r.Body)
	err = decode.Decode(&target)
	if err != nil {
		return nil, nil, err
	}

	pagination, err := client.GetPaginationMeta(r)
	if err != nil {
		return nil, nil, err
	}

	return &TeamsResult{Result: target}, pagination, nil
}
