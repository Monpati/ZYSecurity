package form

import "Dexun/utils"

type Base struct {
	// 订单uuid
	DDUUID string `json:"dd_uuid"`
	// 域名uuid
	DomainUUID string `json:"domain_uuid"`
	// 订单类型，9.CDN，10.DDoS
	ProType string `json:"pro_type"`
}

type Cert struct {
	Gname string `json:"gname"`
	Gnum  string `json:"gnum"`

	Cname       string `json:"cname"`
	Cnum        string `json:"cnum"`
	Compamyname string `json:"compamyname"`
	Scity       string `json:"scity"`
}

type ExclusiveDomain struct {
	DDUUID  string `json:"dd_uuid"`
	ProType int64  `json:"pro_type"`
}

type ExclusiveDomainDevice struct {
	// 操作日志数据，不传则没有记录信息
	ActiveLogData ActiveLogData `json:"active_log_data"`
	// 全量配置列表
	Config []Config `json:"config"`
	// 订单uuid
	DDUUID string `json:"dd_uuid"`
	// 产品类型
	ProType int64 `json:"pro_type"`
}

// 操作日志数据，不传则没有记录信息
type ActiveLogData struct {
	// 监听端口
	Port string `json:"port"`
	// 协议类型，TCP/UDP
	Protocol string `json:"protocol"`
	// 源地址列表
	SourceAddresses []string `json:"source_addresses"`
	// 操作类型，edit编辑/create新增
	Type string `json:"type"`
}

type LoginPanel struct {
	// 订单uuid
	DDUUID string `json:"dd_uuid"`
}

type SimpleInfo struct {
	// 订单uuid
	UUID string `json:"uuid"`
}

type SSLListInfo struct {
	// 关键词
	Keywords string `json:"keywords"`
}

type SSLDownload struct {
	// 下载类型
	// apache
	// iis
	// jks
	// nginx
	// other
	// tomcat
	Type string `json:"type"`
	UUID int64  `json:"uuid"`
}

type SSLDomain struct {
	// 域名验证  DNS验证：dns    文件验证：http
	MType string `json:"m_type"`
	// 订单uuid
	UUID int64 `json:"uuid"`
}

type SSLRenewalInfo struct {
	// 域名验证方式 域名验证 DNS验证：dns 文件验证：http
	Provemethod string `json:"provemethod"`
	// 订单id
	UUID int64 `json:"uuid"`
}

type SSLOrder struct {
	// 申请人信息 格式 姓;名;手机号码;电子邮箱;职务名称
	Admininfo string `json:"admininfo"`
	// 附加域名，以 ; 号分割
	DomainList string `json:"domain_list"`
	// 证书CSR
	Loadcsr string `json:"loadcsr"`
	// 主域名
	Loaddomain string `json:"loaddomain"`
	// 证书KEY
	Loadkey string `json:"loadkey"`
	// 企业信息 公司名称;部门名称; 所在省份;所在城市;详细地址;电话号码;邮编地址
	Orginfo string `json:"orginfo"`
	// 证书列表uuid
	PUUID string `json:"p_uuid"`
	// 域名验证 DNS验证：dns 文件验证：http
	Provemethod string `json:"provemethod"`
	// 技术联系人 格式 姓;名;手机号码;电子邮箱;职务名称
	Techinfo string `json:"techinfo"`
}

type DDoSCC struct {
	// 订单uuid
	DDUUID string `json:"dd_uuid"`
}

type DDoSCCTrigger struct {
	// 订单uuid
	DDUUID string `json:"dd_uuid"`
	// 并发数
	GfCcIn int64 `json:"gf_cc_in"`
	// 新建链接
	GfCcNewtcp int64 `json:"gf_cc_newtcp"`
	// SYN包数
	GfCcSyn int64 `json:"gf_cc_syn"`
}

