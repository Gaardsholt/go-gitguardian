package incidents

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Gaardsholt/go-gitguardian/types"
)

type ListNotesOptions struct {
	Cursor  string `json:"-" url:"cursor"`   // Pagination cursor.
	PerPage *int   `json:"-" url:"per_page"` // Number of items to list per page.	[ 1 .. 100 ]
}

type IncidentListNotesResult struct {
	Result []IncidentListNotesResponse `json:"result"`
	Error  *Error                      `json:"error"`
}

type IncidentListNotesResponse struct {
	ID         int64      `json:"id"`
	IncidentID int64      `json:"incident_id"`
	MemberID   int64      `json:"member_id"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	Comment    string     `json:"comment"`
	IssueID    int64      `json:"issue_id"`
	UserID     int64      `json:"user_id"`
}

func (c *IncidentsClient) ListNotes(IncidentId int, lo ListNotesOptions) (*IncidentListNotesResult, error) {
	ep := types.Endpoints["ListNotes"]

	req, err := c.client.NewRequest(ep.Operation, fmt.Sprintf(ep.Path, IncidentId), lo)
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
		return &IncidentListNotesResult{Error: &target}, fmt.Errorf("%s", target.Detail)
	}

	var target []IncidentListNotesResponse
	decode := json.NewDecoder(r.Body)
	err = decode.Decode(&target)
	if err != nil {
		return nil, err
	}

	return &IncidentListNotesResult{Result: target}, nil
}
