package incidents

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Gaardsholt/go-gitguardian/types"
)

type ListOccurrencesOptions struct {
	Cursor     string     `json:"-" url:"cursor"`      // Pagination cursor.
	PerPage    *int       `json:"-" url:"per_page"`    // Number of items to list per page.	[ 1 .. 100 ]
	DateBefore *time.Time `json:"-" url:"date_before"` // Entries found before this date.
	DateAfter  *time.Time `json:"-" url:"date_after"`  // Entries found after this date.
	SourceId   *int       `json:"-" url:"source_id"`   // Filter on the source ID.
	SourceName string     `json:"-" url:"source_name"` // Entries matching this source name search.
	IncidentId *int       `json:"-" url:"incident_id"` // Filter by incident ID.
	Presence   *Presence  `json:"-" url:"presence"`    // Entries that have the following presence status.
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

	req, err := c.client.NewRequest(ep.Operation, ep.Path, lo)
	if err != nil {
		return nil, err
	}

	// Validate query parameters
	if lo.PerPage != nil {
		if !(*lo.PerPage >= 1 && *lo.PerPage <= 100) {
			return nil, fmt.Errorf("PerPage must be between 1 and 100")
		}
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
