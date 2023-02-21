package teams

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

type ListMembershipsOptions struct {
	Cursor             string             `json:"-" url:"cursor"`              // Pagination cursor.
	PerPage            *int               `json:"-" url:"per_page"`            // [ 1 .. 100 ]
	TeamPermission     TeamPermission     `json:"-" url:"team_permission"`     // Filter team memberships with a specific team permission
	IncidentPermission IncidentPermission `json:"-" url:"incident_permission"` // Filter team memberships with a specific incident permission
	MemberId           string             `json:"-" url:"member_id"`           // Filter team memberships on a specific member
}

type ListMembershipsResponse struct {
	ID                 int64              `json:"id"`
	MemberID           int64              `json:"member_id"`
	TeamID             int64              `json:"team_id"`
	TeamPermission     TeamPermission     `json:"team_permission"`
	IncidentPermission IncidentPermission `json:"incident_permission"`
}
type ListMembershipsResult struct {
	Result ListMembershipsResponse `json:"result"`
	Error  *Error                  `json:"error"`
}

func (c *TeamsClient) ListMemberships(TeamId int, lo ListMembershipsOptions) (*ListMembershipsResult, *client.PaginationMeta, error) {
	ep := types.Endpoints["TeamsListMemberships"]

	req, err := c.client.NewRequest(ep.Operation, fmt.Sprintf(ep.Path, TeamId), lo)
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
		return &ListMembershipsResult{Error: &target}, nil, fmt.Errorf("%s", target.Detail)
	}

	var target ListMembershipsResponse
	decode := json.NewDecoder(r.Body)
	err = decode.Decode(&target)
	if err != nil {
		return nil, nil, err
	}

	pagination, err := client.GetPaginationMeta(r)
	if err != nil {
		return nil, nil, err
	}

	return &ListMembershipsResult{Result: target}, pagination, nil
}
