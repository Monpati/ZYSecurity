package form

type ScdnHeatCon struct {
	OrderUUID   string   `json:"order_uuid"`
	DomainUUID  string   `json:"domain_uuid"`
	ProType     int64    `json:"pro_type"`
	DomainId    int64    `json:"domain_id"`
	CacheConfig []string `json:"cache_config"`
	Url         []string `json:"url"`
	Status      int      `json:"status"`
}

type ScdnHeatInfo struct {
	OrderUUID  string `json:"order_uuid"`
	DomainUUID string `json:"domain_uuid"`
	ProType    int64  `json:"pro_type"`
	DomainId   int64  `json:"domain_id"`
	Url        string `json:"url"`
	Status     int    `json:"status"`
	Active     int    `json:"active"`
}

type HeatFilterForm struct {
	DomainId int64  `json:"domain_id"`
	Id       string `json:"id"`
	Field    string `json:"field"`
	Value    string `json:"value"`
	PageForm
}
