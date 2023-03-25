package members

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Gaardsholt/go-gitguardian/client"
	"github.com/Gaardsholt/go-gitguardian/types"
)

type MembersListRole string

const (
	Owner      MembersListRole = "owner"
	Manager    MembersListRole = "manager"
	Member     MembersListRole = "member"
	Viewer     MembersListRole = "viewer"
	Restricted MembersListRole = "restricted"
)

type ListOptions struct {
	Cursor  string          `json:"-" url:"cursor"`   // Pagination cursor.
	PerPage *int            `json:"-" url:"per_page"` // [ 1 .. 100 ]
	Search  string          `json:"-" url:"search"`   // Search members based on their name or email.
	Role    MembersListRole `json:"-" url:"role"`     // Filter members based on their role.
}

func (c *MembersClient) List(lo ListOptions) (*MembersResult, *client.PaginationMeta, error) {
	ep := types.Endpoints["MembersList"]

	req, err := c.client.NewRequest(ep.Operation, ep.Path, lo)
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
