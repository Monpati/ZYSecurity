package Dexun

import "Dexun/utils"

const (
	Api_UserName  = "温州云锋网络科技有限公司"
	Api_Password  = "yunfeng@025"
	DeXun_Address = "https://apiagent.dexunyun.com/agent"
)

type Response struct {
	Data    Data   `json:"data"`
	Message string `json:"message"`
	Status  int64  `json:"status"`
	Time    int64  `json:"time"`
}

type MultiResponse struct {
	Data    []Data `json:"data"`
	Message string `json:"message"`
	Status  int64  `json:"status"`
	Time    int64  `json:"time"`
}

type CEResponse struct {
	Data    []Datum `json:"data"`
	Message string  `json:"message"`
	Status  int64   `json:"status"`
	Time    int64   `json:"time"`
}

type Data struct {
	Token string `json:"token"`

	DDInfo DDInfo `json:"dd_info"`

	FullName     string `json:"full_name"`
	SmrzBirthday string `json:"smrz_birthday"`
	SmrzCity     string `json:"smrz_city"`
	SmrzGnum     string `json:"smrz_gnum"`
	SmrzSex      string `json:"smrz_sex"`

	SmrzCname string `json:"smrz_cname"`
	SmrzCnum  string `json:"smrz_cnum"`

	CurrentPage int64   `json:"current_page"`
	Data        []Datum `json:"data"`
	LastPage    int64   `json:"last_page"`
	PerPage     int64   `json:"per_page"`
	Total       int64   `json:"total"`
	// 域名列表
	List      []List      `json:"list"`
	AdminList []AdminList `json:"admin_list"`
	// 总数
	MFCount int64 `json:"mf_count"`

	// 全量访问控制状态，2关闭/1开启
	CompleteState int64 `json:"complete_state,omitempty"`
	// DDoS防护峰值
	DdosHh string `json:"ddos_hh"`
	// 域名数限制，0是无限制
	DomainNum int64 `json:"domain_num"`
	// 防火墙控制状态，0关闭/1开启
	FirewallState int64 `json:"firewall_state"`
	// 套餐类型，| 套餐类型 | 描述              |
	// | -------- | ----------------- |
	// | 1        | 独立              |
	// | 2        | 共享              |
	// | 3        | 共享(仅限 443 80) |
	GroupType int64 `json:"group_type"`
	// 套餐原价
	KsMoney int64 `json:"ks_money"`
	// 套餐流量，单位: Gb
	ProFlow int64 `json:"pro_flow"`
	// 套餐说明
	ProNote ProNote `json:"pro_note"`
	// 套餐名称
	TcName string `json:"tc_name"`
	// 套餐uuid、包uuid
	UUID string `json:"uuid"`
	// 高级WAF配置控制状态，0关闭/1开启
	WafState int64 `json:"waf_state"`
	// 代理折扣价
	ZkMoney int64 `json:"zk_money"`
	PortNum int64 `json:"port_num"`
	YwdkNum int64 `gorm:"column:ywdk_num" json:"ywdk_num"`

	// 包流量，单位：Gb
	ProtectLV int64 `json:"protect_lv"`
	// 包名
	ProtectNote string `json:"protect_note"`
	// 包价格
	ProtectPrice int64 `json:"protect_price"`
	//包种类
	Type string `json:"type"`

	// 回调类型，| 产品类型 | 产品名称 |
	// | -------- | -------- |
	// | 1        | 下单     |
	// | 2        | 续费     |
	// | 3        | 升级     |
	// | 0        | 未知     |
	CallbackType int64 `json:"callback_type"`
	// 订单数据，回调返回需要的订单数据
	OrderData OrderData `json:"order_data"`
	// 订单状态，| 产品类型 | 产品名称 |
	// | -------- | -------- |
	// | 0        | 处理失败 |
	// | 1        | 处理成功 |
	// | 2        | 待处理   |
	OrderStatus int64 `json:"order_status"`
	// 订单uuid
	OrderUUID string `json:"order_uuid"`
	// 产品类型，| 产品类型 | 产品名称      |
	// | ---- | ------------- |
	// | 9    | 安全加速 SCDN |
	// | 10   | DDoS 高防 IP  |
	// | 11   | SSL 证书      |
	ProductType int64 `json:"product_type"`

	// 回调key
	AddKey string `json:"add_key"`

	// 节点类型说明
	NodeNote string `json:"node_note"`
	// 节点类型，1 独立节点
	NodeStatus int64 `json:"node_status"`

	AccessActive string `json:"access_active"`
	// 缓存状态
	CacheActive int64 `json:"cache_active"`
	// 缓存配置
	CacheConfig []string `json:"cache_config"`
	// 7层配置列表
	ConfigList []ConfigList `json:"config_list"`
	// 创建时间
	Createtime string `json:"createtime"`
	// 域名
	Domain string `json:"domain"`
	// 域名公共配置
	DomainConfig DomainConfig `json:"domain_config"`
	DomainRecord interface{}  `json:"domain_record"`
	// 域名状态
	DomainStatus int64 `json:"domain_status"`
	// 域名uuid
	DomainUUID string `json:"domain_uuid"`
	// 4层配置列表，> 仅DDoS高防IP的TCP专属域名有效
	FourLayersConfig []FourLayersConfig `json:"four_layers_config"`
	Grouping         string             `json:"grouping"`
	// 是否备案
	IsFilings string `json:"is_filings"`
	// 域名cname后缀，> cname = `domain_uuid.primary_domain`
	PrimaryDomain string `json:"primary_domain"`
	// 最近更新时间
	Updatetime string `json:"updatetime"`
	// 使用流量
	UseFlow int64 `json:"use_flow"`
	// 白名单数量
	WhiteNum int64 `json:"white_num"`

	// 数字证书
	CERT string `json:"cert"`
	// 证书名称
	CERTName string `json:"cert_name"`
	// 无效参数
	Desc string `json:"desc"`
	// HSTS，1.开启，0.关闭
	Hsts int64 `json:"hsts"`
	// id
	ID int64 `json:"id"`
	// 加密秘钥
	Key string `json:"key"`
	// 强制SSL，1.开启，0.关闭
	SSLAlways int64 `json:"ssl_always"`
	// 状态，1.开启，0.关闭
	Status int64 `json:"status"`

	// 开关
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

	// 数量，单位：个
	Count int64 `json:"count,omitempty"`
	// 地区
	Source string `json:"source,omitempty"`

	// 统计
	TotalCount int64 `json:"total_count"`
	// 时间戳
	Time int64 `json:"time"`

	// 请求总流量，单位：G
	RequestSize int64 `json:"request_size"`
	// 响应总流量，单位：G
	ResponseSize int64 `json:"response_size"`

	// 宽带峰值，单位：Mbps
	RequestBandwidthPeak int64 `json:"request_bandwidth_peak"`
	// 请求数，- 单位：次
	// - DDoS高防IP不准确
	Requests              int64 `json:"requests"`
	ResponseBandwidthPeak int64 `json:"response_bandwidth_peak"`
	// 请求总流量，单位：GB
	TotalRequestFlows  int64 `json:"total_request_flows"`
	TotalResponseFlows int64 `json:"total_response_flows"`
	// 缓存命中率，单位：百分比
	UnidentifiedAttack int64 `json:"unidentified_attack"`

	// ip
	ClientIP string `json:"client_ip,omitempty"`
	// 统计
	CountSum int64 `json:"count_sum,omitempty"`

	Loadcsr    string `json:"loadcsr"`
	Loaddomain string `json:"loaddomain"`
	Loadkey    string `json:"loadkey"`

	// 申请人邮箱
	AdminEmail string `json:"admin_email"`
	// 申请人职务
	AdminJob string `json:"admin_job"`
	// 申请人手机
	AdminMobile string `json:"admin_mobile"`
	// 申请人姓名
	AdminNeme string `json:"admin_neme"`
	// 有效时间
	DateDiff string `json:"date_diff"`
	// 有效期结束
	EdTime string `json:"ed_time"`
	// 证书类型 图片
	Img string `json:"img"`
	// 有效期开始
	KsTime string `json:"ks_time"`
	// 订单编号
	OrderID string `json:"order_id"`
	// 证书状态
	OrderName string `json:"order_name"`
	// 购买日期
	OrderTime string `json:"order_time"`
	// 地址
	OrgAddress string `json:"org_address"`
	// 部门
	OrgClass string `json:"org_class"`
	// 邮编
	OrgCode string `json:"org_code"`
	// 公司
	OrgCompany string `json:"org_company"`
	// 电话
	OrgMobile string `json:"org_mobile"`
	// 证书安装服务
	SetupServer string `json:"setup_server"`
	// 证书名称
	SSLName string `json:"ssl_name"`
	// 证书类型 名称
	SSLTypeName string `json:"ssl_type_name"`
	// 技术员邮箱
	TechEmail string `json:"tech_email"`
	// 技术员职务
	TechJob string `json:"tech_job"`
	// 技术员手机
	TechMobile string `json:"tech_mobile"`
	// 技术员姓名
	TechNeme string `json:"tech_neme"`
	// CA证书编号
	VendorID string `json:"vendor_id"`
	// 域名
	ZDomain string `json:"z_domain"`

	URL string `json:"url"`

	// dns主机记录
	DNSHost string `json:"dns_host"`
	// dns记录类型[cname]
	DNSType string `json:"dns_type"`
	// dns记录值
	DNSValue string `json:"dns_value"`
	// 验证文件名称
	FileName string `json:"file_name"`
	// 验证文件内容
	FileValue string `json:"file_value"`
}

