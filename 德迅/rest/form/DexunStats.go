package form

type SCDNDDoS struct {
	DDUUID string `json:"dd_uuid"`
}

type AtkKind struct {
	//将 Unit 写死为hour，day数据通过hour数据整合

	// 订单uuid
	DDUUID string `json:"dd_uuid"`
	// 结束时间，unit=day 必传，秒级时间戳
	EndTime int64 `json:"end_time,omitempty"`
	// 区间，unit=hour 必传，以当前时间往前推算
	Interval int64 `json:"interval,omitempty"`
	// 产品类型
	ProType int64 `json:"pro_type"`
	// 开始时间，unit=day 必传，秒级时间戳
	StartTime int64 `json:"start_time,omitempty"`
	// 类型，| 类型 | 数值单位   |
	// | ---- | ---------- |
	// | hour | 按小时展示 |
	// | day  | 按小时展示 |
	Unit string `json:"unit"`
}

type IpRanking struct {
	// 订单uuid
	DDUUID string `json:"dd_uuid"`
	// 结束时间，unit=day 必传，秒级时间戳
	EndTime int64 `json:"end_time,omitempty"`
	// 产品类型
	ProType int64 `json:"pro_type"`
	// 开始时间，unit=day 必传，秒级时间戳
	StartTime int64 `json:"start_time,omitempty"`
}

type Info struct {
	// 订单uuid
	DDUUID string `json:"dd_uuid"`
	// 订单类型，9.CDN，10.DDoS
	ProType int64 `json:"pro_type"`
}

type InterceptStatsInfo struct {
	OrderId    int64 `json:"order_id"`
	AppCC      int64 `json:"app_cc"`
	CC         int64 `json:"cc"`
	IpBlack    int64 `json:"ip_black"`
	Referer    int64 `json:"referer"`
	UrlBlack   int64 `json:"url_black"`
	WebProtect int64 `json:"web_protect"`
	Other      int64 `json:"other"`
	AreaAcc    int64 `json:"area_acc"`
	SafeAcc    int64 `json:"safe_acc"`
	PreAcc     int64 `json:"pre_acc"`
}

type AtkStatsInfo struct {
	Time       int64 `json:"time"`
	OrderId    int64 `json:"order_id"`
	TotalCount int64 `json:"total_count"`
}

type AreaStatsInfo struct {
	OrderId int64  `json:"order_id"`
	Source  string `json:"source"`
	Count   int64  `json:"count"`
}

type AreaRankStatsInfo struct {
	OrderId    int64  `json:"order_id"`
	Domain     string `json:"domain"`
	TotalCount int64  `json:"total_count"`
}

type AtkInterStatsInfo struct {
	OrderId    int64  `json:"order_id"`
	Domain     string `json:"domain"`
	TotalCount int64  `json:"total_count"`
}

type HttpPackStatsInfo struct {
	OrderId    int64 `json:"order_id"`
	Time       int64 `json:"time"`
	TotalCount int64 `json:"total_count"`
}

type LineChartStatsInfo struct {
	OrderId      int64 `json:"order_id"`
	Time         int64 `json:"time"`
	ResponseSize int64 `json:"response_size"`
	RequestSize  int64 `json:"request_size"`
}

type BWStatsInfo struct {
	OrderId    int64  `json:"order_id"`
	IP         string `json:"ip"`
	BwListType string `json:"bw_list_type"`
	TotalCount int64  `json:"total_count"`
}

type TotalFlowInfo struct {
	OrderId               int64 `json:"order_id"`
	RequestBandWidthPeak  int64 `json:"request_bandwidth_peak"`
	Requests              int64 `json:"requests"`
	ResponseBandWidthPeak int64 `json:"response_bandwidth_peak"`
	TotalRequestFlows     int64 `json:"total_request_flows"`
	TotalResponseFlows    int64 `json:"total_response_flows"`
	UnidentifiedAttack    int64 `json:"unidentified_attack"`
}

type AccCDNRankInfo struct {
	OrderId  int64  `json:"order_id"`
	ClientIp string `json:"client_ip"`
	CountSum int64  `json:"count_sum"`
}
