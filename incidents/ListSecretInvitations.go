package incidents

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Gaardsholt/go-gitguardian/types"
)

type IncidentPermission string

const (
	CanView    IncidentPermission = "can_view"
	CanEdit    IncidentPermission = "can_edit"
	FullAccess IncidentPermission = "full_access"
)

type ListSecretInvitationsOptions struct {
	Cursor             string              `json:"cursor"`              // Pagination cursor.
	InvitationId       *int                `json:"invitation_id"`       // Filter on a specific invitation id.
	IncidentPermission *IncidentPermission `json:"incident_permission"` // Filter accesses with a specific permission.
}

type ListSecretInvitationsResult struct {
	Result []ListSecretInvitationsResponse `json:"result"`
	Error  *Error                          `json:"error"`
}

type ListSecretInvitationsResponse struct {
	InvitationId       int64  `json:"invitation_id"`
	IncidentId         int64  `json:"incident_id"`
	IncidentPermission string `json:"incident_permission"`
}

func (c *IncidentsClient) ListSecretInvitations(IncidentId int, lo ListSecretInvitationsOptions) (*ListSecretInvitationsResult, error) {
	ep := types.Endpoints["ListSecretInvitations"]

	req, err := c.client.NewRequest(ep.Operation, fmt.Sprintf(ep.Path, IncidentId), nil)
	if err != nil {
		return nil, err
	}

	// Add query parameters
	q := req.URL.Query()

	if lo.InvitationId != nil {
		q.Add("invitation_id", strconv.Itoa(*lo.InvitationId))
	}
	if lo.IncidentPermission != nil {
		q.Add("incident_permission", string(*lo.IncidentPermission))
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
		return &ListSecretInvitationsResult{Error: &target}, fmt.Errorf("%s", target.Detail)
	}

	var target []ListSecretInvitationsResponse
	decode := json.NewDecoder(r.Body)
	err = decode.Decode(&target)
	if err != nil {
		return nil, err
	}

	return &ListSecretInvitationsResult{Result: target}, nil
}
