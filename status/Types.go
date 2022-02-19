package status

type Error struct {
	Detail string `json:"detail"`
}

type HealthResult struct {
	Result HealthResponse `json:"result"`
	Error  *Error         `json:"error"`
}

type HealthResponse struct {
	Detail string `json:"detail"`
}
