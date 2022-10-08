package auditlogs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Gaardsholt/go-gitguardian/client"
)

type AuditLogsListResult struct {
	Result []AuditLogsListResponse `json:"result"`
	Error  *Error                  `json:"error"`
}

type AuditLogsListResponse struct {
	ID          int64                  `json:"id"`
	Date        time.Time              `json:"date"`
	MemberEmail string                 `json:"member_email"`
	MemberName  string                 `json:"member_name"`
	MemberID    int64                  `json:"member_id"`
	APITokenID  int64                  `json:"api_token_id"`
	IPAddress   string                 `json:"ip_address"`
	TargetIDS   []string               `json:"target_ids"`
	ActionType  string                 `json:"action_type"`
	EventName   string                 `json:"event_name"`
	Data        map[string]interface{} `json:"data"`
}

type AuditLogsListOptions struct {
	Cursor      string     `json:"cursor"`       // Pagination cursor.
	PerPage     *int       `json:"per_page"`     // Number of items to list per page.	[ 1 .. 100 ]
	DateBefore  *time.Time `json:"date_before"`  // Entries found before this date.
	DateAfter   *time.Time `json:"date_after"`   // Entries found after this date.
	EventName   string     `json:"event_name"`   // Entries matching this event name.
	MemberId    *int       `json:"member_id"`    // The id of the member to retrieve.
	MemberName  string     `json:"member_name"`  // Entries matching this member name.
	MemberEmail string     `json:"member_email"` // Entries matching this member email.
}

func (c *AuditLogsClient) List(lo AuditLogsListOptions) (*AuditLogsListResult, *client.PaginationMeta, error) {
	req, err := c.client.NewRequest("GET", "/v1/audit_logs", nil)
	if err != nil {
		return nil, nil, err
	}

	// Add query parameters
	q := req.URL.Query()

	if lo.PerPage != nil {
		if !(*lo.PerPage >= 1 && *lo.PerPage <= 100) {
			return nil, nil, fmt.Errorf("PerPage must be between 1 and 100")
		}
		q.Add("per_page", strconv.Itoa(*lo.PerPage))
	}

	if lo.Cursor != "" {
		q.Add("cursor", lo.Cursor)
	}

	if lo.DateBefore != nil {
		q.Add("date_before", lo.DateBefore.Format("2019-08-30T14:15:22Z"))
	}

	if lo.DateAfter != nil {
		q.Add("date_after", lo.DateAfter.Format("2019-08-30T14:15:22Z"))
	}

	if lo.EventName != "" {
		q.Add("event_name", string(lo.EventName))
	}

	if lo.MemberId != nil {
		q.Add("member_id", strconv.Itoa(*lo.MemberId))
	}

	if lo.MemberName != "" {
		q.Add("member_name", string(lo.MemberName))
	}

	if lo.MemberEmail != "" {
		q.Add("member_email", string(lo.MemberEmail))
	}

	req.URL.RawQuery = q.Encode()

	r, err := c.client.Client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		var target Error
		decode := json.NewDecoder(r.Body)
		err = decode.Decode(&target)
		if err != nil {
			return nil, nil, err
		}
		return &AuditLogsListResult{Error: &target}, nil, fmt.Errorf("%s", target.Detail)
	}

	var target []AuditLogsListResponse
	decode := json.NewDecoder(r.Body)
	err = decode.Decode(&target)
	if err != nil {
		return nil, nil, err
	}

	pagination, err := client.GetPaginationMeta(r)
	if err != nil {
		return nil, nil, err
	}

	return &AuditLogsListResult{Result: target}, pagination, nil
}
