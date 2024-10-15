package form

type Log struct {
	// 订单uuid
	DDUUID string `json:"dd_uuid"`
	// 结束时间，单位：毫秒
	End   int64 `json:"end"`
	Limit int64 `json:"limit"`
	Page  int64 `json:"page"`
	// 产品类型
	ProType int64 `json:"pro_type"`
	// 开始时间，单位：毫秒
	Start int64 `json:"start"`
}

type QueryLog struct {
	// 订单uuid
	DDUUID string `json:"dd_uuid"`
	// 结束时间，unit=day 必传，秒级时间戳
	EndTime int64 `json:"end_time"`
	// 单页数量
	Limit int64 `json:"limit"`
	// 页码
	Page int64 `json:"page"`
	// 产品类型
	ProType int64 `json:"pro_type"`
	// 开始时间，unit=day 必传，秒级时间戳
	StartTime int64 `json:"start_time"`
}

type OtherLog struct {
	// 订单uuid
	DDUUID string `json:"dd_uuid"`
	// 结束时间，unit=day 必传，秒级时间戳
	EndTime int64 `json:"end_time,omitempty"`
	// 产品类型
	ProType int64 `json:"pro_type"`
	// 开始时间，unit=day 必传，秒级时间戳
	StartTime int64 `json:"start_time,omitempty"`
}

type AtkLogsInfo struct {
	AttackInfo     string `json:"attackinfo"`
	AttackType     string `json:"attacktype"`
	ClientIp       string `json:"clientip"`
	ClientPort     int64  `json:"clientport"`
	ClientRegion   string `json:"clientregion"`
	Count          int64  `json:"count"`
	Domain         string `json:"domain"`
	DomainId       string `json:"doaminid"`
	OrderId        int64  `json:"order_id"`
	HttpMethod     string `json:"httpmethod"`
	AlId           int64  `json:"al_id"`
	InstanceId     int64  `json:"instanceid"`
	LocalIp        string `json:"localip"`
	Method         string `json:"method"`
	ProtectType    string `json:"protecttype"`
	RequestInfo    string `json:"requestinfo"`
	TargetUrl      string `json:"targeturl"`
	TimeRangeEnd   string `json:"timerangeend"`
	TimeRangeStart string `json:"timerangestart"`
}

type AccLogsInfo struct {
	CacheHit           string      `json:"cachehit"`
	ClientIp           string      `json:"clientip"`
	ClientPort         int64       `json:"clientport"`
	ClientRegion       string      `json:"clientregion"`
	Count              int64       `json:"count"`
	CreateDat          interface{} `json:"createdat"`
	Domain             string      `json:"domain"`
	DomainId           string      `json:"domainid"`
	Form               string      `json:"form"`
	DxId               int64       `json:"dx_id"`
	InstanceId         int64       `json:"instanceid"`
	LocalAddr          string      `json:"localaddr"`
	LocalIp            string      `json:"localip"`
	LocalPort          int64       `json:"localport"`
	Method             string      `json:"method"`
	Nodeid             string      `json:"nodeid"`
	Packagesize        int64       `json:"packagesize"`
	Remoteaddr         string      `json:"remoteaddr"`
	Responsesize       int64       `json:"responsesize"`
	Responsestatuscode int64       `json:"responsestatuscode"`
	Timerangeend       string      `json:"timerangeend"`
	Timerangestart     string      `json:"timerangestart"`
	URL                string      `json:"url"`
	Wblist             string      `json:"wblist"`

	Accept                  string  `json:"Accept"`
	AcceptEncoding          string  `json:"accept-encoding"`
	AcceptLanguage          string  `json:"accept-language"`
	Authorization           string  `json:"authorization"`
	CacheControl            string  `json:"cache-control"`
	Connection              string  `json:"connection"`
	Pragma                  string  `json:"pragma"`
	Purpose                 *string `json:"purpose,omitempty"`
	Referer                 string  `json:"referer"`
	UpgradeInsecureRequests *string `json:"upgrade-insecure-requests,omitempty"`
	UserAgent               string  `json:"user-agent"`
	XForwardedHost          string  `json:"x-forwarded-host"`
	XForwardedPort          string  `json:"x-forwarded-port"`
	XForwardedProto         string  `json:"x-forwarded-proto"`
	XForwardedServer        string  `json:"x-forwarded-server"`
	XRealIP                 string  `json:"x-real-ip"`

	AcceptRanges    *string `json:"accept-ranges,omitempty"`
	ContentEncoding string  `json:"content-encoding"`
	ContentLength   string  `json:"content-length"`
	ContentType     string  `json:"content-type"`
	Date            string  `json:"date"`
	Etag            string  `json:"etag"`
	LastModified    string  `json:"Last-Modified"`
	Server          string  `json:"server"`
	Vary            string  `json:"vary"`
	WWWAuthenticate string  `json:"www-authenticate"`
}

type FlowLogsInfo struct {
	OrderId      int64  `json:"order_id"`
	Domain       string `json:"domain"`
	RequestSize  int64  `json:"request_size"`
	ResponseSize int64  `json:"response_size"`
}

type QueryLogsInfo struct {
	CacheCalls int64  `json:"cache_calls"`
	CacheRate  int64  `json:"cache_rate"`
	Domain     string `json:"domain"`
	DomainId   int64  `json:"domain_id"`
	TotalCalls int64  `json:"total_calls"`
}
