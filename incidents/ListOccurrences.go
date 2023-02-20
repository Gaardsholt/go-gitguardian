package incidents

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Gaardsholt/go-gitguardian/types"
)

type ListOccurrencesOptions struct {
	Cursor     string     `json:"cursor"`      // Pagination cursor.
	PerPage    *int       `json:"per_page"`    // Number of items to list per page.	[ 1 .. 100 ]
	DateBefore *time.Time `json:"date_before"` // Entries found before this date.
	DateAfter  *time.Time `json:"date_after"`  // Entries found after this date.
	SourceId   *int       `json:"source_id"`   // Filter on the source ID.
	SourceName string     `json:"source_name"` // Entries matching this source name search.
	IncidentId *int       `json:"incident_id"` // Filter by incident ID.
	Presence   *Presence  `json:"presence"`    // Entries that have the following presence status.
}

type IncidentListOccurrencesResult struct {
	Result []IncidentListOccurrencesResponse `json:"result"`
	Error  *Error                            `json:"error"`
}

type IncidentListOccurrencesResponse struct {
	ID         int64   `json:"id"`
	IncidentID int64   `json:"incident_id"`
	Kind       string  `json:"kind"`
	SHA        string  `json:"sha"`
	Source     Source  `json:"source"`
	AuthorName string  `json:"author_name"`
	AuthorInfo string  `json:"author_info"`
	Date       string  `json:"date"`
	Presence   string  `json:"presence"`
	URL        string  `json:"url"`
	Matches    []Match `json:"matches"`
	Filepath   string  `json:"filepath"`
}

func (c *IncidentsClient) ListOccurrences(lo ListOccurrencesOptions) (*IncidentListOccurrencesResult, error) {
	ep := types.Endpoints["ListOccurrences"]

	req, err := c.client.NewRequest(ep.Operation, ep.Path, nil)
	if err != nil {
		return nil, err
	}

	// Add query parameters
	q := req.URL.Query()

	if lo.PerPage != nil {
		if !(*lo.PerPage >= 1 && *lo.PerPage <= 100) {
			return nil, fmt.Errorf("PerPage must be between 1 and 100")
		}
		q.Add("per_page", strconv.Itoa(*lo.PerPage))
	}

	if lo.DateBefore != nil {
		q.Add("date_before", lo.DateBefore.Format("2019-08-30T14:15:22Z"))
	}

	if lo.DateAfter != nil {
		q.Add("date_after", lo.DateAfter.Format("2019-08-30T14:15:22Z"))
	}

	if lo.SourceId != nil {
		q.Add("source_id", strconv.Itoa(*lo.SourceId))
	}

	if lo.SourceName != "" {
		q.Add("source_name", string(lo.SourceName))
	}

	if lo.IncidentId != nil {
		q.Add("incident_id", strconv.Itoa(*lo.IncidentId))
	}

	if lo.Presence != nil {
		q.Add("presence", string(*lo.Presence))
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
		return &IncidentListOccurrencesResult{Error: &target}, fmt.Errorf("%s", target.Detail)
	}

	var target []IncidentListOccurrencesResponse
	decode := json.NewDecoder(r.Body)
	err = decode.Decode(&target)
	if err != nil {
		return nil, err
	}

	return &IncidentListOccurrencesResult{Result: target}, nil
}
