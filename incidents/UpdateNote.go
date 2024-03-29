package incidents

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Gaardsholt/go-gitguardian/types"
)

type UpdateNoteOptions struct {
	Comment string `json:"comment"` // Content of the incident note
}

type IncidentUpdateNoteResult struct {
	Result []IncidentUpdateNoteResponse `json:"result"`
	Error  *Error                       `json:"error"`
}

type IncidentUpdateNoteResponse struct {
	ID         int64      `json:"id"`
	IncidentID int64      `json:"incident_id"`
	MemberID   int64      `json:"member_id"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	Comment    string     `json:"comment"`
	IssueID    int64      `json:"issue_id"`
	UserID     int64      `json:"user_id"`
}

func (c *IncidentsClient) UpdateNote(IncidentId int, NoteId int, lo UpdateNoteOptions) (*IncidentUpdateNoteResult, error) {
	ep := types.Endpoints["UpdateNote"]

	req, err := c.client.NewRequest(ep.Operation, fmt.Sprintf(ep.Path, IncidentId, NoteId), lo)
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
		return &IncidentUpdateNoteResult{Error: &target}, fmt.Errorf("%s", target.Detail)
	}

	var target []IncidentUpdateNoteResponse
	decode := json.NewDecoder(r.Body)
	err = decode.Decode(&target)
	if err != nil {
		return nil, err
	}

	return &IncidentUpdateNoteResult{Result: target}, nil
}
