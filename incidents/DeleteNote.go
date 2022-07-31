package incidents

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type IncidentDeleteNoteResult struct {
	Result []IncidentDeleteNoteResponse `json:"result"`
	Error  *Error                       `json:"error"`
}

type IncidentDeleteNoteResponse struct {
	ID         int64      `json:"id"`
	IncidentID int64      `json:"incident_id"`
	MemberID   int64      `json:"member_id"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	Comment    string     `json:"comment"`
	IssueID    int64      `json:"issue_id"`
	UserID     int64      `json:"user_id"`
}

func (c *IncidentsClient) DeleteNote(IncidentId int, NoteId int) (*IncidentDeleteNoteResult, error) {
	req, err := c.client.NewRequest("DELETE", fmt.Sprintf("/v1/incidents/secrets/%d/notes/%d", IncidentId, NoteId), nil)
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
		return &IncidentDeleteNoteResult{Error: &target}, fmt.Errorf("%s", target.Detail)
	}

	var target []IncidentDeleteNoteResponse
	decode := json.NewDecoder(r.Body)
	err = decode.Decode(&target)
	if err != nil {
		return nil, err
	}

	return &IncidentDeleteNoteResult{Result: target}, nil
}
