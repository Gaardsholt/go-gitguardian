package invitations

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type InvitationsListResult struct {
	Result []InvitationsListResponse `json:"result"`
	Error  *Error                    `json:"error"`
}

type InvitationsListResponse struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
}

type ListOptions struct {
	Cursor  string `json:"cursor"`   // Pagination cursor.
	PerPage *int   `json:"per_page"` // Number of items to list per page.	[ 1 .. 100 ]
}

func (c *InvitationsClient) List(lo ListOptions) (*InvitationsListResult, error) {
	req, err := c.client.NewRequest("GET", "/v1/invitations", nil)
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

	if lo.PerPage != nil {
		q.Add("cursor", lo.Cursor)
	}

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
		return &InvitationsListResult{Error: &target}, fmt.Errorf("%s", target.Detail)
	}

	var target []InvitationsListResponse
	decode := json.NewDecoder(r.Body)
	err = decode.Decode(&target)
	if err != nil {
		return nil, err
	}

	return &InvitationsListResult{Result: target}, nil
}
