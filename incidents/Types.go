package incidents

type Error struct {
	Detail string `json:"detail"`
}

type IncidentListResult struct {
	Result []IncidentListResponse `json:"result"`
	Error  *Error                 `json:"error"`
}

type IncidentListResponse struct {
	ID              int64       `json:"id"`
	Date            string      `json:"date"`
	Detector        Detector    `json:"detector"`
	SecretHash      string      `json:"secret_hash"`
	GitguardianURL  string      `json:"gitguardian_url"`
	Regression      bool        `json:"regression"`
	Status          string      `json:"status"`
	AssigneeEmail   string      `json:"assignee_email"`
	OccurrenceCount int64       `json:"occurrence_count"`
	Occurrences     interface{} `json:"occurrences"`
	IgnoreReason    string      `json:"ignore_reason"`
	IgnoredAt       string      `json:"ignored_at"`
	SecretRevoked   bool        `json:"secret_revoked"`
	Severity        string      `json:"severity"`
	Validity        string      `json:"validity"`
	ResolvedAt      interface{} `json:"resolved_at"`
	ShareURL        string      `json:"share_url"`
}

type Detector struct {
	Name                     string `json:"name"`
	DisplayName              string `json:"display_name"`
	Nature                   string `json:"nature"`
	Family                   string `json:"family"`
	DetectorGroupName        string `json:"detector_group_name"`
	DetectorGroupDisplayName string `json:"detector_group_display_name"`
}