type SSL struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Img   string `json:"img"`
}

type FourLayersConfig struct {
	// 健康检查开关，字符串true开启/字符串false关闭
	HealthCheckActive string `json:"health_check_active"`
	// 间隔时间
	Interval string `json:"interval"`
	// 负载均衡，0轮询/1权重
	LoadBalancing string `json:"load_balancing"`
	// 频率限制-并发连接数，0不限制
	MaxConnection string `json:"max_connection"`
	// 频率限制-新建连接数，0不限制
	NewConnection string `json:"new_connection"`
	// 监听端口
	Port string `json:"port"`
	// 协议类型，tcp/udp
	Protocol string `json:"protocol"`
	// Proxy Protocol，| 类型 | 说明          |
	// | ---- | ------------- |
	// | 0    | 关闭          |
	// | 1    | US-ASCII 编码 |
	// | 2    | 二进制编码    |
	ProxyProtocol string `json:"proxy_protocol"`
	// 频率限制-发包频率，- 0不限制
	// - 单位：次/秒
	RateLimit string `json:"rate_limit"`
	// 恢复次数，单位：秒
	RecoverTimes string `json:"recover_times"`
	// 源地址列表
	SourceAddress []SourceAddress `json:"source_address"`
	// 超时时间
	Timeout string `json:"timeout"`
}

