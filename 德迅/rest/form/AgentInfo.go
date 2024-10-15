package form

type AgentInfo struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	TelNum     string `json:"tel_num"`
	Email      string `json:"email"`
	InviteCode string `json:"invite_code"`
	Role       string `json:"role"`
}

type InviteCode struct {
	Code   string `json:"code"`
	Status int    `json:"status"`
}

type AgentFilterForm struct {
	Username string `json:"username"`
	TelNum   string `json:"tel_num"`
	Field    string `json:"field"`
	Value    string `json:"value"`
	PageForm
}
