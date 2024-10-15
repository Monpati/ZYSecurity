package dexun

import (
	"Dexun/model/Dexun"
)

func Api() string {
	var d Dexun.DeXunBody
	d.ApiLogin()
	return d.Token
}
