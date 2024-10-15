package form

import (
	"Dexun/utils"
)

type DomainsInfo struct {
	Domain  string `json:"domain"`
	OrderId string `json:"order_id"`
}

type DomainInfo struct {
	Domain              string      `json:"domain"`
	DomainUuid          string      `json:"domain_uuid"`
	DomainId            int64       `json:"domain_id"`
	OrderId             *int64      `json:"order_id"`
	PrimaryDomain       string      `json:"primary_domain"`
	DomainStatus        int64       `json:"domain_status"`
	DomainRecord        interface{} `json:"domain_record"`
	FourLayersConfig    utils.JSON  `json:"four_layers_config"`
	CacheFileSizeLimit  int64       `json:"cache_file_size_limit"`
	CacheTotalSizeLimit int64       `json:"cache_total_size_limit"`
	CacheConfig         utils.JSON  `json:"cache_config"`
	CacheActive         int64       `json:"cache_active"`
	WhiteNum            int64       `json:"white_num"`
	UseFlow             int64       `json:"use_flow"`
	CreateTime          string      `json:"createtime"`
	UpdateTime          string      `json:"updatetime"`
	AccessActive        string      `json:"access_active"`
	Grouping            string      `json:"grouping"`
	IsFiling            string      `json:"is_filing"`
	Status              int         `json:"status"`
	UserId              int64       `json:"user_id"`
	Cname               string      `json:"cname"`
	WafSwitch           bool        `json:"waf_switch"`
	WafFile             bool        `json:"waf_file"`
	WafCode             bool        `json:"waf_code"`
	WafSession          bool        `json:"waf_session"`
	WafShellShock       bool        `json:"waf_shellshock"`
	WafZombie           bool        `json:"waf_zombie"`
	WafMetadata         bool        `json:"waf_metadata"`
	WafSql              bool        `json:"waf_sql"`
	WafProxy            bool        `json:"waf_proxy"`
	WafXss              bool        `json:"waf_xss"`
	DDoSDDId            *int64      `json:"ddosdd_id"`
	Field               string      `json:"field"`
	Value               string      `json:"value"`
	PageForm
}

type ConfigListInfo struct {
	DomainId            int64  `json:"domain_id"`
	LoadBalancing       string `json:"load_balancing"`
	OverloadRedirectUrl string `json:"overload_redirect_url"`
	OverloadStatusCode  string `json:"overload_status_code"`
	OverloadType        string `json:"overload_type"`
	Port                string `json:"port"`
	Protocol            string `json:"protocol"`
	Redirect            string `json:"redirect"`
	Server              string `json:"server"`
	UriForward          string `json:"uri_forward"`
	Field               string `json:"field"`
	Value               string `json:"value"`
	PageForm
}

type SourceAddressInfo struct {
	ConfigListId int64  `json:"config_list_id"`
	Address      string `json:"address"`
	Concurrent   string `json:"concurrent"`
	Port         string `json:"port"`
	Protocol     string `json:"protocol"`
	Sni          string `json:"sni"`
	Weight       string `json:"weight"`
	Field        string `json:"field"`
	Value        string `json:"value"`
	PageForm
}

type DomainFilterForm struct {
	Port     string `json:"port"`
	Server   string `json:"server"`
	Protocol string `json:"protocol"`
	Field    string `json:"field"`
	Value    string `json:"value"`
	PageForm
}

type DomainConfigInfo struct {
	Domain              string              `json:"domain"`
	LoadBalancing       string              `json:"load_balancing"`
	OverloadRedirectUrl string              `json:"overload_redirect_url"`
	OverloadStatusCode  string              `json:"overload_status_code"`
	OverloadType        string              `json:"overload_type"`
	Port                string              `json:"port"`
	Protocol            string              `json:"protocol"`
	Redirect            string              `json:"redirect"`
	Server              string              `json:"server"`
	UriForward          string              `json:"uri_forward"`
	SaInfos             []SourceAddressInfo `json:"source_addresses"`
}
