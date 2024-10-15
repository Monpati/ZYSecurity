package form

type PreAccInfo struct {
	DomainUuid string `json:"domain_uuid"`
	DomainId   int64  `json:"domain_id"`
	OrderUuid  string `json:"order_uuid"`
	OrderId    int64  `json:"order_id"`
	DaId       int64  `json:"da_id"`

	// 处理方式，reject:丢弃,pass:回源,block3layer:拉黑,redirect:跳转
	Action string `json:"action"`
	// 开关
	Active int64 `json:"active"`
	// 拉黑时间(秒)：
	BlockTime string `json:"block_time"`
	// 订单uuid
	DDUUID string `json:"dd_uuid"`
	// 域名uuid
	DomainUUID string `json:"domain_uuid"`
	// 跳转地址
	Location string `json:"location"`
	// 订单类型，9.CDN，10.DDoS
	ProType int64 `json:"pro_type"`
	// 已经选择的匹配项类型
	CheckList []string      `json:"check_list"`
	Rule      []RuleElement `json:"rule"`
}

type PreAccConInfo struct {
	DomainUuid string `json:"domain_uuid"`
	DomainId   int64  `json:"domain_id"`
	OrderUuid  string `json:"order_uuid"`
	OrderId    int64  `json:"order_id"`
	DaId       int64  `json:"da_id"`
	Action     string `json:"action"`
	Active     int64  `json:"active"`
	CheckList  string `json:"check_list"`
	MItem      string `json:"m_item"`
	MValue     string `json:"m_value"`
	MOperate   string `json:"m_operate"`
	MValueXs   string `json:"m_value_xs"`
}
