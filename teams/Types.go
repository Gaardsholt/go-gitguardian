package teams

type Error struct {
	Detail string `json:"detail"`
}

type TeamsResult struct {
	Result []TeamsResponse `json:"result"`
	Error  *Error          `json:"error"`
}

type TeamGetResult struct {
	Result TeamsResponse `json:"result"`
	Error  *Error        `json:"error"`
}

type TeamsResponse struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	IsGlobal       bool   `json:"is_global"`
	GitguardianURL string `json:"gitguardian_url"`
}
