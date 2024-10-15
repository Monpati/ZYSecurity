package form

type DDoSServiceInfo struct {
	ComboId         int64  `json:"combo_id"`
	Uuid            string `json:"uuid"`
	UUserId         int64  `json:"u_user_id"`
	ServerIp        string `json:"server_ip"`
	TcName          string `json:"tc_name"`
	KsMoney         int64  `json:"ks_money"`
	ProductSitename string `json:"product_sitename"`
	StatTime        string `json:"stat_time"`
	EndTime         string `json:"end_time"`
	SiteStart       int64  `json:"site_start"`
	DdosHh          string `json:"ddos_hh"`
	KsStart         int64  `json:"ks_start"`
	DomainNum       int64  `json:"domain_num"`
	RechargeDomain  int64  `json:"recharge_domain"`
	PortNum         int64  `json:"port_num"`
	RechargePort    int64  `json:"recharge_port"`
	UserId          int64  `json:"user_id"`
	Agent           string `json:"agent"`
	ProType         int64  `json:"pro_type"`
}
