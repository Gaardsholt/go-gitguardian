package members

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Gaardsholt/go-gitguardian/client"
	"github.com/Gaardsholt/go-gitguardian/types"
)

type ListMembersTeamRequestsOptions struct {
	Cursor  string `json:"-" url:"cursor"`   // Pagination cursor.
	PerPage *int   `json:"-" url:"per_page"` // [ 1 .. 100 ]
	TeamId  int64  `json:"-" url:"team_id"`  // Filter requests to a specific team
}

func (c *MembersClient) ListMembersTeamRequests(memberId int64, lo ListMembersTeamRequestsOptions) (*MembersResult, *client.PaginationMeta, error) {
	ep := types.Endpoints["MembersListMembersTeamRequests"]

	req, err := c.client.NewRequest(ep.Operation, fmt.Sprintf(ep.Path, memberId), lo)
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
		return &MembersResult{Error: &target}, nil, fmt.Errorf("%s", target.Detail)
	}

	var target []MembersResponse
	decode := json.NewDecoder(r.Body)
	err = decode.Decode(&target)
	if err != nil {
		return nil, nil, err
	}

	pagination, err := client.GetPaginationMeta(r)
	if err != nil {
		return nil, nil, err
	}

	return &MembersResult{Result: target}, pagination, nil
}
