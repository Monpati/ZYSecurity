package form

type WafInfo struct {
	DomainId   int64  `json:"domain_id"`
	Active     int64  `json:"active"`
	DDUUID     string `json:"dd_uuid"`
	DomainUUID string `json:"domain_uuid"`
	ProType    string `json:"pro_type"`
	Type       string `json:"type"`
}
