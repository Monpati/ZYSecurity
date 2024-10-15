package form

import "Dexun/utils"

type AreaAccInfo struct {
	OrderId    int64      `json:"order_id"`
	OrderUuid  string     `json:"order_uuid"`
	DomainId   int64      `json:"domain_id"`
	DomainUuid string     `json:"domain_uuid"`
	ProType    string     `json:"pro_type"`
	DaId       int64      `json:"da_id"`
	Regions    utils.JSON `json:"regions"`
	Active     string     `json:"active"`
}

type AreaAccFilterForm struct {
	Active string `json:"active"`
	Field  string `json:"field"`
	Value  string `json:"value"`
	PageForm
}
