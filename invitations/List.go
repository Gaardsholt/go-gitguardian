package invitations

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Gaardsholt/go-gitguardian/types"
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
	Cursor  string `json:"-" url:"cursor"`   // Pagination cursor.
	PerPage *int   `json:"-" url:"per_page"` // Number of items to list per page.	[ 1 .. 100 ]
}

func (c *InvitationsClient) List(lo ListOptions) (*InvitationsListResult, error) {
	ep := types.Endpoints["InvitationsList"]

	req, err := c.client.NewRequest(ep.Operation, ep.Path, lo)
	if err != nil {
		return nil, err
	}

	// Validate query parameters
	if lo.PerPage != nil {
		if !(*lo.PerPage >= 1 && *lo.PerPage <= 100) {
			return nil, fmt.Errorf("PerPage must be between 1 and 100")
		}
	}

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