type DDInfo struct {
	// 并发数，默认 10000
	GfCcIn int64 `json:"gf_cc_in"`
	// CC安全防护模式，7开头普通防御 8开头高级防御 9开头验证防御
	GfCcInfo int64 `json:"gf_cc_info"`
	// 链接数，默认 1000
	GfCcNewtcp int64 `json:"gf_cc_newtcp"`
	// SYN包数，默认 2000
	GfCcSyn int64 `json:"gf_cc_syn"`
	// CC安全防护状态，0AI防御 1手动防御
	GfYnCckg int64 `json:"gf_yn_cckg"`
	// 节点IP
	ServerIP string `json:"server_ip"`
	// 订单uuid
	UUID string `json:"uuid"`
}

type AdminList struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type List struct {
	SSLOrderInfo interface{} `json:"ssl_order_info"`
	// 是否开启
	Active int64 `json:"active"`
	// 缓存uuid
	CacheUUID string `json:"cache_uuid"`
	// 类型：后缀匹配，类型三选一，可为空
	Cacheextensions string `json:"cacheextensions"`
	// 缓存匹配模式，选项 path，ext，reg
	Cachemode string `json:"cachemode"`
	// 类型：路径匹配类型：路径匹配，类型三选一，可为空
	Cachepath string `json:"cachepath"`
	// 类型：正则表达式匹配，类型三选一，可为空
	Cachereg string `json:"cachereg"`
	// 创建时间
	//Createtime                 string `json:"createtime"`
	// 域名uuid
	//DomainUUID                 string `json:"domain_uuid"`
	// id
	ID int64 `json:"id"`
	// 缓存过期时间，秒
	Timeout int64 `json:"timeout"`
	// 跟新时间
	//Updatetime                 string `json:"updatetime"`
	// url匹配模式，选项 full，onlypath
	Urlmode string `json:"urlmode"`
	// uuid
	UUID string `json:"uuid"`
	// 权重，1-100
	Weight int64 `json:"weight"`

	// 缓存预热状态，1开启/0关闭
	CacheActive int64 `json:"cache_active"`
	// 缓存预热配置，url列表
	CacheConfig []string `json:"cache_config"`
	// 解析状态，1已接入/0未接入
	CnameStatus int64 `json:"cname_status"`
	// 7层配置信息，详情见 域名信息
	ConfigList []ConfigList `json:"config_list"`
	// 创建时间
	Createtime string `json:"createtime"`
	// 域名
	Domain       string      `json:"domain"`
	DomainRecord interface{} `json:"domain_record"`
	// 状态，1启用/0禁用
	DomainStatus int64 `json:"domain_status"`
	// 域名uuid
	DomainUUID string `json:"domain_uuid"`
	// 备案状态，0未备案/1已备案/3查询中/4查询失败
	IsFilings string `json:"is_filings"`
	// cname后缀
	PrimaryDomain string `json:"primary_domain"`
	// 最近更新时间
	Updatetime string `json:"updatetime"`
	// 使用流量
	UseFlow int64 `json:"use_flow"`

	// 拦截信息
	Attackinfo string `json:"attackinfo"`
	// 拦截规则类型
	Attacktype string `json:"attacktype"`
	// 访问IP
	Clientip string `json:"clientip"`
	// 访问端口
	Clientport int64 `json:"clientport"`
	// 访问地区
	Clientregion string `json:"clientregion"`
	// 统计
	Count int64 `json:"count"`
	// 域名uuid
	Domainid string `json:"domainid"`
	// 请求方法
	Httpmethod string `json:"httpmethod"`
	// 实例id
	Instanceid int64 `json:"instanceid"`
	// 使用的节点ip
	Localip string `json:"localip"`
	// 方法
	Method string `json:"method"`
	// 节点uuid
	Nodeid string `json:"nodeid"`
	// 防护措施
	Protecttype string `json:"protecttype"`
	// 请求信息
	Requestinfo string `json:"requestinfo"`
	// 请求地址
	Targeturl string `json:"targeturl"`
	// 访问结束时间
	Timerangeend string `json:"timerangeend"`
	// 访问开始时间
	Timerangestart string `json:"timerangestart"`

	// 缓存命中
	Cachehit string `json:"cachehit"`
	// 客户端ip
	Createdat interface{} `json:"createdat"`
	Form      string      `json:"form"`
	// 使用节点地址
	Localaddr string `json:"localaddr"`
	// 使用节点端口
	Localport int64 `json:"localport"`
	// 包大小
	Packagesize int64 `json:"packagesize"`
	// 远程地址
	Remoteaddr string `json:"remoteaddr"`
	// 请求头
	Requestheaders Requestheaders `json:"requestheaders"`
	// 响应头
	Responseheaders Responseheaders `json:"responseheaders"`
	// 请求大小
	Responsesize int64 `json:"responsesize"`
	// 响应代码
	Responsestatuscode int64 `json:"responsestatuscode"`
	// 访问地址
	URL    string `json:"url"`
	Wblist string `json:"wblist"`

	// 请求总流量，单位：Gb
	RequestSize int64 `json:"request_size"`
	// 响应总流量，单位：Gb
	ResponseSize int64 `json:"response_size"`

	// 缓存命中，单位：次
	CacheCalls int64 `json:"cache_calls"`
	// 缓存命中率，单位：百分比
	CacheRate int64 `json:"cache_rate"`
	// 请求总数，单位：次
	TotalCalls int64 `json:"total_calls"`

	// 时间戳
	Time int64 `json:"time"`
	// 统计
	TotalCount int64 `json:"total_count"`

	// 类型，0:白名单 1:黑名单
	BWListType string `json:"bw_list_type"`
	// ip
	IP string `json:"ip"`

	// 域名数量[包含域名数]
	DomainNum string `json:"domain_num"`
	// 域名类型，1:单域名 2:多域名 3:通配符(*)
	DomainType string `json:"domain_type"`
	// 额外增加域名每个的费用
	EyMoney int64 `json:"ey_money"`
	// 开设价格
	KsMoney int64 `json:"ks_money"`
	// 市场价
	MarketMoney int64 `json:"market_money"`
	// 备注
	PNote string `json:"p_note"`
	// 品牌id号
	PTypeID int64 `json:"p_type_id"`
	// 品牌名称
	PTypeName string `json:"p_type_name"`
	// 推荐，并排序
	SType int64 `json:"s_type"`
	// 证书编码[购买使用]'
	SSLCode string `json:"ssl_code"`
	// 证书名称
	SSLName string `json:"ssl_name"`
	// 证书类型
	SSLType string `json:"ssl_type"`
	// 申请期限
	Term string `json:"term"`

	// 申请人信息[姓;名;手机号;邮箱;职务]
	AdminInfo string `json:"admin_info,omitempty"`
	// 0已付款(接口未执行)/1COM接口执行中/2COM接口执行失败/3完成
	COMStatus int64 `json:"com_status"`
	// dns主机记录
	DNSHost string `json:"dns_host,omitempty"`
	// dns记录类型[cname]
	DNSType string `json:"dns_type,omitempty"`
	// dns记录值
	DNSValue string `json:"dns_value,omitempty"`
	// 证书附加域名
	DomainList string `json:"domain_list,omitempty"`
	// 证书结束时间
	EdTime string `json:"ed_time"`
	// 验证文件名称
	FileName string `json:"file_name,omitempty"`
	// 验证文件内容
	FileValue string `json:"file_value,omitempty"`
	// 证书开始时间
	KsTime interface{} `json:"ks_time"`
	// 返回订单号
	OrderID string `json:"order_id,omitempty"`
	// 订单状态[0:待验证  1:待签发  2:已签发 3:申请失败  4.已取消]
	OrderStart int64 `json:"order_start,omitempty"`
	// 下单时间
	OrderTime string `json:"order_time,omitempty"`
	// 公司信息[企业名称;部门名称;省份;城市;地址;电话;邮编]
	OrgInfo string `json:"org_info,omitempty"`
	// 域名验证方式[http/dns]
	PMethod string `json:"p_method,omitempty"`
	// 是否有证书安装服务 0 为无 1 为有
	SetupServer int64 `json:"setup_server,omitempty"`
	// 生成或提交的crs代码
	SSLCsr string `json:"ssl_csr,omitempty"`
	// 生成的key
	SSLKey string `json:"ssl_key,omitempty"`
	// 服务器证书代码crt
	SSLPem string `json:"ssl_pem"`
	// 技术人员信息[姓;名;手机号;邮箱;职务]
	TechInfo string `json:"tech_info,omitempty"`
	// vendor_id
	VendorID string `json:"vendor_id,omitempty"`
	// 续费后的新订单号
	XufeiOrderid string `json:"xufei_orderid,omitempty"`
	// 是否验证 1为是 0为否
	YnProve int64 `json:"yn_prove,omitempty"`
	// 是否重新签发 0为否 1为是
	YnReplace int64 `json:"yn_replace,omitempty"`
	// 是否续费[0 未续费  1已续费]
	YnXufei string `json:"yn_xufei,omitempty"`
	// 证书主域名
	ZDomain     string `json:"z_domain,omitempty"`
	OrderName   string `json:"order_name,omitempty"`
	ZDomainList string `json:"z_domain_list,omitempty"`
	DateDiff    int64  `json:"date_doff,omitempty"`

	Name  string `json:"name"`
	Value string `json:"value"`
	Img   string `json:"img"`
}

