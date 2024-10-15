package Dexun

import (
	"Dexun/form"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var (
	Username = Api_UserName
	Password = Api_Password
	Address  = DeXun_Address
)

type DeXunBody struct {
	Token         string
	Username      string
	Password      string
	Cert          string
	SCDNOrderList string
	FullName      string
	SmrzBirthday  string
	SmrzCity      string
	SmrzGnum      string
	SmrzSex       string
	SmrzCname     string
	SmrzCnum      string

	CurrentPage int64
	SingleData  Data
	Data        []Datum
	MultiData   []Data
	LastPage    int64
	PerPage     int64
	Total       int64
	List        []List
	AdminList   []AdminList
	MFCount     int64

	CompleteState int64
	DdosHh        string
	DomainNum     int64
	FirewallState int64
	GroupType     int64
	KsMoney       int64
	ProFlow       int64
	ProNote       ProNote
	TcName        string
	UUID          string
	WafState      int64
	ZkMoney       int64

	CallbackType int64
	OrderData    OrderData
	OrderStatus  int64
	OrderUUID    string
	ProductType  int64

	AddKey string

	NodeNote   string
	NodeStatus int64

	AccessActive     string
	CacheActive      int64
	CacheConfig      []string
	ConfigList       []ConfigList
	Createtime       string
	Domain           string
	DomainConfig     DomainConfig
	DomainRecord     interface{}
	DomainStatus     int64
	DomainUUID       string
	FourLayersConfig []FourLayersConfig
	Grouping         string
	IsFilings        string
	PrimaryDomain    string
	Updatetime       string
	UseFlow          int64
	WhiteNum         int64

	CERT      string
	CERTName  string
	Desc      string
	Hsts      int64
	ID        int64
	Key       string
	SSLAlways int64
	Status    int64

	Active string `json:"active"`
	Config Config `json:"config"`

	AppCC      int64 `json:"APP专用防CC策略"`
	CC         int64 `json:"CC防护"`
	IPBlack    int64 `json:"IP黑名单"`
	Referer    int64 `json:"Referer防盗链"`
	UrLBlack   int64 `json:"URL黑名单"`
	WebProtect int64 `json:"Web攻击防护"`
	Other      int64 `json:"其他"`
	AreaAcc    int64 `json:"区域访问限制"`
	SafeAcc    int64 `json:"安全访问控制"`
	PreAcc     int64 `json:"精准访问控制"`

	RequestBandwidthPeak  int64 `json:"request_bandwidth_peak"`
	Requests              int64 `json:"requests"`
	ResponseBandwidthPeak int64 `json:"response_bandwidth_peak"`
	TotalRequestFlows     int64 `json:"total_request_flows"`
	TotalResponseFlows    int64 `json:"total_response_flows"`
	UnidentifiedAttack    int64 `json:"unidentified_attack"`

	Loadcsr    string `json:"loadcsr"`
	Loaddomain string `json:"loaddomain"`
	Loadkey    string `json:"loadkey"`

	DDInfo DDInfo `json:"dd_info"`

	SSLOrderInfo []SSL `json:"ssl_order_info"`
}

func NewDeXun() *DeXunBody {
	return &DeXunBody{
		Username:     Username,
		Password:     Password,
		FullName:     "",
		SmrzBirthday: "",
		SmrzCity:     "",
		SmrzGnum:     "",
		SmrzSex:      "",
		SmrzCname:    "",
		SmrzCnum:     "",
	}
}

// json
func (p *DeXunBody) SendRequest(method, path string, params url.Values, param interface{}, responseBody interface{}, contentType string) error {
	url := fmt.Sprintf("%s/%s", Address, path)

	var req *http.Request
	var err error
	if params != nil && param == nil {
		switch contentType {
		case "application/json":
			jsonData, err := json.Marshal(params)
			if err != nil {
				return err
			}
			req, err = http.NewRequest(method, url, bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")
		case "application/x-www-form-urlencoded":
			formData := params.Encode()
			req, err = http.NewRequest(method, url, strings.NewReader(formData))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		default:
			return fmt.Errorf("unsupported content type: %s", contentType)
		}
	} else if params == nil && param != nil {
		switch contentType {
		case "application/json":
			jsonData, err := json.Marshal(param)
			if err != nil {
				return err
			}
			req, err = http.NewRequest(method, url, bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")
		case "application/x-www-form-urlencoded":
			formData := params.Encode()
			req, err = http.NewRequest(method, url, strings.NewReader(formData))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		default:
			return fmt.Errorf("unsupported content type: %s", contentType)
		}
	} else {
		req, err = http.NewRequest(method, url, nil)
	}
	if err != nil {
		return err
	}

	req.Header.Set("agent-token", p.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	if responseBody != nil {
		var tmp Response
		err = json.Unmarshal(body, &tmp)
		if tmp.Message == "您的套餐暂不支持此功能，请升级套餐。" {
			return errors.New("您的套餐暂不支持此功能，请升级套餐。")
		}
		err = json.Unmarshal(body, responseBody)
		if err != nil {
			return err
		}

		//if responseBody.Message == ""
	}
	return nil
}

// ApiLogin 接口登录获取token
func (p *DeXunBody) ApiLogin() string {
	var resp Response
	//request := url.Values{"username": {p.Username}, "password": {p.Password}}
	//tmp, _ := http.PostForm(p.LoginUrl, request)
	request := url.Values{"username": {Username}, "password": {Password}}
	tmp, err := http.PostForm("https://apiagent.dexunyun.com/agent/user_manager/apicheckIn", request)
	if err != nil {
		return err.Error()
	}
	//defer tmp.Body.Close()
	bodyC, _ := ioutil.ReadAll(tmp.Body)
	err = json.Unmarshal(bodyC, &resp)
	if err != nil {
		return err.Error()
	} else if resp.Status != 200 {
		return resp.Message
	}
	p.Token = resp.Data.Token
	return ""
}

// SwitchCert 开启实名认证接口
func (p *DeXunBody) SwitchCert() error {
	var result Response
	params := url.Values{"enable_smrz": {"1"}}
	if err := p.SendRequest("POST", "authentication.real/switch", params, nil, &result, "application/x-www-form-urlencoded"); err != nil {
		return err
	}
	return nil
}

// GetPersonCertDetails 个人实名认证
func (p *DeXunBody) GetPersonCertDetails(gname, gnum string) error {
	var result Response
	var info form.Cert

	params := url.Values{info.Gname: {gname}, info.Gnum: {gnum}}
	if err := p.SendRequest("POST", "authentication.real/nameauthgr", params, nil, &result, "application/x-www-form-urlencoded"); err != nil {
		return err
	}
	p.FullName = result.Data.FullName
	p.SmrzBirthday = result.Data.SmrzBirthday
	p.SmrzCity = result.Data.SmrzCity
	p.SmrzGnum = result.Data.SmrzGnum
	p.SmrzSex = result.Data.SmrzSex
	return nil
}

// GetCorpCertDetails 企业实名认证
func (p *DeXunBody) GetCorpCertDetails(compamyname, cnum, cname, gnum, scity string) error {
	var result Response
	var info form.Cert
	params := url.Values{info.Cname: {cname}, info.Cnum: {cnum}, info.Compamyname: {compamyname}, info.Gnum: {gnum}, info.Scity: {scity}}
	if err := p.SendRequest("POST", "authentication.real/nameauthgr", params, nil, &result, "application/x-www-form-urlencoded"); err != nil {
		return err
	}
	p.FullName = result.Data.FullName
	p.SmrzGnum = result.Data.SmrzGnum
	p.SmrzCname = result.Data.SmrzCname
	p.SmrzCnum = result.Data.SmrzCnum

	return nil
}

// GetSCDNList SCDN订单列表
func (p *DeXunBody) GetSCDNList(keywords, listrows, ks_money string, page int64) []Datum {
	var result Response

	params := form.SCDN{
		Keywords: keywords,
		Page:     page,
		ListRows: listrows,
		Order: form.Order{
			KsMoney: ks_money,
		},
	}
	if err := p.SendRequest("POST", "cdn.dd/lists", nil, params, &result, "application/json"); err != nil {
		return nil
	}

	p.Total = result.Data.Total
	p.PerPage = result.Data.PerPage
	p.CurrentPage = result.Data.CurrentPage
	p.LastPage = result.Data.LastPage
	p.Data = result.Data.Data

	return p.Data
}

// GetComboList SCDN套餐列表
func (p *DeXunBody) GetComboList() []Data {
	var result MultiResponse
	var end []Data

	params := url.Values{}
	if err := p.SendRequest("POST", "cdn.tc/lists", params, nil, &result, "application/x-www-form-urlencoded"); err != nil {
		return nil
	}
	end = result.Data
	return end
}

// GetFlowPackageLists 流量增量包列表
func (p *DeXunBody) GetFlowPackageLists() []Data {
	var result MultiResponse
	var end []Data

	params := url.Values{}
	if err := p.SendRequest("POST", "cdn.flow/lists", params, nil, &result, "application/x-www-form-urlencoded"); err != nil {
		return nil
	}
	end = result.Data
	for i, _ := range end {
		end[i].Type = "flow"
	}
	return end
}

// GetDDPackageLists 域名增量包列表
func (p *DeXunBody) GetDDPackageLists() []Data {
	var result MultiResponse
	var end []Data

	params := url.Values{}
	if err := p.SendRequest("POST", "cdn.dd/domainbaoList", params, nil, &result, "application/x-www-form-urlencoded"); err != nil {
		return nil
	}
	end = result.Data
	for i, _ := range end {
		end[i].Type = "dd"
	}
	return end
}

// GetBillCallBack 订单回调
func (p *DeXunBody) GetBillCallBack(uuid []string) []Data {
	var result MultiResponse

	params := form.CallBack{
		UUID: uuid,
	}
	if err := p.SendRequest("POST", "callback/orderHandle", nil, params, &result, "application/json"); err != nil {
		return nil
	}
	return result.Data
}

// CreateSCDNOrder SCDN下单接口
func (p *DeXunBody) CreateSCDNOrder(tc_uuid string, months int64) string {
	var result Response

	params := form.OrderConsume{UUID: tc_uuid, Months: months}
	if err := p.SendRequest("POST", "cdn.dd/platformCreate", nil, params, &result, "application/json"); err != nil {
		return ""
	}
	p.OrderUUID = result.Data.UUID
	p.AddKey = result.Data.AddKey

	return p.OrderUUID
}

// OrderRenewal 订单按月续费
func (p *DeXunBody) OrderRenewal(uuid string, months int64) error {
	var result Response

	params := form.OrderConsume{UUID: uuid, Months: months}
	if err := p.SendRequest("POST", "cdn.dd/xufei", nil, params, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// GetMultiThreat 前端异步批量获取今日威胁
func (p *DeXunBody) GetMultiThreat(instance_name_list []string) []Data {
	var result MultiResponse

	params := form.InstanceList{InstanceNameList: instance_name_list}
	if err := p.SendRequest("POST", "cdn.dd/getDayWaf", nil, params, &result, "application/json"); err != nil {
		return nil
	} else {
		return result.Data
	}
}

// Upgrade 套餐升级
func (p *DeXunBody) Upgrade(dd_uuid, new_tc_uuid string) error {
	var result Response

	params := url.Values{"dd_uuid": {dd_uuid}, "new_tc_uuid": {new_tc_uuid}}
	if err := p.SendRequest("POST", "cdn.dd/platformUpgrade", params, nil, &result, "application/x-www-form-urlencoded"); err != nil {
		return err
	} else {
		return nil
	}
}

// PurchaseFlow 购买流量包
func (p *DeXunBody) PurchaseFlow(flow_uuid, dd_uuid string) error {
	var result Response

	params := url.Values{"flow_uuid": {flow_uuid}, "dd_uuid": {dd_uuid}}
	if err := p.SendRequest("POST", "cdn.dd/buyFlow", params, nil, &result, "application/x-www-form-urlencoded"); err != nil {
		return err
	} else {
		return nil
	}
}

// PurchaseDomains 购买域名增量包
func (p *DeXunBody) PurchaseDomains(domainbao_uuid, dd_uuid string) error {
	var result Response

	params := url.Values{"domainbao_uuid": {domainbao_uuid}, "dd_uuid": {dd_uuid}}
	if err := p.SendRequest("POST", "cdn.dd/buyDomainbao", nil, params, &result, "application/x-www-form-urlencoded"); err != nil {
		return err
	} else {
		return nil
	}
}

// LoginConPanel 登录控制面板
func (p *DeXunBody) LoginConPanel(dd_uuid string) error {
	var result Response

	params := url.Values{"dd_uuid": {dd_uuid}}
	if err := p.SendRequest("POST", "cdn.dd/controlPanel", params, nil, &result, "application/x-www-form-urlencoded"); err != nil {
		return err
	} else {
		return nil
	}
}

// EnableSite 启用站点
func (p *DeXunBody) EnableSite(dd_uuid string) error {
	var result Response

	params := url.Values{"dd_uuid": {dd_uuid}}

	if err := p.SendRequest("POST", "cdn.dd/enableSite", params, nil, &result, "application/x-www-form-urlencoded"); err != nil {
		return err
	} else {
		return nil
	}
}

// CleanCache 清理接口缓存，所有模块都走这个接口
func (p *DeXunBody) CleanCache(domain_uuid, dd_uuid string, pro_type int64) error {
	var result Response

	params := form.Domain{
		DDUUID:     dd_uuid,
		DomainUUID: domain_uuid,
		ProType:    pro_type,
	}

	if err := p.SendRequest("POST", "cdncp.domain/clean", nil, params, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// GetDomainInfo 获取域名信息，所有模块都走这个接口
func (p *DeXunBody) GetDomainInfo(dd_uuid, domain_uuid string, pro_type int64) error {
	var result Response

	params := form.Domain{
		DDUUID:     dd_uuid,
		DomainUUID: domain_uuid,
		ProType:    pro_type,
	}
	if err := p.SendRequest("POST", "cdncp.domain/info", nil, params, &result, "application/json"); err != nil {
		return err
	} else {
		p.AccessActive = result.Data.AccessActive
		p.CacheActive = result.Data.CacheActive
		p.CacheConfig = result.Data.CacheConfig
		p.ConfigList = result.Data.ConfigList
		p.Createtime = result.Data.Createtime
		p.Domain = result.Data.Domain
		p.DomainConfig.CacheTotalSizeLimit = result.Data.DomainConfig.CacheTotalSizeLimit
		p.DomainConfig.CacheFileSizeLimit = result.Data.DomainConfig.CacheFileSizeLimit
		p.DomainRecord = result.Data.DomainRecord
		p.DomainStatus = result.Data.DomainStatus
		p.DomainUUID = result.Data.DomainUUID
		p.FourLayersConfig = result.Data.FourLayersConfig
		p.Grouping = result.Data.Grouping
		p.IsFilings = result.Data.IsFilings
		p.PrimaryDomain = result.Data.PrimaryDomain
		p.Updatetime = result.Data.Updatetime
		p.UseFlow = result.Data.UseFlow
		p.WhiteNum = result.Data.WhiteNum
		return nil
	}
}

// AddDomains 添加域名，所有模块都走这个接口，
func (p *DeXunBody) AddDomains(domain, dd_uuid string, pro_type int64) error {
	var result Response

	params := form.Domain{
		DDUUID:  dd_uuid,
		Domain:  domain,
		ProType: pro_type,
	}

	if err := p.SendRequest("POST", "cdncp.domain/domainAdd", nil, params, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// GetDomainsLists 获取域名的列表，所有模块都走这个接口
func (p *DeXunBody) GetDomainsLists(dd_uuid string, pro_type int64) error {
	var result Response

	params := form.DomainsLists{
		DDUUID:  dd_uuid,
		ProType: pro_type,
		Page:    1,
		Limit:   15,
	}
	if err := p.SendRequest("POST", "cdncp.domain/index", nil, params, &result, "application/json"); err != nil {
		return err
	} else {
		p.List = result.Data.List
		p.Total = result.Data.Total
		return nil
	}
}

// GetQuickStatus 获取加速状态，所有模块都走这个接口
func (p *DeXunBody) GetQuickStatus(dd_uuid, domain_uuid string, pro_type int64) error {
	var result Response

	params := form.Domain{
		DDUUID:     dd_uuid,
		DomainUUID: domain_uuid,
		ProType:    pro_type,
	}
	if err := p.SendRequest("POST", "cdncp.domain/quicken", nil, params, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// UpdateHttpDevices http更新设备，即更新域名缓存的相关网络配置，所有模块都走这个接口
func (p *DeXunBody) UpdateHttpDevices(info form.DomainDevices) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.domain/httpEdit", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// UpdatePreheat 域名开启关闭缓存预热，所有模块都走这个接口
func (p *DeXunBody) UpdatePreheat(info *form.DomainHeat) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.domain/preheat", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// EditPreheat 域名缓存预热路由编辑，所有模块都走这个接口
func (p *DeXunBody) EditPreheat(info *form.DomainHeatUpdate) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.domain/editPreheat", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// DeleteDomains 域名删除，所有模块都走这个接口
func (p *DeXunBody) DeleteDomains(dd_uuid string, pro_type int64, domain_uuid string) error {
	var result Response

	params := form.Domain{
		DDUUID:     dd_uuid,
		ProType:    pro_type,
		DomainUUID: domain_uuid,
	}
	if err := p.SendRequest("POST", "cdncp.domain/delete", nil, params, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// CacheModelAdd 缓存模版添加
func (p *DeXunBody) CacheModelAdd(info *form.Cache) error {
	var result Response

	params := form.Cache{
		Active:          info.Active,
		Cacheextensions: info.Cacheextensions,
		Cachemode:       info.Cachemode,
		Cachepath:       info.Cachepath,
		Cachereg:        info.Cachereg,
		DDUUID:          info.DDUUID,
		ProType:         info.ProType,
		Timeout:         info.Timeout,
		Urlmode:         info.Urlmode,
		Weight:          info.Weight,
	}
	if err := p.SendRequest("POST", "cdncp.cache_cert/templateAdd", nil, params, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// CacheModelLists 缓存模版列表
func (p *DeXunBody) CacheModelLists(info *form.Cache) error {
	var result Response

	params := form.Cache{
		DDUUID:  info.DDUUID,
		ProType: info.ProType,
	}
	if err := p.SendRequest("POST", "cdncp.cache_cert/templateLists", nil, params, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// CacheModelUpdate 缓存模版修改
func (p *DeXunBody) CacheModelUpdate(info *form.Cache) error {
	var result Response

	params := form.Cache{
		Active:          info.Active,
		Cacheextensions: info.Cacheextensions,
		Cachemode:       info.Cachemode,
		Cachepath:       info.Cachepath,
		Cachereg:        info.Cachereg,
		DDUUID:          info.DDUUID,
		ProType:         info.ProType,
		Timeout:         info.Timeout,
		Urlmode:         info.Urlmode,
		Weight:          info.Weight,
	}
	if err := p.SendRequest("POST", "cdncp.cache_cert/templateSave", nil, params, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// CacheModelDelete 缓存模版删除
func (p *DeXunBody) CacheModelDelete(info *form.Cache) error {
	var result Response

	params := form.Cache{
		DDUUID:  info.DDUUID,
		ProType: info.ProType,
	}
	if err := p.SendRequest("POST", "cdncp.cache_cert/templateDel", nil, params, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// CacheAdd 缓存添加
func (p *DeXunBody) CacheAdd(info *form.Cache) error {
	var result Response

	params := form.Cache{
		Active:          info.Active,
		Cacheextensions: info.Cacheextensions,
		Cachemode:       info.Cachemode,
		Cachepath:       info.Cachepath,
		Cachereg:        info.Cachereg,
		DDUUID:          info.DDUUID,
		ProType:         info.ProType,
		Timeout:         info.Timeout,
		Urlmode:         info.Urlmode,
		Weight:          info.Weight,
		DomainUUID:      info.DomainUUID,
	}
	if err := p.SendRequest("POST", "cdncp.cache/add", nil, params, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// CacheLists 缓存列表
func (p *DeXunBody) CacheLists(info *form.Cache) error {
	var result Response

	params := form.Cache{
		DDUUID:     info.DDUUID,
		ProType:    info.ProType,
		DomainUUID: info.DomainUUID,
	}
	if err := p.SendRequest("POST", "cdncp.cache/index", nil, params, &result, "application/json"); err != nil {
		return err
	} else {
		p.List = result.Data.List
		p.Total = result.Data.Total
		return nil
	}
}

// CacheUpdate 缓存更新
func (p *DeXunBody) CacheUpdate(info *form.Cache) error {
	var result Response

	params := form.Cache{
		Active:          info.Active,
		Cacheextensions: info.Cacheextensions,
		Cachemode:       info.Cachemode,
		Cachepath:       info.Cachepath,
		Cachereg:        info.Cachereg,
		DDUUID:          info.DDUUID,
		ProType:         info.ProType,
		Timeout:         info.Timeout,
		Urlmode:         info.Urlmode,
		Weight:          info.Weight,
		DomainUUID:      info.DomainUUID,
	}
	if err := p.SendRequest("POST", "cdncp.cache/edit", nil, params, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// CacheDelete 缓存删除
func (p *DeXunBody) CacheDelete(info *form.Cache) error {
	var result Response

	params := form.Cache{
		DDUUID:     info.DDUUID,
		ProType:    info.ProType,
		DomainUUID: info.DomainUUID,
	}
	if err := p.SendRequest("POST", "cdncp.cache/delete", nil, params, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// CacheDomainLists 域名菜单
func (p *DeXunBody) CacheDomainLists(info *form.Cache) error {
	var result Response

	params := form.Cache{
		DDUUID:  info.DDUUID,
		ProType: info.ProType,
	}
	if err := p.SendRequest("POST", "cdncp.domain/domain", nil, params, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// CertUpdate 更新证书
func (p *DeXunBody) CertUpdate(info *form.Certification) error {
	var result Response

	params := form.Certification{
		CERT:       info.CERT,
		CERTName:   info.CERTName,
		DDUUID:     info.DDUUID,
		DomainUUID: info.DomainUUID,
		Hsts:       info.Hsts,
		Key:        info.Key,
		ProType:    info.ProType,
		SSLAlways:  info.SSLAlways,
		Status:     info.Status,
	}
	if err := p.SendRequest("POST", "cdncp.cert/edit", nil, params, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// CertGetDomains 证书获取域名
func (p *DeXunBody) CertGetDomains(info *form.Certification) error {
	var result Response

	params := form.Certification{
		DDUUID:     info.DDUUID,
		DomainUUID: info.DomainUUID,
		ProType:    info.ProType,
	}
	if err := p.SendRequest("POST", "cdncp.cert/getdata", nil, params, &result, "application/json"); err != nil {
		return err
	} else {
		p.CERT = result.Data.CERT
		p.CERTName = result.Data.CERTName
		p.Createtime = result.Data.Createtime
		p.Desc = result.Data.Desc
		p.DomainUUID = result.Data.DomainUUID
		p.Hsts = result.Data.Hsts
		p.Key = result.Data.Key
		p.SSLAlways = result.Data.SSLAlways
		p.Status = result.Data.Status
		p.Updatetime = result.Data.Updatetime
		return nil
	}
}

// CustomizedUpdate 定制页面更新
func (p *DeXunBody) CustomizedUpdate(info *form.Customized) error {
	var result Response

	params := form.Customized{
		Content:    info.Content,
		DDUUID:     info.DDUUID,
		DomainUUID: info.DomainUUID,
		PageType:   info.PageType,
		ProType:    info.ProType,
		Type:       info.Type,
	}
	if err := p.SendRequest("POST", "cdncp/customized/edit", nil, params, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// CustomizedLists 定制页面更新
func (p *DeXunBody) CustomizedLists(info *form.Customized) error {
	var result Response

	params := form.Customized{
		DDUUID:     info.DDUUID,
		DomainUUID: info.DomainUUID,
		ProType:    info.ProType,
		Type:       info.Type,
	}
	if err := p.SendRequest("POST", "cdncp/customized/index", nil, params, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// IpBWUpdate 防火墙修改黑白名单ip
func (p *DeXunBody) IpBWUpdate(info *form.IpBW) error {
	var result Response

	params := form.IpBW{
		DDUUID:     info.DDUUID,
		DomainUUID: info.DomainUUID,
		IPList:     info.IPList,
		ProType:    info.ProType,
		Type:       info.Type,
	}
	if err := p.SendRequest("POST", "cdncp.firewall/domainIpsEdit", nil, params, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// IppBWDelete 防火墙删除黑白名单ip
func (p *DeXunBody) IppBWDelete(info *form.IpBW) error {
	var result Response

	params := form.IpBW{
		DDUUID:     info.DDUUID,
		DomainUUID: info.DomainUUID,
		IPList:     info.IPList,
		ProType:    info.ProType,
		Type:       info.Type,
	}
	if err := p.SendRequest("POST", "cdncp.firewall/domainIpsDel", nil, params, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// ClearIpBW 防火墙清空黑白名单ip
func (p *DeXunBody) ClearIpBW(info *form.IpBW) error {
	var result Response

	params := form.IpBW{
		DDUUID:     info.DDUUID,
		DomainUUID: info.DomainUUID,
		IPList:     info.IPList,
		ProType:    info.ProType,
		Type:       info.Type,
	}
	if err := p.SendRequest("POST", "cdncp.firewall/domainIpsDel", nil, params, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// GetIpBWInsLists 防火墙获取黑白名单列表(实例)
func (p *DeXunBody) GetIpBWInsLists(info *form.IpBW) error {
	var result Response

	params := form.IpBW{
		DDUUID:  info.DDUUID,
		ProType: info.ProType,
	}
	if err := p.SendRequest("POST", "cdncp.firewall/domainIpsIndex", nil, params, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// GetIpBWSinLists 防火墙获取黑白名单信息(单域名)
func (p *DeXunBody) GetIpBWSinLists(info *form.IpBW) error {
	var result Response

	params := form.IpBW{
		DDUUID:     info.DDUUID,
		DomainUUID: info.DomainUUID,
		ProType:    info.ProType,
		Type:       info.Type,
	}
	if err := p.SendRequest("POST", "cdncp.firewall/domainIpsGetdata", nil, params, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// UrlBWUpdate 防火墙修改URl黑白名单
func (p *DeXunBody) UrlBWUpdate(info *form.UrlBW) error {
	var result Response

	params := form.UrlBW{
		DDUUID:     info.DDUUID,
		DomainUUID: info.DomainUUID,
		Type:       info.Type,
		URLList:    info.URLList,
		ProType:    info.ProType,
	}
	if err := p.SendRequest("POST", "cdncp.firewall/domainUrlsEdit", nil, params, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// DeleteUrlBW 防火墙删除URL黑白名单
func (p *DeXunBody) DeleteUrlBW(info *form.UrlBW) error {
	var result Response

	params := form.UrlBW{
		DDUUID:     info.DDUUID,
		DomainUUID: info.DomainUUID,
		Type:       info.Type,
		URLList:    info.URLList,
		ProType:    info.ProType,
	}
	if err := p.SendRequest("POST", "cdncp.firewall/domainUrlsDel", nil, params, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// GetUrlBWInfo 防火墙获取黑白名单信息
func (p *DeXunBody) GetUrlBWInfo(info *form.UrlBW) error {
	var result Response

	params := form.UrlBW{
		DDUUID:     info.DDUUID,
		DomainUUID: info.DomainUUID,
		Type:       info.Type,
		URLList:    info.URLList,
		ProType:    info.ProType,
	}
	if err := p.SendRequest("POST", "cdncp.firewall/domainUrlsGetdata", nil, params, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// GetUrlBWLists 防火墙获取黑白名单列表
func (p *DeXunBody) GetUrlBWLists(info *form.UrlBW) error {
	var result Response

	params := form.UrlBW{
		DDUUID:  info.DDUUID,
		ProType: info.ProType,
	}
	if err := p.SendRequest("POST", "cdncp.firewall/domainUrlsIndex", nil, params, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// GetCCInfo 获取CC配置
func (p *DeXunBody) GetCCInfo(info *form.Base) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.firewall/domainFirewallGetcc", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// UpdateCCInfo 更新防CC配置
func (p *DeXunBody) UpdateCCInfo(info *form.CC) error {
	var result Response

	//params := form.CC{
	//	DomainUUID: info.DomainUUID,
	//	Config:     info.Config,
	//	Active:     info.Active,
	//	UseDefault: info.UseDefault,
	//	DDUUID:     info.DDUUID,
	//	ProType:    info.ProType,
	//}
	if err := p.SendRequest("POST", "cdncp.firewall/domainFirewallEditcc", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// UpdateAreaCon 更新区域访问限制
func (p *DeXunBody) UpdateAreaCon(info *form.AreaAccCon) error {
	var result Response

	params := form.AreaAccCon{
		Active:     info.Active,
		DDUUID:     info.DDUUID,
		DomainUUID: info.DomainUUID,
		ProType:    info.ProType,
		Regions:    info.Regions,
	}
	if err := p.SendRequest("POST", "cdncp.waf/domainFirewallEditarea", nil, params, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// GetAreaCon 获取区域访问限制
func (p *DeXunBody) GetAreaCon(info *form.AreaAccCon) error {
	var result Response

	params := form.AreaAccCon{
		DomainUUID: info.DomainUUID,
		DDUUID:     info.DDUUID,
		ProType:    info.ProType,
	}
	if err := p.SendRequest("POST", "cdncp.waf/domainFirewallGetarea", nil, params, &result, "application/json"); err != nil {
		return err
	} else {
		p.Active = result.Data.Active
		p.Config.Regions = result.Data.Config.Regions
		return nil
	}
}

func (p *DeXunBody) DeleteAreaCon(info *form.AreaAccCon) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.waf/domainFirewallDelarea", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// GetLeechLink 获取防盗链
func (p *DeXunBody) GetLeechLink(info *form.Base) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.firewall/domainFirewallGetreferer", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// UpdateLeechLink 更新防盗链配置
func (p *DeXunBody) UpdateLeechLink(info *form.LeechLink) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.firewall/domainFirewallEditreferer", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// GetPicRC 获取内容安全风控-图片
func (p *DeXunBody) GetPicRC(info *form.Base) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.waf/domainFirewallGetriskimg", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// UpdatePicRC 更新内容安全风控-图片
func (p *DeXunBody) UpdatePicRC(info *form.PicRc) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.waf/domainFirewallEditriskimg", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// GetWordsRC 获取内容安全风控-关键字
func (p *DeXunBody) GetWordsRC(info *form.Base) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.waf/domainFirewallGetriskword", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// UpdateWordsRC 更新内容安全风控-关键字
func (p *DeXunBody) UpdateWordsRC(info *form.WordsRc) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.waf/domainFirewallEditriskword", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// DeleteWordsRC 删除内容安全风控-关键字
func (p *DeXunBody) DeleteWordsRC(info *form.WordsRc) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.waf/domainFirewallDelriskword", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// GetSafeCon 获取安全访问控制
func (p *DeXunBody) GetSafeCon(info *form.Base) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.firewall/domainFirewallGetSafe", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// UpdateSafeCon 设置安全访问控制
func (p *DeXunBody) UpdateSafeCon(info *form.SafeCon) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.firewall/domainFirewallEditsafe", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// CleanSafeCon 删除安全访问控制
func (p *DeXunBody) CleanSafeCon(info *form.SafeAcc) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.firewall/domainFirewallEditsafe", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// SwitchPreAccCon 精准访问开关
func (p *DeXunBody) SwitchPreAccCon(info *form.SwitchPreAcc) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.firewall/domainFirewallRenewPreciseOpen", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// UpdateSinglePreAcc 更新单个精准访问控制
func (p *DeXunBody) UpdateSinglePreAcc(info *form.PreAcc) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.firewall/domainFirewallRenewprecise", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// GetPreAcc 获取精准访问控制
func (p *DeXunBody) GetPreAcc(info *form.Base) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.firewall/domainFirewallGetprecise", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// UpdatePreAccCon 添加精准访问控制
func (p *DeXunBody) UpdatePreAccCon(info *form.PreAcc) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.firewall/domainFirewallEditprecise", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// DeletePreAcc 删除精准访问控制
func (p *DeXunBody) DeletePreAcc(info *form.PreAccDel) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.firewall/domainFirewallDelprecise", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// UpdateTotalWaf 设置WAF总体开关
func (p *DeXunBody) UpdateTotalWaf(info *form.WafStatus) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.waf/domainWafEditwaf", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// GetTotalWaf 获取WAF总体开关
func (p *DeXunBody) GetTotalWaf(info *form.Base) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.waf/domainWafGetwaf", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// UpdateFileWaf 设置WAF文件包含开关
func (p *DeXunBody) UpdateFileWaf(info *form.WafStatus) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.waf/domainWafEditfile", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// GetFileWaf 获取WAF文件包含
func (p *DeXunBody) GetFileWaf(info *form.Base) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.waf/domainWafGetfile", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// UpdateCodeWaf 设置WAF代码注入开关
func (p *DeXunBody) UpdateCodeWaf(info *form.WafStatus) error {
	var result Response

	if err := p.SendRequest("POST", "/agent/cdncp.waf/domainWafEditcode", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// GetCodeWaf 获取WAF代码注入
func (p *DeXunBody) GetCodeWaf(info *form.Base) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.waf/domainWafGetcode", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// UpdateSessionWaf 设置WAF会话固定攻击开关
func (p *DeXunBody) UpdateSessionWaf(info *form.WafStatus) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.waf/domainWafEditsession", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// GetSessionWaf 获取WAF会话固定
func (p *DeXunBody) GetSessionWaf(info *form.Base) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.waf/domainWafGetsession", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// UpdateShellShockWaf 设置WAF的ShellShock后门开关
func (p *DeXunBody) UpdateShellShockWaf(info *form.WafStatus) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.waf/domainWafEditshellshock", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// GetShellShockWaf 获取WAF的ShellShock后门
func (p *DeXunBody) GetShellShockWaf(info *form.Base) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.waf/domainWafGetshellshock", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// UpdateScriptWaf 设置WAF程序检测开关
func (p *DeXunBody) UpdateScriptWaf(info *form.WafStatus) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.waf/domainWafEditscript", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// GetScriptWaf 获取WAF程序检测
func (p *DeXunBody) GetScriptWaf(info *form.Base) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.waf/domainWafGetshellshock", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// UpdateMetaWaf 设置WAF元数据/错误泄露开关
func (p *DeXunBody) UpdateMetaWaf(info *form.WafStatus) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.waf/domainWafEditmetadata", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// GetMetaWaf 获取WAF元数据/错误泄露
func (p *DeXunBody) GetMetaWaf(info *form.Base) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.waf/domainWafGetmetadata", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// UpdateSqlWaf 设置WAF防注入开关
func (p *DeXunBody) UpdateSqlWaf(info *form.WafStatus) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.waf/domainWafEditsqli", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// GetSqlWaf 获取WAF防注入
func (p *DeXunBody) GetSqlWaf(info *form.Base) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.waf/domainWafGetsqli", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// UpdateProxyWaf 设置WAF高级过滤
func (p *DeXunBody) UpdateProxyWaf(info *form.WafProxyStatus) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.waf/domainWafEditfiltering", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// UpdateXssWaf 设置WAF高级过滤
func (p *DeXunBody) UpdateXssWaf(info *form.WafXssStatus) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.waf/domainWafEditfiltering", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// GetProWaf 获取WAF高级过滤
func (p *DeXunBody) GetProWaf(info *form.Base) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.waf/domainWafGetfiltering", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// GetRC 获取非法内容风控信息
func (p *DeXunBody) GetRC(info *form.Base) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.waf/domainFirewallGetrisk", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// GetCC 获取全局CC配置
func (p *DeXunBody) GetCC(info *form.Base) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.firewall/domainOverallCc", nil, info, &result, "application/json"); err != nil {
		return err
	} else {

		return nil
	}
}

// GetAtkLog 获取攻击日志
func (p *DeXunBody) GetAtkLog(info *form.Log) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.log/attackLogAll", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		p.List = result.Data.List
		return nil
	}
}

// GetAccLog 获取访问日志
func (p *DeXunBody) GetAccLog(info *form.Log) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.log/accessLogAll", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		p.List = result.Data.List
		return nil
	}
}

// GetContentLog 内容防护日志
func (p *DeXunBody) GetContentLog(info *form.OtherLog) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.log/contentblocklog", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// GetFlowLog 获取流量统计日志
func (p *DeXunBody) GetFlowLog(info *form.OtherLog) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.log/flowtrafficlist", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		p.List = result.Data.List
		return nil
	}
}

// GetQueryLog 请求统计列表
func (p *DeXunBody) GetQueryLog(info *form.QueryLog) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.log/counttrafficlist", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		p.List = result.Data.List
		return nil
	}
}

// GetAtkKind 拦截类型统计
func (p *DeXunBody) GetAtkKind(info *form.AtkKind) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.log/attackcountvolleyrule", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		p.AppCC = result.Data.AppCC
		p.CC = result.Data.CC
		p.IPBlack = result.Data.CC
		p.Referer = result.Data.Referer
		p.UrLBlack = result.Data.UrLBlack
		p.WebProtect = result.Data.WebProtect
		p.Other = result.Data.Other
		p.AreaAcc = result.Data.AreaAcc
		p.SafeAcc = result.Data.SafeAcc
		p.PreAcc = result.Data.PreAcc
		return nil
	}
}

// GetAtkCount 攻击记录统计
func (p *DeXunBody) GetAtkCount(info *form.AtkKind) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.log/attackcountindex", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		p.List = result.Data.List
		return nil
	}
}

// GetAreaCount 地区统计
func (p *DeXunBody) GetAreaCount(info *form.AtkKind) error {
	var result MultiResponse

	if err := p.SendRequest("POST", "cdncp.log/attackcountarea", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		p.MultiData = result.Data
		return nil
	}
}

// GetDomainAccRanking 域名访问量排行
func (p *DeXunBody) GetDomainAccRanking(info *form.AtkKind) error {
	var result MultiResponse

	if err := p.SendRequest("POST", "cdncp.log/domainAccess", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		p.MultiData = result.Data
		return nil
	}
}

// GetAtkDomain 攻击接口
func (p *DeXunBody) GetAtkDomain(info *form.AtkKind) error {
	var result MultiResponse

	if err := p.SendRequest("POST", "cdncp.log/attackDomain", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		p.MultiData = result.Data
		return nil
	}
}

// GetHttpPack http抓包数据
func (p *DeXunBody) GetHttpPack(info *form.AtkKind) error {
	var result MultiResponse

	if err := p.SendRequest("POST", "cdncp.log/counthttpdataline", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		p.MultiData = result.Data
		return nil
	}
}

// GetFlowLineChart 流量统计折线图
func (p *DeXunBody) GetFlowLineChart(info *form.Info) error {
	var result MultiResponse

	if err := p.SendRequest("POST", "cdncp.log/counttrafficline", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		p.MultiData = result.Data
		return nil
	}
}

// GetBWList 黑白名单统计
func (p *DeXunBody) GetBWList(info *form.Info) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.log/countbwlist", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		p.List = result.Data.List
		return nil
	}
}

// GetTotalFlow 流量统计
func (p *DeXunBody) GetTotalFlow(info *form.Info) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.log/counttrafficsum", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		p.RequestBandwidthPeak = result.Data.RequestBandwidthPeak
		p.Requests = result.Data.Requests
		p.ResponseBandwidthPeak = result.Data.ResponseBandwidthPeak
		p.TotalRequestFlows = result.Data.TotalRequestFlows
		p.TotalResponseFlows = result.Data.TotalResponseFlows
		p.UnidentifiedAttack = result.Data.UnidentifiedAttack
		return nil
	}
}

// GetIPRanking CDN IP访问排名
func (p *DeXunBody) GetIPRanking(info *form.IpRanking) error {
	var result MultiResponse

	if err := p.SendRequest("POST", "cdncp.log/accessranking", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		p.MultiData = result.Data
		return nil
	}
}

// GetDDoSChart1 DDoS告警图表1
func (p *DeXunBody) GetDDoSChart1(info *form.SCDNDDoS) error {
	var result Response

	if err := p.SendRequest("POST", "cdn.gaojing/chart1", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// GetDDoSChart2 DDoS告警图表2
func (p *DeXunBody) GetDDoSChart2(info *form.SCDNDDoS) error {
	var result Response

	if err := p.SendRequest("POST", "cdn.gaojing/chart2", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// GetDDoSChart3 DDoS告警图表3
func (p *DeXunBody) GetDDoSChart3(info *form.SCDNDDoS) error {
	var result Response

	if err := p.SendRequest("POST", "cdn.gaojing/chart3", nil, info, &result, "application/json"); err != nil {
		return err
	} else {
		return nil
	}
}

// GetDDoSComboList 获取DDoS高防IP套餐列表
func (p *DeXunBody) GetDDoSComboList() []Data {
	var result MultiResponse
	var end []Data

	params := url.Values{}
	if err := p.SendRequest("POST", "ddos.tc/lists", params, nil, &result, "application/x-www-form-urlencoded"); err != nil {
		return nil
	}
	end = result.Data
	return end
}

// CreateDDoSOrder DDoS下单接口
func (p *DeXunBody) CreateDDoSOrder(info *form.OrderConsume) string {
	var result Response

	params := form.OrderConsume{UUID: info.UUID, Months: info.Months}
	if err := p.SendRequest("POST", "ddos.dd/platformCreate", nil, params, &result, "application/json"); err != nil {
		return ""
	}
	p.OrderUUID = result.Data.UUID
	p.AddKey = result.Data.AddKey

	return p.OrderUUID
}

// GetDDoSList 获取DDoS订单列表
func (p *DeXunBody) GetDDoSList(keywords, listrows, ks_money string, page int64) []Datum {
	var result Response

	params := form.SCDN{
		Keywords: keywords,
		Page:     page,
		ListRows: listrows,
		Order: form.Order{
			KsMoney: ks_money,
		},
	}
	if err := p.SendRequest("POST", "ddos.dd/lists", nil, params, &result, "application/json"); err != nil {
		return nil
	}

	p.Total = result.Data.Total
	p.PerPage = result.Data.PerPage
	p.CurrentPage = result.Data.CurrentPage
	p.LastPage = result.Data.LastPage
	p.Data = result.Data.Data

	return p.Data
}

// LoginDDoSPanel DDoS登录控制面板，调用控制面板的相关接口前需要先调用这个接口
func (p *DeXunBody) LoginDDoSPanel(info *form.LoginPanel) error {
	var result Response
	if err := p.SendRequest("POST", "ddos.dd/controlPanel", nil, info, &result, "application/json"); err != nil {
		return err
	}
	return nil
}

// GetDDoSCCStatus 获取DDoS的CC防御总揽信息
func (p *DeXunBody) GetDDoSCCStatus(info *form.DDoSCC) error {
	var result Response
	if err := p.SendRequest("POST", "ddos.port/ccDefend", nil, info, &result, "application/json"); err != nil {
		return err
	}
	p.DDInfo = result.Data.DDInfo

	return nil
}

// UpdateDDoSCCTrigger 设置DDoSCC安全防护触发规则
func (p *DeXunBody) UpdateDDoSCCTrigger(info *form.DDoSCCTrigger) error {
	var result Response
	if err := p.SendRequest("POST", "ddos.port/ccSafeDefendRule", nil, info, &result, "application/json"); err != nil {
		return err
	}
	return nil
}

// UpdateDDoSCC 设置DDoSCC安全防护规则
func (p *DeXunBody) UpdateDDoSCC(info *form.DDoSCCInfo) error {
	var result Response

	if err := p.SendRequest("POST", "ddos.port/ccSafeDefend", nil, info, &result, "application/json"); err != nil {
		return err
	}
	return nil
}

// GetExclusiveDomainInfo DDoS获取专属域名信息
func (p *DeXunBody) GetExclusiveDomainInfo(info *form.ExclusiveDomain) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.tcp.domain/info", nil, info, &result, "application/json"); err != nil {
		return err
	}
	p.List = result.Data.List

	return nil
}

// UpdateExclusiveDomainDevice DDoS专属域名的设备更新
func (p *DeXunBody) UpdateExclusiveDomainDevice(info *form.ExclusiveDomainDevice) error {
	var result Response

	if err := p.SendRequest("POST", "cdncp.tcp.domain/tcpEdit", nil, info, &result, "application/json"); err != nil {
		return err
	}

	return nil
}

// GetSSLList SSL证书产品列表接口
func (p *DeXunBody) GetSSLList() []List {
	var result Response

	params := url.Values{}
	if err := p.SendRequest("POST", "ssl.dd/productsslproduct", params, nil, &result, "application/x-www-form-urlencoded"); err != nil {
		return nil
	}
	p.MFCount = result.Data.MFCount
	p.List = result.Data.List

	return p.List
}

// GetSSLCSRandKey 在线生成SSL的CSR和KEY
func (p *DeXunBody) GetSSLCSRandKey(info *form.SSLCsr) error {
	var result Response

	if err := p.SendRequest("POST", "ssl.dd/productsslbuy", nil, info, &result, "application/json"); err != nil {
		return err
	}
	p.Loadcsr = result.Data.Loadcsr
	p.Loadkey = result.Data.Loadkey
	p.Loaddomain = result.Data.Loaddomain

	return nil
}

// CreateSSLOrder 下单SSL
func (p *DeXunBody) CreateSSLOrder(info *form.SSLOrderInfo) string {
	var result Response

	if err := p.SendRequest("POST", "ssl.dd/platformSsl", nil, info, &result, "application/json"); err != nil {
		return ""
	}
	p.OrderUUID = result.Data.UUID
	p.AddKey = result.Data.AddKey

	return p.OrderUUID
}

// GetSSLOrderInfo 获取SSL证书订单详情
func (p *DeXunBody) GetSSLOrderInfo(info *form.SimpleInfo) error {
	var result Response

	if err := p.SendRequest("POST", "ssl.dd/productsslxq", nil, info, &result, "application/json"); err != nil {
		return err
	}
	p.List = result.Data.List
	p.AdminList = result.Data.AdminList
	return nil
}

// GetSSLOrderStatus 获取SSL订单状态
func (p *DeXunBody) GetSSLOrderStatus(info *form.SimpleInfo) error {
	var result Response

	if err := p.SendRequest("POST", "ssl.dd/productsslstatu", nil, info, &result, "application/json"); err != nil {
		return err
	}
	p.SingleData = result.Data

	return nil
}

// GetSSLDownload 下载SSL证书
func (p *DeXunBody) GetSSLDownload(info *form.SSLDownload) error {
	var result Response

	if err := p.SendRequest("POST", "ssl.dd/productsslstatu", nil, info, &result, "application/json"); err != nil {
		return err
	}
	p.SingleData = result.Data

	return nil
}

// GetSSLDomainConfirm 更新/提交SSL证书验证方式
func (p *DeXunBody) GetSSLDomainConfirm(info *form.SSLDomain) error {
	var result Response
	if err := p.SendRequest("POST", "Ssl/productsslmethod", nil, info, &result, "application/json"); err != nil {
		return err
	}

	return nil
}

// GetSSLOrderList 获取SSL证书订单列表
func (p *DeXunBody) GetSSLOrderList(info *form.SSLListInfo) error {
	var result Response
	if err := p.SendRequest("POST", "ssl.dd/productssl", nil, info, &result, "application/json"); err != nil {
		return err
	}
	p.List = result.Data.List

	return nil
}

// SSLOrderRenewal SSL续费
func (p *DeXunBody) SSLOrderRenewal(info *form.SSLRenewalInfo) error {
	var result Response
	if err := p.SendRequest("POST", "ssl.dd/platformSslXuFei", nil, info, &result, "application/json"); err != nil {
		return err
	}

	return nil
}

// GetSSLDns SSL查询域名解析记录或文件验证
func (p *DeXunBody) GetSSLDns(info *form.SimpleInfo) error {
	var result Response
	if err := p.SendRequest("POST", "ssl.dd/getSslDns", nil, info, &result, "application/json"); err != nil {
		return err
	}
	p.SingleData = result.Data

	return nil
}

// QuashSSLOrder 撤销SSL订单
func (p *DeXunBody) QuashSSLOrder(info *form.SimpleInfo) error {
	var result Response
	if err := p.SendRequest("POST", "ssl.dd/ddDel", nil, info, &result, "application/json"); err != nil {
		return err
	}

	return nil
}

// DeleteSSLQuash 已撤销的SSL订单记录删除
func (p *DeXunBody) DeleteSSLQuash(info *form.SimpleInfo) error {
	var result Response
	if err := p.SendRequest("POST", "ssl.dd/ssldel", nil, info, &result, "application/json"); err != nil {
		return err
	}
	return nil
}

// GetCloudEyeComboList 获取云眼套餐列表
func (p *DeXunBody) GetCloudEyeComboList() []Datum {
	var result CEResponse
	if err := p.SendRequest("POST", "ssl.dd/productssl", nil, nil, &result, "application/json"); err != nil {
		return nil
	}
	p.Data = result.Data

	return p.Data
}

// CreateCloudEyeOrder 云眼下单接口
func (p *DeXunBody) CreateCloudEyeOrder(info *form.CloudEyeOrderInfo) string {
	var result Response

	if err := p.SendRequest("POST", "yunjiance.dd/yunjianceddrecordadd", nil, info, &result, "application/json"); err != nil {
		return ""
	}
	p.OrderUUID = result.Data.UUID
	p.AddKey = result.Data.AddKey

	return p.OrderUUID
}
