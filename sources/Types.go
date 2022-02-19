package sources

type Error struct {
	Detail string `json:"detail"`
}

type SourcesListResult struct {
	Result []SourcesResponse `json:"result"`
	Error  *Error            `json:"error"`
}

type SourcesGetResult struct {
	Result SourcesResponse `json:"result"`
	Error  *Error          `json:"error"`
}

type SourcesResponse struct {
	ID         int64  `json:"id"`
	URL        string `json:"url"`
	Type       string `json:"type"`
	FullName   string `json:"full_name"`
	Visibility string `json:"visibility"`
}