type DDoSCCInfo struct {
	// 订单uuid
	DDUUID string `json:"dd_uuid"`
	// 模式，| 模式类型 | 对应模式       | 备注     |
	// | -------- | -------------- | -------- |
	// | 1        | `7,True,False` | 普通防御 |
	// | 2        | `8,True,False` | 高级防御 |
	// | 3        | `9,True,False` | 验证防御 |
	GfCcInfo int64 `json:"gf_cc_info"`
	// 状态，| 状态 | 防御模式 |
	// | ---- | -------- |
	// | 0    | AI 防御  |
	// | 1    | 手动防御 |
	GfYnCckg int64 `json:"gf_yn_cckg"`
}

type SSLCsr struct {
	// 绑定域名
	Csrdomain string `json:"csrdomain"`
	// 所在城市
	Csrlocality string `json:"csrlocality"`
	// 公司名称
	Csrorg string `json:"csrorg"`
	// 所在省份
	Csrstate string `json:"csrstate"`
	// 部门名称
	Department string `json:"department"`
}

type CloudEyeOrderInfo struct {
	// 额外任务数量
	Count int64 `json:"count"`
	// 月份
	Months int64 `json:"months"`
	// 套餐uuid
	TcUUID string `json:"tc_uuid"`
}

type SSLOrderInfo struct {
	// 申请人信息 格式 姓;名;手机号码;电子邮箱;职务名称
	Admininfo string `json:"admininfo"`
	// 附加域名，以 ; 号分割
	DomainList string `json:"domain_list"`
	// 证书CSR
	Loadcsr string `json:"loadcsr"`
	// 主域名
	Loaddomain string `json:"loaddomain"`
	// 证书KEY
	Loadkey string `json:"loadkey"`
	// 企业信息 公司名称;部门名称; 所在省份;所在城市;详细地址;电话号码;邮编地址
	Orginfo string `json:"orginfo"`
	// 证书列表uuid
	PUUID string `json:"p_uuid"`
	// 域名验证 DNS验证：dns 文件验证：http
	Provemethod string `json:"provemethod"`
	// 技术联系人 格式 姓;名;手机号码;电子邮箱;职务名称
	Techinfo string `json:"techinfo"`
}

type ScdnFilterForm struct {
	Field string `json:"field"`
	Value string `json:"value"`
	PageForm
}

type CallBack struct {
	UUID []string `json:"uuid"`
}

type DDoSMove struct {
	Id     int64 `json:"id"`
	Status int   `json:"status"`
}

type ScdnMove struct {
	Id     int64 `json:"id"`
	Status int   `json:"status"`
}

type SCDN struct {
	Keywords string `json:"keywords"`
	ListRows string `json:"list_rows"`
	Order    Order  `json:"order"`
	Page     int64  `json:"page"`
}

type Order struct {
	KsMoney string `json:"ks_money"`
}

type ScdnCombos struct {
	Id            int64  `json:"id"`
	UUID          int64  `json:"uuid"`
	TcName        string `json:"tc_name"`
	KsMoney       string `json:"ks_money"`
	Yms           string `json:"yms"`
	Ccfy          string `json:"ccfy"`
	Gjwaf         string `json:"gjwaf"`
	Ywll          string `json:"ywll"`
	ProFlow       int64  `json:"pro_flow"`
	DdosHh        string `json:"ddos_hh"`
	DomainNum     int64  `json:"domain_num"`
	CompleteState int64  `json:"complete_state"`
	FirewallState int64  `json:"firewall_state"`
	WafState      int64  `json:"waf_state"`
}

type OrderConsume struct {
	UUID   string `json:"tc_uuid"`
	Months int64  `json:"months"`
}

type InstanceList struct {
	InstanceNameList []string `json:"instance_name_list"`
}

type SafeAcc struct {
	// 订单uuid
	DDUUID string `json:"dd_uuid"`
	// 域名uuid
	DomainUUID string `json:"domain_uuid"`
	// 订单类型，9.CDN，10.DDoS
	ProType string     `json:"pro_type"`
	URLList []SafeList `json:"url_list"`
}

