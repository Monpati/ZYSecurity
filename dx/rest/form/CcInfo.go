package form

import "Dexun/utils"

type CCInfo struct {
	CcId                int64      `json:"cc_id"`
	DomainId            int64      `json:"domain_id"`
	DDUuid              string     `json:"dd_uuid"`
	OrderId             int64      `json:"order_id"`
	DomainUuid          string     `json:"domain_uuid"`
	ProType             string     `json:"pro_type"`
	CsActive            string     `gorm:"column:cs_active" json:"cs_active"`
	Policy              string     `json:"policy"`
	Url                 string     `json:"url"`
	Rate                string     `json:"rate"`
	WaitSeconds         string     `json:"waitseconds"`
	BlockMinutes        string     `json:"blockminutes"`
	RedirectLocation    string     `json:"redirectlocation"`
	GlobalConcurrent    string     `json:"global_concurrent"`
	WaitPolicyMinutes   string     `json:"waitpolicyminutes"`
	RedirectWaitSeconds string     `json:"redirectwaitseconds"`
	Count               string     `json:"count"`
	BlockTime           string     `json:"block_time"`
	BlockActive         string     `json:"block_active"`
	RrActive            string     `json:"rr_active"`
	RrRate              string     `json:"rr_rate"`
	UrUrl               string     `json:"ur_url"`
	UrRate              string     `json:"ur_rate"`
	CookieName          string     `json:"cookieName"`
	ExcludeExt          string     `json:"excludeExt"`
	Concurrency         string     `json:"concurrency"`
	RBlockMinutes       string     `json:"r-blockMinutes"`
	WhiteMinutes        string     `json:"whiteMinutes"`
	ChallengeLimit      string     `json:"challengeLimit"`
	ProtectMinutes      string     `json:"protectMinutes"`
	ChallengeMethods    utils.JSON `json:"challengeMethods"`
	ChallengePolicy     string     `json:"challengePolicy"`
	Active              string     `json:"active"`
	UseDefault          string     `json:"use_default"`
}

type CcFilterForm struct {
	CsActive string `json:"cs_active"`
	Policy   string `json:"policy"`
	Url      string `json:"url"`
	Field    string `json:"field"`
	Value    string `json:"value"`
	PageForm
}
