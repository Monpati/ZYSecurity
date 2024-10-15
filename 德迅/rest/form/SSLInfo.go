package form

type SSLInfo struct {
	SSLId       int64  `json:"ssl_id"`
	UUID        string `json:"uuid"`
	PTypeId     int64  `json:"p_type_id"`
	PTypeName   string `json:"p_type_name"`
	SSLName     string `json:"ssl_name"`
	SSLType     string `json:"ssl_type"`
	DomainType  int64  `json:"domain_type"`
	DomainNum   int64  `json:"domain_num"`
	SSLCode     string `json:"ssl_code"`
	KsMoney     int64  `json:"ks_money"`
	EyMoney     int64  `json:"ey_money"`
	PNote       string `json:"p_note"`
	Term        string `json:"term"`
	SType       int64  `json:"s_type"`
	MarketMoney int64  `json:"market_money"`
	SellStatus  int    `json:"sell_status"`
}
