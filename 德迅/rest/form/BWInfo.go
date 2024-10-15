package form

import "Dexun/utils"

type BWInfo struct {
	BwId       int64      `json:"bw_id"`
	OrderId    int64      `json:"order_id"`
	DomainId   int64      `json:"domain_id"`
	DomainUuid string     `json:"domain_uuid"`
	Type       int64      `json:"type"`
	IpList     utils.JSON `json:"ip_list"`
}

type BWInfoFilterForm struct {
	DomainUuid string `json:"domain_uuid"`
	Field      string `json:"field"`
	Value      string `json:"value"`
	PageForm
}
