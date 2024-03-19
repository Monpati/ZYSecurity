package form

type AccountInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	TelNum   string `json:"tel_num"`
	CertType string `json:"cert_type"`
}