type SafeList struct {
	// 访问密码
	Password string `json:"password"`
	// URI
	URL string `json:"url"`
}

type UrlBW struct {
	// 订单uuid
	DDUUID string `json:"dd_uuid"`
	// 域名uuid
	DomainUUID string `json:"domain_uuid"`
	// 订单类型，9.CDN，10.DDoS
	ProType string `json:"pro_type"`
	// 黑白名单类型，0：白名单，1：黑名单
	Type string `json:"type"`
	// url列表
	URLList URLList `json:"url_list"`
}

type URLList struct {
	// 方法，("contains","包含"), ("ends","匹配末尾"), ("regex","正则")
	Method string `json:"method,omitempty"`
	// 路径
	Path string `json:"path,omitempty"`
}

type WafStatus struct {
	// 是否开启
	Active int64 `json:"active"`
	// 订单uuid
	DDUUID string `json:"dd_uuid"`
	// 域名uuid
	DomainUUID string `json:"domain_uuid"`
	// 订单类型，9.CDN，10.DDoS
	ProType string `json:"pro_type"`
}

type WafProxyStatus struct {
	// 订单uuid
	DDUUID string `json:"dd_uuid"`
	// 域名uuid
	DomainUUID string `json:"domain_uuid"`
	// 订单类型，9.CDN，10.DDoS
	ProType string `json:"pro_type"`
	// 代理
	Proxy int64 `json:"proxy"`
}

type WafXssStatus struct {
	// 订单uuid
	DDUUID string `json:"dd_uuid"`
	// 域名uuid
	DomainUUID string `json:"domain_uuid"`
	// 订单类型，9.CDN，10.DDoS
	ProType string `json:"pro_type"`
	// xss开关
	XSS int64 `json:"xss"`
}

type IpBW struct {
	// 订单uuid
	DDUUID string `json:"dd_uuid"`
	// 域名uuid
	DomainUUID string `json:"domain_uuid"`
	// ip列表，["1.1.1.1","1.1.1.2"]
	IPList []string `json:"ip_list"`
	// 订单类型，9.CDN，10.DDoS
	ProType string `json:"pro_type"`
	// 黑白名单类型，0：白名单，1：黑名单
	Type int64 `json:"type"`
}

type Customized struct {
	// 内容
	Content string `json:"content"`
	// 订单uuid
	DDUUID string `json:"dd_uuid"`
	// 域名uuid
	DomainUUID string `json:"domain_uuid"`
	// 类型，定制页面：写死的， 错误页面：500，501，404等
	PageType string `json:"page_type"`
	// 订单类型，9.CDN，10.DDoS
	ProType string `json:"pro_type"`
	// 定制类型，0：定制页面， 1：错误页面
	Type int64 `json:"type"`
}

type Certification struct {
	// 数字证书
	CERT string `json:"cert"`
	// 证书名称
	CERTName string `json:"cert_name"`
	// 订单uuid
	DDUUID string `json:"dd_uuid"`
	// 域名uuid
	DomainUUID string `json:"domain_uuid"`
	// HSTS，1.开启，0.关闭
	Hsts int64 `json:"hsts"`
	// 加密秘钥
	Key string `json:"key"`
	// 订单类型，9.CDN，10.DDoS
	ProType string `json:"pro_type"`
	// 强制SSL，1.开启，0.关闭
	SSLAlways int64 `json:"ssl_always"`
	// 状态，1.开启，0.关闭
	Status int64 `json:"status"`
}

type Cache struct {
	// 缓存开关，1.开启，0.关闭
	Active int64 `json:"active"`
	// 扩展名匹配，类型三选一，可为空
	Cacheextensions string `json:"cacheextensions"`
	// url匹配模式，path，ext，reg
	Cachemode string `json:"cachemode"`
	// 路径匹配，类型三选一，可为空
	Cachepath string `json:"cachepath"`
	// 正则匹配，类型三选一，可为空
	Cachereg string `json:"cachereg"`
	// 订单uuid
	DDUUID string `json:"dd_uuid"`
	// 订单类型，9.CDN，10.DDoS
	ProType string `json:"pro_type"`
	// 缓存过期时间，分
	Timeout string `json:"timeout"`
	// url匹配模式，full，onlypath
	Urlmode string `json:"urlmode"`
	// 权重，1-100
	Weight string `json:"weight"`

	// 域名uuid
	DomainUUID string `json:"domain_uuid"`
}

