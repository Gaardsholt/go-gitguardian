package incidents

import (
	"fmt"
	"net/http"

	"github.com/Gaardsholt/go-gitguardian/types"
)

type RevokeAccessSecretIncidentOptions struct {
	Email    string `json:"email"`
	MemberID int64  `json:"member_id"`
}

type RevokeAccessSecretIncidentResponse struct {
	ShareURL           string `json:"share_url"`
	IncidentID         int64  `json:"incident_id"`
	FeedbackCollection bool   `json:"feedback_collection"`
	AutoHealing        bool   `json:"auto_healing"`
	Token              string `json:"token"`
	ExpireAt           string `json:"expire_at"`
}

type RevokeAccessSecretIncidentResult struct {
	Result RevokeAccessSecretIncidentResponse `json:"result"`
	Error  *Error                             `json:"error"`
}

func (c *IncidentsClient) RevokeAccessSecretIncident(IncidentId int, lo RevokeAccessSecretIncidentOptions) (bool, error) {
	ep := types.Endpoints["RevokeAccessSecretIncident"]

	r, err := c.client.NewRequest(ep.Operation, fmt.Sprintf(ep.Path, IncidentId), lo)
	if err != nil {
		return false, err
	}

	return r.Response.StatusCode == http.StatusNoContent, nil

}
