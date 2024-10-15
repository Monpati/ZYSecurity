package form

type Purchase struct {
	Username  string `json:"username"`
	UserId    int64  `json:"user_id"`
	ComboUuid string `json:"combo_uuid"`
	ComboId   string `json:"combo_id"`
	Months    string `json:"months"`
	ServiceId *int64 `json:"service_id"`
	DDoSId    *int64 `json:"ddos_id"`
	Agent     string `json:"agent"`
	Count     int64  `json:"count"`
	Recharge  int64  `json:"recharge"`
	Method    string `json:"method"`
}

type BillsFilterForm struct {
	Username string `json:"username"`
	Months   string `json:"months"`
	Agent    string `json:"agent"`
	Field    string `json:"field"`
	Value    string `json:"value"`
	PageForm
}