type Requestheaders struct {
	Accept                  string  `json:"Accept"`
	AcceptEncoding          string  `json:"Accept-Encoding"`
	AcceptLanguage          string  `json:"Accept-Language"`
	Authorization           string  `json:"Authorization"`
	CacheControl            string  `json:"Cache-Control"`
	Connection              string  `json:"Connection"`
	Pragma                  string  `json:"Pragma"`
	Purpose                 *string `json:"Purpose,omitempty"`
	Referer                 string  `json:"Referer"`
	UpgradeInsecureRequests *string `json:"Upgrade-Insecure-Requests,omitempty"`
	UserAgent               string  `json:"User-Agent"`
	// 真实域名
	XForwardedHost string `json:"X-Forwarded-Host"`
	// 真实端口
	XForwardedPort string `json:"X-Forwarded-Port"`
	// 真实协议
	XForwardedProto string `json:"X-Forwarded-Proto"`
	// 服务商
	XForwardedServer string `json:"X-Forwarded-Server"`
	// 真实IP
	XRealIP string `json:"X-Real-Ip"`
}

// 响应头
type Responseheaders struct {
	AcceptRanges    *string `json:"Accept-Ranges,omitempty"`
	ContentEncoding string  `json:"Content-Encoding"`
	ContentLength   string  `json:"Content-Length"`
	ContentType     string  `json:"Content-Type"`
	Date            string  `json:"Date"`
	Etag            string  `json:"Etag"`
	LastModified    string  `json:"Last-Modified"`
	Server          string  `json:"Server"`
	Vary            string  `json:"Vary"`
	WWWAuthenticate string  `json:"Www-Authenticate"`
}

