package members

type Error struct {
	Detail string `json:"detail"`
}

type MembersResult struct {
	Result []MembersRepsonse `json:"result"`
	Error  *Error            `json:"error"`
}

type MembersRepsonse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
