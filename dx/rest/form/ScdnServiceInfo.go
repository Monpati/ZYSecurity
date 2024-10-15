package form

type ScdnServiceInfo struct {
	ComboId         int64  `json:"combo_id"`
	Uuid            string `json:"uuid"`
	TcName          string `json:"tc_name"`
	KsMoney         int64  `json:"ks_money"`
	ProFlow         int64  `json:"pro_flow"`
	DdosHh          string `json:"ddos_hh"`
	DomainNum       int64  `json:"domain_num"`
	ZkMoney         int64  `json:"zk_money"`
	Source          string `json:"source"`
	Status          int    `json:"status"`
	ZyzkMoney       int64  `json:"zyzk_money"`
	ZyksMoney       int64  `json:"zyks_money"`
	WafSwitch       bool   `json:"waf_switch"`
	WafFile         bool   `json:"waf_file"`
	WafCode         bool   `json:"waf_code"`
	WafSession      bool   `json:"waf_session"`
	WafShellShock   bool   `json:"waf_shellshock"`
	WafZombie       bool   `json:"waf_zombie"`
	WafMetadata     bool   `json:"waf_metadata"`
	WafSql          bool   `json:"waf_sql"`
	WafPro          bool   `json:"waf_pro"`
	Months          int64  `json:"months"`
	ActuaFlow       int64  `json:"actua_flow"`
	EndTime         string `json:"end_time"`
	KsStart         int64  `json:"ks_start"`
	ProductSitename string `json:"product_sitename"`
	RechargeFlow    int64  `json:"recharge_flow"`
	RechargeDomain  int64  `json:"recharge_domain"`
	ServerIp        string `json:"server_ip"`
	StartTime       string `json:"stat_time"`
	TotalFlow       int64  `json:"total_flow"`
	UUserId         int64  `json:"u_user_id"`
	SiteStart       int64  `json:"site_stat"`
	UserId          int64  `json:"user_id"`
	PackageId       int64  `json:"package_id"`
}

type Filter struct {
	Field string `json:"field"`
	Value string `json:"value"`
	PageForm
}
