package form

type SafeAccInfo struct {
	OrderId    int64    `json:"order_id"`
	OrderUuid  string   `json:"order_uuid"`
	DomainId   int64    `json:"domain_id"`
	DomainUuid string   `json:"domain_uuid"`
	DsId       int64    `json:"ds_id"`
	ProType    string   `json:"pri_type"`
	Password   []string `json:"password"`
	URL        []string `json:"url"`
}

type SafeAccCon struct {
	OrderId    int64  `json:"order_id"`
	OrderUuid  string `json:"order_uuid"`
	DomainId   int64  `json:"domain_id"`
	DomainUuid string `json:"domain_uuid"`
	DsId       int64  `json:"ds_id"`
	ProType    string `json:"pri_type"`
	Password   string `json:"password"`
	URL        string `json:"url"`
}
