package members

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
	Page    *int            `json:"page"`     // Page number.
	PerPage *int            `json:"per_page"` // [ 1 .. 100 ]
	Search  string          `json:"search"`   // Search members based on their name or email.
	Role    MembersListRole `json:"role"`     // Filter members based on their role.
}

func (c *MembersClient) List(lo ListOptions) (*MembersResult, error) {
	req, err := c.client.NewRequest("GET", "/v1/members", nil)
	if err != nil {
		return nil, err
	}

	// Add query parameters
	q := req.URL.Query()

	if lo.PerPage != nil {
		if !(*lo.PerPage >= 1 && *lo.PerPage <= 100) {
			return nil, fmt.Errorf("PerPage must be between 1 and 100")
		}
		q.Add("per_page", strconv.Itoa(*lo.PerPage))
	}

	if lo.Page != nil {
		if !(*lo.Page >= 1) {
			return nil, fmt.Errorf("Page must be geater than zero")
		}
		q.Add("page", strconv.Itoa(*lo.Page))
	}

	q.Add("search", lo.Search)
	q.Add("role", string(lo.Role))
	req.URL.RawQuery = q.Encode()

	r, err := c.client.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		var target Error
		decode := json.NewDecoder(r.Body)
		err = decode.Decode(&target)
		if err != nil {
			return nil, err
		}
		return &MembersResult{Error: &target}, fmt.Errorf("%s", target.Detail)
	}

	var target []MembersRepsonse
	decode := json.NewDecoder(r.Body)
	err = decode.Decode(&target)
	if err != nil {
		return nil, err
	}

	return &MembersResult{Result: target}, nil
}
