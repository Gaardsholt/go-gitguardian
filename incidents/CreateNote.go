package incidents

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type CreateNoteOptions struct {
	Comment string `json:"comment"` // Content of the incident note
}

type IncidentCreateNoteResult struct {
	Result []IncidentCreateNoteResponse `json:"result"`
	Error  *Error                       `json:"error"`
}

type IncidentCreateNoteResponse struct {
	ID         int64       `json:"id"`
	IncidentID int64       `json:"incident_id"`
	MemberID   int64       `json:"member_id"`
	APIToken   interface{} `json:"api_token"`
	CreatedAt  *time.Time  `json:"created_at"`
	UpdatedAt  *time.Time  `json:"updated_at"`
	Comment    string      `json:"comment"`
	IssueID    int64       `json:"issue_id"`
	UserID     int64       `json:"user_id"`
}

func (c *IncidentsClient) CreateNote(IncidentId int, lo CreateNoteOptions) (*IncidentCreateNoteResult, error) {
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(lo)
	if err != nil {
		return nil, err
	}

	req, err := c.client.NewRequest("POST", fmt.Sprintf("/v1/incidents/secrets/%d/notes", IncidentId), b)
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
		return &IncidentCreateNoteResult{Error: &target}, fmt.Errorf("%s", target.Detail)
	}

	var target []IncidentCreateNoteResponse
	decode := json.NewDecoder(r.Body)
	err = decode.Decode(&target)
	if err != nil {
		return nil, err
	}

	return &IncidentCreateNoteResult{Result: target}, nil
}
