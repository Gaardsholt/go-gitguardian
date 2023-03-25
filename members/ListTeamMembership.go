package members

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Gaardsholt/go-gitguardian/client"
	"github.com/Gaardsholt/go-gitguardian/types"
)

type TeamPermission string

const (
	CanManage    TeamPermission = "can_manage"
	CanNotManage TeamPermission = "cannot_manage"
)

type IncidentPermission string

const (
	CanView    IncidentPermission = "can_view"
	CanEdit    IncidentPermission = "can_edit"
	FullAccess IncidentPermission = "full_access"
)

type ListTeamMembershipOptions struct {
	Cursor  string `json:"-" url:"cursor"`   // Pagination cursor.
	PerPage *int   `json:"-" url:"per_page"` // [ 1 .. 100 ]
	TeamId  string `json:"-" url:"team_id"`  // The id of a team to filter on
}

type ListTeamMembershipResult struct {
	Result []ListTeamMembershipResponse `json:"result"`
	Error  *Error                       `json:"error"`
}

type ListTeamMembershipResponse struct {
	ID                 int64              `json:"id"`
	MemberId           int64              `json:"member_id"`
	TeamId             int64              `json:"team_id"`
	TeamPermission     TeamPermission     `json:"team_permission"`
	IncidentPermission IncidentPermission `json:"incident_permission"`
}

func (c *MembersClient) ListTeamMembership(memberId int64, lo ListTeamMembershipOptions) (*ListTeamMembershipResult, *client.PaginationMeta, error) {
	ep := types.Endpoints["ListTeamMembership"]

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
		return &ListTeamMembershipResult{Error: &target}, nil, fmt.Errorf("%s", target.Detail)
	}

	var target []ListTeamMembershipResponse
	decode := json.NewDecoder(r.Body)
	err = decode.Decode(&target)
	if err != nil {
		return nil, nil, err
	}

	pagination, err := client.GetPaginationMeta(r)
	if err != nil {
		return nil, nil, err
	}

	return &ListTeamMembershipResult{Result: target}, pagination, nil
}
