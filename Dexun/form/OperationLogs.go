package form

import (
	"Dexun/utils"
)

type OperationLog struct {
	UserId    int64      `json:"user_id"`
	ReqUrl    string     `json:"req_url"`
	OriginUrl string     `json:"origin_url"`
	UserAgent string     `json:"user_agent"`
	Request   utils.JSON `json:"request"`
	ReqTime   int        `json:"req_time"`
}
