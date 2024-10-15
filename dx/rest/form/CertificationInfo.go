package form

type CertificationInfo struct {
	Ddc_id     int64  `gorm:"column:ddc_id" json:"ddc_Id"`
	OrderId    int64  `gorm:"column:order_id" json:"order_id"`
	DomainId   int64  `gorm:"column:domain_id" json:"domain_id"`
	DomainUuid string `gorm:"column:domain_uuid" json:"domain_uuid"`
	Status     int64  `gorm:"column:status" json:"status"`
	SslAlways  int64  `gorm:"column:ssl_always" json:"ssl_always"`
	Hsts       int64  `gorm:"column:hsts" json:"hsts"`
	CertName   string `gorm:"column:cert_name" json:"cert_name"`
	Cert       string `gorm:"column:cert" json:"cert"`
	Key        string `gorm:"column:key" json:"key"`
	Desc       string `gorm:"column:desc" json:"desc"`
	Createtime string `gorm:"column:createtime" json:"createtime"`
	UpdateTime string `gorm:"column:updatetime" json:"updatetime"`
}

type CertificationFilterForm struct {
	CertName string `json:"cert_name"`
	Field    string `json:"field"`
	Value    string `json:"value"`
	PageForm
}
