package form

import "Dexun/utils"

type LeechLinkInfo struct {
	DlId       int64      `json:"dl_id"`
	DomainId   int64      `json:"domain_id"`
	DomainUuid string     `json:"domain_uuid"`
	OrderId    int64      `json:"order_id"`
	DdUuid     string     `json:"dd_uuid"`
	ProType    string     `json:"pro_type"`
	Type       string     `json:"type"`
	Active     string     `json:"active"`
	Domains    utils.JSON `json:"domains"`
	AllowEmpty string     `json:"allow_empty"`
}

type LeechLinkFilterForm struct {
	ProType string `json:"pro_type"`
	Field   string `json:"field"`
	Value   string `json:"value"`
	PageForm
}
