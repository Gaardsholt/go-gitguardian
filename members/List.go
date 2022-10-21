package members

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Gaardsholt/go-gitguardian/client"
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
	Cursor  string          `json:"string"`   // Pagination cursor.
	PerPage *int            `json:"per_page"` // [ 1 .. 100 ]
	Search  string          `json:"search"`   // Search members based on their name or email.
	Role    MembersListRole `json:"role"`     // Filter members based on their role.
}

func (c *MembersClient) List(lo ListOptions) (*MembersResult, *client.PaginationMeta, error) {
	req, err := c.client.NewRequest("GET", "/v1/members", nil)
	if err != nil {
		return nil, nil, err
	}

	// Add query parameters
	q := req.URL.Query()

	if lo.PerPage != nil {
		if !(*lo.PerPage >= 1 && *lo.PerPage <= 100) {
			return nil, nil, fmt.Errorf("PerPage must be between 1 and 100")
		}
		q.Add("per_page", strconv.Itoa(*lo.PerPage))
	}

	if lo.Cursor != "" {
		q.Add("cursor", string(lo.Cursor))
	}

	q.Add("search", lo.Search)
	q.Add("role", string(lo.Role))
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
		return &MembersResult{Error: &target}, nil, fmt.Errorf("%s", target.Detail)
	}

	var target []MembersRepsonse
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