type ConfigList struct {
	// 负载均衡-类型，0.轮询,1.权重,2.IP哈希,3.URL哈希
	LoadBalancing string `json:"load_balancing"`
	// 并发超限处置-301重定向跳转地址
	OverloadRedirectURL string `json:"overload_redirect_url"`
	// 并发超限处置-自定义阻断状态码
	OverloadStatusCode *string `json:"overload_status_code,omitempty"`
	OverloadSelect     *string `json:"overload_select,omitempty"`
	// 并发超限处置-类型，| 类型 | 并发超限处置     |
	// | ---- | ---------------- |
	// | 1    | 400 阻断         |
	// | 2    | 301 跳转         |
	// | 3    | 自定义阻断状态码 |
	OverloadType string `json:"overload_type"`
	// 监听端口
	Port string `json:"port"`
	// 协议类型
	Protocol string `json:"protocol"`
	// 取源协议-301重定向状态，true启用/false禁用
	Redirect string `json:"redirect"`
	// 服务器响应头，默认空
	Server string `json:"server"`
	// 取源协议-数据，redirect = false 生效
	SourceAddresses []SourceAddress `json:"source_addresses"`
	// 取源协议-301重定向跳转地址，redirect = true 生效
	URIForward *string `json:"uri_forward,omitempty"`
}

