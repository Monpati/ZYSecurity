package dexun

import (
	"Dexun/model/Dexun"
	"fmt"
)

func Lists(keywords, listrows, ks_money string, page int64) {
	var d Dexun.DeXunBody
	d.GetSCDNList(keywords, listrows, ks_money, page)
	fmt.Println(d)
}