type DomainSni struct {
	Id        string `json:"id"`
	PageIndex int    `json:"pageIndex"`
	PageSize  int    `json:"pageSize"`
}

type SafeCon struct {
	Config Config `json:"config"`
	// 订单uuid
	DDUUID string `json:"dd_uuid"`
	// 域名uuid
	DomainUUID string `json:"domain_uuid"`
	// 订单类型，9.CDN，10.DDoS
	ProType string `json:"pro_type"`
	Active  int64  `json:"active"`
}

type SwitchPreAcc struct {
	// 开关
	Active int64 `json:"active"`
	// 订单uuid
	DDUUID string `json:"dd_uuid"`
	// 域名uuid
	DomainUUID string `json:"domain_uuid"`
	// 订单类型，9.CDN，10.DDoS
	ProType string `json:"pro_type"`
}

type PreAccDel struct {
	// 规则uuid
	AccessUUID string `json:"access_uuid"`
	// 订单uuid
	DDUUID string `json:"dd_uuid"`
	// 域名uuid
	DomainUUID string `json:"domain_uuid"`
	// 订单类型，9.CDN，10.DDoS
	ProType string `json:"pro_type"`
}

type PreAcc struct {
	// 规则uuid
	AccessUUID string `json:"access_uuid"`
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
	ProType int64      `json:"pro_type"`
	Rule    PreAccRule `json:"rule"`
}

type PreAccRule struct {
	// 已经选择的匹配项类型
	CheckList []string      `json:"check_list"`
	Rule      []RuleElement `json:"rule"`
}

type RuleElement struct {
	//
	// 匹配项，path：URL，method：方法，referer：来源URL，ua：User-Agent，statuscode：URL重写，clientip：来源IP，contains_header：自定义请求头，params:URL参数
	MItem string `json:"m_item"`
	// 操作符，equal:等于,notEqual:不等于,contain:包含,notContain:不包含,regular:正则匹配
	MOperate string `json:"m_operate"`
	// 匹配值
	MValue string `json:"m_value"`
	// 匹配值
	MValueXs string `json:"m_value_xs"`
}

type LeechLink struct {
	// 是否开启
	Active string `json:"active"`
	// 允许空Referer
	AllowEmpty string `json:"allow_empty"`
	DDUUID     string `json:"dd_uuid"`
	// 域名uuid
	DomainUUID string `json:"domain_uuid"`
	// 域名列表
	Domains utils.JSON `json:"domains"`
	ProType string     `json:"pro_type"`
	// 黑白名单，true:白名单 1:黑名单
	Type string `json:"type"`
}

type AreaAccCon struct {
	// 开关
	Active string `json:"active"`
	// 订单uuid
	DDUUID string `json:"dd_uuid"`
	// 域名uuid
	DomainUUID string `json:"domain_uuid"`
	// 订单类型，9.CDN，10.DDoS
	ProType string `json:"pro_type"`
	// 地区列表
	Regions utils.JSON `json:"regions"`
}

type Domain struct {
	// 订单uuid
	DDUUID string `json:"dd_uuid"`
	// 域名
	Domain string `json:"domain"`
	// 域名uuid
	DomainUUID string `json:"domain_uuid"`
	// 订单类型，9.CDN，10.DDoS
	ProType int64 `json:"pro_type"`
	// 域名uuid数组
	IDS []string `json:"ids"`
}

type DomainsLists struct {
	// 订单uuid
	DDUUID string `json:"dd_uuid"`
	// 订单类型，9.CDN，10.DDoS
	ProType int64 `json:"pro_type"`
	Page    int64 `json:"page"`
	Limit   int64 `json:"limit"`
}