type SourceAddress struct {
	// 地址
	Address string `json:"address"`
	// 并发，0不限制
	Concurrent       string  `json:"concurrent"`
	ConcurrentSelect *string `json:"concurrent_select,omitempty"`
	// 端口
	Port string `json:"port"`
	// 协议，http/https
	Protocol string `json:"protocol"`
	// SNI，空字符串或指定域名
	Sni       string  `json:"sni"`
	SniSelect *string `json:"sni_select,omitempty"`
	// 权重，1-100
	Weight string `json:"weight"`
}

type Config struct {
	// 地址列表
	Regions utils.JSON `json:"regions"`
}

// 域名公共配置
type DomainConfig struct {
	// 文件缓存大小
	CacheFileSizeLimit int64 `json:"cache_file_size_limit"`
	// 统计缓存大小
	CacheTotalSizeLimit int64 `json:"cache_total_size_limit"`
}

// OrderData 订单数据，回调返回需要的订单数据
type OrderData struct {
	// 没用到
	CompleteState *int64 `json:"complete_state,omitempty"`
	// 产品实例名称
	ProductSitename *string `json:"product_sitename,omitempty"`
	// ip
	ServerIP *string `json:"server_ip,omitempty"`
	// 订单uuid
	UUID string `json:"uuid,omitempty"`
}

