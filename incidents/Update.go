package incidents

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Gaardsholt/go-gitguardian/types"
)

type Update struct {
	Severity IncidentsListSeverity `json:"severity" url:"-"`
}

func (c *IncidentsClient) Update(IncidentId int, lo Update) (*IncidentGetResult, error) {
	ep := types.Endpoints["UpdateSecretIncidents"]

	req, err := c.client.NewRequest(ep.Operation, fmt.Sprintf(ep.Path, IncidentId), lo)
	if err != nil {
		return nil, err
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
		return &IncidentGetResult{Error: &target}, fmt.Errorf("%s", target.Detail)
	}

	var target IncidentGetResponse
	decode := json.NewDecoder(r.Body)
	err = decode.Decode(&target)
	if err != nil {
		return nil, err
	}

	return &IncidentGetResult{Result: target}, nil
}
