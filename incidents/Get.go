package incidents

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type GetOptions struct {
	WithOccurrences *int `json:"with_occurrences"` // Retrieve a number of occurrences of this incident.	[ 1 .. 100 ]
}

func (c *IncidentsClient) Get(IncidentId int, lo GetOptions) (*IncidentGetResult, error) {
	req, err := c.client.NewRequest("GET", fmt.Sprintf("/v1/incidents/secrets/%d", IncidentId), nil)
	if err != nil {
		return nil, err
	}

	// Add query parameters
	q := req.URL.Query()

	if lo.WithOccurrences != nil {
		if !(*lo.WithOccurrences >= 1 && *lo.WithOccurrences <= 100) {
			return nil, fmt.Errorf("WithOccurrences must be between 1 and 100")
		}
		q.Add("with_occurrences", strconv.Itoa(*lo.WithOccurrences))
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
