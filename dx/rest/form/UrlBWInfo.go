package form

type UrlBWInfo struct {
	BwlId      int64    `json:"bwl_id"`
	OrderId    int64    `json:"order_id"`
	DomainId   int64    `json:"domain_id"`
	DomainUuid string   `json:"domain_uuid"`
	OrderUuid  string   `json:"order_uuid"`
	ProType    string   `json:"pro_type"`
	Type       string   `json:"type"`
	Path       []string `json:"path"`
	Method     []string `json:"method"`
	Active     int64    `json:"active"`
	Uuid       string   `json:"uuid"`
}

type UrlBWCon struct {
	BwlId      int64  `json:"bwl_id"`
	OrderId    int64  `json:"order_id"`
	DomainId   int64  `json:"domain_id"`
	DomainUuid string `json:"domain_uuid"`
	OrderUuid  string `json:"order_uuid"`
	ProType    string `json:"pro_type"`
	Type       string `json:"type"`
	Path       string `json:"path"`
	Method     string `json:"method"`
	Active     int64  `json:"active"`
	Uuid       string `json:"uuid"`
}
