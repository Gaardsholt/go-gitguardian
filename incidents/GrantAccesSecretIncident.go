package incidents

import (
	"fmt"
	"net/http"

	"github.com/Gaardsholt/go-gitguardian/types"
)

type GrantAccessSecretIncidentOptions struct {
	Email    string `json:"email"`
	MemberID int64  `json:"member_id"`
}

type GrantAccessSecretIncidentResponse struct {
	ShareURL           string `json:"share_url"`
	IncidentID         int64  `json:"incident_id"`
	FeedbackCollection bool   `json:"feedback_collection"`
	AutoHealing        bool   `json:"auto_healing"`
	Token              string `json:"token"`
	ExpireAt           string `json:"expire_at"`
}

type GrantAccessSecretIncidentResult struct {
	Result GrantAccessSecretIncidentResponse `json:"result"`
	Error  *Error                            `json:"error"`
}

func (c *IncidentsClient) GrantAccessSecretIncident(IncidentId int, lo GrantAccessSecretIncidentOptions) (bool, error) {
	ep := types.Endpoints["GrantAccessSecretIncident"]

	req, err := c.client.NewRequest(ep.Operation, fmt.Sprintf(ep.Path, IncidentId), lo)
	if err != nil {
		return false, err
	}
	r, err := c.client.Client.Do(req)
	if err != nil {
		return false, err
	}
	defer r.Body.Close()

	return r.StatusCode == http.StatusOK, nil

}
