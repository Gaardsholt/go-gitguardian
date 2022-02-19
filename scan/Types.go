package scan

type Error struct {
	Detail string `json:"detail"`
}

// ContentScanPayload this shouldn't exceed 1MB.
type ContentScanPayload struct {
	Filename string `json:"filename"` // <= 256 characters
	Document string `json:"document"`
}

type ContentScanResult struct {
	Result ContentScanResponse `json:"result"`
	Error  *Error              `json:"error"`
}
type MultipleContentScanResult struct {
	Result []ContentScanResponse `json:"result"`
	Error  *Error                `json:"error"`
}

type ContentScanResponse struct {
	PolicyBreakCount int64         `json:"policy_break_count"`
	Policies         []string      `json:"policies"`
	PolicyBreaks     []PolicyBreak `json:"policy_breaks"`
}

type PolicyBreak struct {
	Type     string  `json:"type"`
	Policy   string  `json:"policy"`
	Matches  []Match `json:"matches"`
	Validity string  `json:"validity"`
}

type Match struct {
	Type       string `json:"type"`
	Match      string `json:"match"`
	IndexStart int64  `json:"index_start"`
	IndexEnd   int64  `json:"index_end"`
	LineStart  int64  `json:"line_start"`
	LineEnd    int64  `json:"line_end"`
}

// MultipleContentScanPayload this shouldn't exceed 2MB.
type MultipleContentScanPayload []ContentScanPayload

type MultipleContentScanResponse []ContentScanResponse

type QuotaResult struct {
	Result QuotaResponse `json:"result"`
	Error  *Error        `json:"error"`
}

type QuotaResponse struct {
	Content QuotaContent `json:"content"`
}

type QuotaContent struct {
	Count     int64  `json:"count"`
	Limit     int64  `json:"limit"`
	Remaining int64  `json:"remaining"`
	Since     string `json:"since"`
}