type Datum struct {
	ActuaFlow       int64  `json:"actua_flow,omitempty"`
	AIModeSum       int64  `json:"ai_mode_sum,omitempty"`
	AIOpen          int64  `json:"ai_open,omitempty"`
	DateDiff        int64  `json:"date_diff,omitempty"`
	DdosHh          string `json:"ddos_hh,omitempty"`
	DomainNum       int64  `json:"domain_num,omitempty"`
	EndTime         string `json:"end_time,omitempty"`
	KsMoney         int64  `json:"ks_money,omitempty"`
	KsStart         int64  `json:"ks_start,omitempty"`
	ProFlow         int64  `json:"pro_flow,omitempty"`
	ProductSitename string `json:"product_sitename,omitempty"`
	RechargeFlow    int64  `json:"recharge_flow,omitempty"`
	ServerIP        string `json:"server_ip,omitempty"`
	SiteStart       int64  `json:"site_start,omitempty"`
	StatTime        string `json:"stat_time,omitempty"`
	TcName          string `json:"tc_name,omitempty"`
	TotalFlow       int64  `json:"total_flow,omitempty"`
	UUserID         int64  `json:"u_user_id,omitempty"`
	UUID            string `json:"uuid,omitempty"`

	// 数量，单位：个
	Count *int64 `json:"count,omitempty"`
	// 地区
	Source *string `json:"source,omitempty"`

	// 端口数量
	PortNum *int64 `json:"port_num,omitempty"`
	// 域名数增量包
	RechargeDomain *int64 `json:"recharge_domain,omitempty"`
	// 端口数增量包
	RechargePort *int64 `json:"recharge_port,omitempty"`

	Content     string     `json:"content"`
	MonitorType utils.JSON `json:"monitor_type"`
	OrderMoney  int64      `json:"order_money"`
	StartCount  int64      `json:"start_count"`
	TaskType    utils.JSON `json:"task_type"`
	ZkMoney     int64      `json:"zk_money"`
}

type ProNote struct {
	Ccfy  string `json:"ccfy"`
	Gjcs  string `json:"gjcs"`
	Gjwaf string `json:"gjwaf"`
	Yms   string `json:"yms"`
	Ywll  string `json:"ywll"`
	XL    string `json:"xl"`
	Zfdks string `json:"zfdks"`
	Fhyms string `json:"fhyms"`
	Ywdk  string `json:"ywdk"`
}