type DomainDevices struct {
	Config []Config `json:"config"`
	// 订单uuid
	DDUUID string `json:"dd_uuid"`
	// 域名uuid
	DomainUUID string `json:"domain_uuid"`
	// 订单类型，9.CDN，10.DDoS
	ProType int64 `json:"pro_type"`
}

type DomainHeat struct {
	// 订单uuid
	DDUUID string `json:"dd_uuid"`
	// 域名uuid
	DomainUUID string `json:"domain_uuid"`
	// 订单类型，9.CDN，10.DDoS
	ProType int64 `json:"pro_type"`
}

type DomainHeatUpdate struct {
	// 订单uuid
	DDUUID string `json:"dd_uuid"`
	// 域名uuid
	DomainUUID string `json:"domain_uuid"`
	// 订单类型，9.CDN，10.DDoS
	ProType     int64    `json:"pro_type"`
	CacheConfig []string `json:"cache_config"`
}

type Config struct {
	// 负载均衡，0.轮询,1.权重,2.IP哈希,3.URL哈希
	LoadBalancing string `json:"load_balancing"`
	// 并发超限处置-301重定向跳转地址
	OverloadRedirectURL string `json:"overload_redirect_url"`
	// 并发超限处置-自定义阻断状态码，403/404/500等http状态码
	OverloadStatusCode *string `json:"overload_status_code,omitempty"`
	// 并发超限处置-类型，1.400阻断，2.301跳转，3.自定义阻断状态码
	OverloadType string `json:"overload_type"`
	// 端口
	Port string `json:"port"`
	// 协议类型，http：HTTP协议，https：HTTPS协议
	Protocol string `json:"protocol"`
	// 取源协议-301重定向状态，true启用/false禁用
	Redirect string `json:"redirect"`
	// 服务器响应头，默认空
	Server string `json:"server"`
	// 取源协议-数据，redirect = false 生效
	SourceAddresses []SourceAddress `json:"source_addresses"`
	// 取源协议-301重定向跳转地址，redirect = true 生效
	URIForward *string `json:"uri_forward,omitempty"`

	BlockConfig BlockConfig `json:"block_config"`
	// 请求频率限制开关
	ResuestRate ResuestRate `json:"resuest_rate"`
	Site        Site        `json:"site"`

	// 访问密码
	Password string `json:"password"`
	// URI，需去掉域名部分路径，如：/login.php
	URL string `json:"url"`

	// 健康检查开关，字符串true开启/字符串false关闭
	HealthCheckActive string `json:"health_check_active"`
	// 间隔时间，单位：秒
	Interval string `json:"interval"`
	// 频率限制-并发连接数，0不限制
	MaxConnection string `json:"max_connection"`
	// 频率限制-新建连接数，0不限制
	NewConnection string `json:"new_connection"`
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
	// 超时时间，- 单位：秒
	// - 超时时间必须小于间隔时间
	Timeout string `json:"timeout"`
}

type SourceAddress struct {
	// 地址
	Address string `json:"address"`
	// 并发
	Concurrent string `json:"concurrent"`
	// 端口
	Port string `json:"port"`
	// 取源协议，http，https
	Protocol string `json:"protocol"`
	// SNI
	Sni string `json:"sni"`
	// 权重
	Weight string `json:"weight"`
}

type CC struct {
	DomainId int64 `json:"domain_id"`
	OrderId  int64 `json:"order_id"`
	// 整体开关
	Active string `json:"active"`
	Config Config `json:"config"`
	// 订单uuid
	DDUUID string `json:"dd_uuid"`
	// 域名uuid
	DomainUUID string `json:"domain_uuid"`
	// 订单类型，9.CDN，10.DDoS
	ProType string `json:"pro_type"`
	// CC默认配置开关
	UseDefault string `json:"use_default"`
}

