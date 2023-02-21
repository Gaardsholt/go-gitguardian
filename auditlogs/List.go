package auditlogs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Gaardsholt/go-gitguardian/client"
	"github.com/Gaardsholt/go-gitguardian/types"
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
	Cursor      string     `json:"-" url:"cursor"`       // Pagination cursor.
	PerPage     *int       `json:"-" url:"per_page"`     // Number of items to list per page.	[ 1 .. 100 ]
	DateBefore  *time.Time `json:"-" url:"date_before"`  // Entries found before this date.
	DateAfter   *time.Time `json:"-" url:"date_after"`   // Entries found after this date.
	EventName   string     `json:"-" url:"event_name"`   // Entries matching this event name.
	MemberId    *int       `json:"-" url:"member_id"`    // The id of the member to retrieve.
	MemberName  string     `json:"-" url:"member_name"`  // Entries matching this member name.
	MemberEmail string     `json:"-" url:"member_email"` // Entries matching this member email.
}

func (c *AuditLogsClient) List(lo AuditLogsListOptions) (*AuditLogsListResult, *client.PaginationMeta, error) {
	ep := types.Endpoints["AuditLogsList"]

	req, err := c.client.NewRequest(ep.Operation, ep.Path, lo)
	if err != nil {
		return nil, nil, err
	}

	// Validate query parameters
	if lo.PerPage != nil {
		if !(*lo.PerPage >= 1 && *lo.PerPage <= 100) {
			return nil, nil, fmt.Errorf("PerPage must be between 1 and 100")
		}
	}

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
