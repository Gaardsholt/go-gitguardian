package types

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type Error struct {
	Detail string `json:"detail"`
}

// ContentScanPayload this shouldn't exceed 1MB.
type ContentScanPayload struct {
	Filename string `json:"filename"` // <= 256 characters
	Document string `json:"document"`
}

type ContentScanResult struct {
	Result *ContentScanResponse `json:"result"`
	Error  *Error               `json:"error"`
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
	Validity *string `json:"validity,omitempty"`
}

type Match struct {
	Type       string `json:"type"`
	Match      string `json:"match"`
	IndexStart *int64 `json:"index_start,omitempty"`
	IndexEnd   *int64 `json:"index_end,omitempty"`
	LineStart  *int64 `json:"line_start,omitempty"`
	LineEnd    *int64 `json:"line_end,omitempty"`
}

// MultipleContentScanPayload this shouldn't exceed 2MB.
type MultipleContentScanPayload []ContentScanPayload

type MultipleContentScanResponse []ContentScanResponse

// postStructAsJSON posts the structs as json to the specified url and returns the result in the interface you passes along
func postStructAsJSON(url string, payload interface{}, target interface{}) error {
	var myClient = &http.Client{Timeout: 10 * time.Second}

	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, b)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	r, err := myClient.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	decode := json.NewDecoder(r.Body)
	err = decode.Decode(&target)
	if err != nil {
		return err
	}
	return nil
}