type BlockConfig struct {
	// 3层拦截开关
	BlockActive string `json:"block_active"`
	// 封禁时间，封禁时间，单位为秒，默认为3600
	BlockTime string `json:"block_time"`
	// 质询次数，当天累计质询失败超过此次数后移入封禁名单那，默认为3
	Count string `json:"count"`
}

// 请求频率限制开关
type ResuestRate struct {
	// 开关
	Active string `json:"active"`
	// 质询失败后阻断的时间（分钟）
	BlockMinutes string `json:"blockMinutes"`
	// 质询次数限制，质询次数限制。失败若干次数后，直接七层阻断，为0则不限制
	ChallengeLimit string `json:"challengeLimit"`
	// 需要质询的请求方法
	ChallengeMethod utils.JSON `json:"challengeMethod"`
	// 触发请求限制后的质询策略，challengePolicy 修改为 1. 浏览器环境验证（不延迟）- "js" 2. JS延迟验证（延迟5s）- "js_delay" 3.
	// 交互验证（滑动验证码）- "human"
	ChallengePolicy string `json:"challengePolicy"`
	Concurrency     string `json:"concurrency"`
	// 质询使用的cookie名称
	CookieName string `json:"cookieName"`
	// 不触发质询的请求地址后缀
	ExcludeEXT string `json:"excludeExt"`
	// 触发请求限制后的质询保护持续时间分钟
	ProtectMinutes string `json:"protectMinutes"`
	// 请求频率上线 次/秒
	Rate     string             `json:"rate"`
	URLRates ResuestRateURLRate `json:"url_rates"`
	// 质询通过后，免于质询的时间（分钟）
	WhiteMinutes string `json:"whiteMinutes"`
}

type ResuestRateURLRate struct {
	// 对单独URL设置并发量的QPS
	Rate *string `json:"rate,omitempty"`
	// URL地址
	URL *string `json:"url,omitempty"`
}

type Site struct {
	// 开关
	Active string `json:"active"`
	// 阻断时长，触发阻断策略或封禁策略后的阻断时长:分钟
	Blockminutes string `json:"blockminutes"`
	// 网站全局并发
	GlobalConcurrent string `json:"global_concurrent"`
	// 防护策略，("block","阻断"), ("wait","等待"),("redirect","重定向"),("ban","封禁")
	Policy string `json:"policy"`
	// 触发重定向策略后的重定向地址
	Redirectlocation string `json:"redirectlocation"`
	// 重定向等待时长，触发重定向策略后的重定向等待时长（秒）
	Redirectwaitseconds string `json:"redirectwaitseconds"`
	// 访问频率
	URLRates SiteURLRate `json:"url_rates"`
	// 执行时长，触发等待策略后等待策略执行时长（分钟）
	Waitpolicyminutes string `json:"waitpolicyminutes"`
	// 等待时长，触发等待策略后，请求等待时长:秒
	Waitseconds string `json:"waitseconds"`
}

type SiteURLRate struct {
	// 对单独URL设置并发量的QPS
	Rate *string `json:"rate,omitempty"`
	// URL地址，单独URL设置并发量的URL地址
	URL *string `json:"url,omitempty"`
}

type PicRc struct {
	// 开关
	Active int64 `json:"active"`
	// 订单uuid
	DDUUID string `json:"dd_uuid"`
	// 域名uuid
	DomainUUID string `json:"domain_uuid"`
	// 订单类型，9.CDN，10.DDoS
	ProType string `json:"pro_type"`
}

type WordsRc struct {
	// 开关
	Active string `json:"active"`
	// 订单uuid
	DDUUID string `json:"dd_uuid"`
	// 域名uuid
	DomainUUID string `json:"domain_uuid"`
	// 回源gzip标识，根据 content-encoding 的gzip是否开启关闭gzip，如果站点存在gzip，设置为false，会无法解压页面导致功能失效
	Gzip string `json:"gzip"`
	// 关键词列表，关键列表：["test", "mock"]
	Keywords utils.JSON `json:"keywords"`
	// 订单类型，9.CDN，10.DDoS
	ProType string `json:"pro_type"`
}
