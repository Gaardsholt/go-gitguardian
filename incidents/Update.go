package incidents

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Update struct {
	Severity IncidentsListSeverity `json:"severity"`
}

func (c *IncidentsClient) Update(IncidentId int, lo Update) (*IncidentGetResult, error) {
	req, err := c.client.NewRequest("PATCH", fmt.Sprintf("/v1/incidents/secrets/%d", IncidentId), nil)
	if err != nil {
		return nil, err
	}

	// Add query parameters
	q := req.URL.Query()

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
