package form

type CacheListInfo struct {
	CmId            int64  `json:"cm_id"`
	DdId            int64  `json:"dd_id"`
	DdType          int64  `json:"dd_type"`
	CacheName       string `json:"cache_name"`
	Active          string `json:"active"`
	UrlMode         string `json:"urlmode"`
	CacheMode       string `json:"cachemode"`
	CachePath       string `json:"cachepath"`
	Cacheextensions string `json:"cacheextensions"`
	CacheReg        string `json:"cachereg"`
	TimeOut         int64  `json:"timeout"`
	Weight          int64  `json:"weight"`
	Status          int    `json:"status"`
}

type CacheListFilterForm struct {
	UrlMode   string `json:"urlmode"`
	CacheMode string `json:"cachemode"`
	Field     string `json:"field"`
	Value     string `json:"value"`
	PageForm
}
