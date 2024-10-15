package dexun

import "Dexun/model/Dexun"

func CertInfo(gname, gnum string) (string, string, string, string) {
	var d Dexun.DeXunBody
	d.GetPersonCertDetails(gname, gnum)
	return d.SmrzBirthday, d.SmrzSex, d.SmrzGnum, d.SmrzCity
}
