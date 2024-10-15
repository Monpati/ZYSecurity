package form

type CacheInfo struct {
	DxCacheId       int64  `json:"dxcache_id"`
	CacheId         int64  `json:"cache_id"`
	DdUuid          string `json:"dd_uuid"`
	OrderId         int64  `json:"order_id"`
	ProType         string `json:"pro_type"`
	CacheUuid       string `json:"cache_uuid"`
	DomainUuid      string `json:"domain_uuid"`
	DomainId        int64  `json:"domain_id"`
	Active          string `json:"active"`
	UrlMode         string `json:"urlmode"`
	CacheMode       string `json:"cachemode"`
	CachePath       string `json:"cachepath"`
	CacheExtensions string `json:"cacheextensions"`
	CacheReg        string `json:"cachereg"`
	TimeOut         string `json:"timeout"`
	Weight          string `json:"weight"`
	CreateTime      string `json:"createtime"`
	UpdateTime      string `json:"updatetime"`
	Status          int    `json:"status"`
}

type CacheFilterForm struct {
	UrlMode   string `json:"urlmode"`
	CacheMode string `json:"cachemode"`
	Field     string `json:"field"`
	Value     string `json:"value"`
	PageForm
}
