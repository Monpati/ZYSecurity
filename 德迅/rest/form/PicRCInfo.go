package form

type PicRCInfo struct {
	OrderId    int64  `json:"order_id"`
	OrderUuid  string `json:"order_uuid"`
	DomainId   int64  `json:"domain_id"`
	DomainUuid string `json:"domain_uuid"`
	ProType    string `json:"pro_type"`
	DpId       int64  `json:"dp_id"`
	Active     int64  `json:"active"`
}
