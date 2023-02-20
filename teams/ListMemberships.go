package teams

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
	Cursor             string             `json:"cursor"`              // Pagination cursor.
	PerPage            *int               `json:"per_page"`            // [ 1 .. 100 ]
	TeamPermission     TeamPermission     `json:"team_permission"`     // Filter team memberships with a specific team permission
	IncidentPermission IncidentPermission `json:"incident_permission"` // Filter team memberships with a specific incident permission
	MemberId           string             `json:"member_id"`           // Filter team memberships on a specific member
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

	req, err := c.client.NewRequest(ep.Operation, fmt.Sprintf(ep.Path, TeamId), nil)
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

	q.Add("team_permission", string(lo.TeamPermission))
	q.Add("incident_permission", string(lo.IncidentPermission))
	q.Add("member_id", lo.MemberId)
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
