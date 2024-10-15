package form

type SSLServiceInfo struct {
	SSLId              int64  `json:"ssl_id"`
	AdminInfo          string `json:"admin_info,omitempty"`
	COMStatus          int64  `json:"com_status"`
	DNSHost            string `json:"dns_host,omitempty"`
	DNSType            string `json:"dns_type,omitempty"`
	DNSValue           string `json:"dns_value,omitempty"`
	DomainList         string `json:"domain_list,omitempty"`
	DomainNum          string `son:"domain_num,omitempty"`
	DomainType         string `json:"domain_type,omitempty"`
	FileName           string `json:"file_name,omitempty"`
	FileValue          string `json:"file_value,omitempty"`
	OrderStart         int64  `json:"order_start,omitempty"`
	OrgInfo            string `json:"org_info,omitempty"`
	PMethod            string `json:"p_method,omitempty"`
	PTypeName          string `json:"p_type_name,omitempty"`
	SetupServer        int64  `json:"setup_server,omitempty"`
	SSLCode            string `json:"ssl_code,omitempty"`
	SSLCsr             string `json:"ssl_csr,omitempty"`
	SSLKey             string `json:"ssl_key,omitempty"`
	SSLPem             string `json:"ssl_pem"`
	SSLType            string `json:"ssl_type,omitempty"`
	TechInfo           string `json:"tech_info,omitempty"`
	UUID               string `json:"uuid,omitempty"`
	XufeiOrderid       string `json:"xufei_orderid,omitempty"`
	YnProve            int64  `json:"yn_prove,omitempty"`
	YnReplace          int64  `json:"yn_replace,omitempty"`
	YnXufei            string `json:"yn_xufei,omitempty"`
	ZDomain            string `json:"z_domain,omitempty"`
	OrderId            string `json:"order_id,omitempty"`
	OrderName          string `json:"order_name,omitempty"`
	SSLName            string `json:"ssl_name,omitempty"`
	ZDomainList        string `json:"z_domain_list,omitempty"`
	KsMoney            int64  `json:"ks_money,omitempty"`
	EdTime             string `json:"ed_time,omitempty"`
	DateDiff           int64  `json:"date_doff,omitempty"`
	Img                string `json:"img,omitempty"`
	SSLTypeDetail      string `json:"ssl_type_detail"`
	SetUpServiceDetail string `json:"setup_service_detail"`
	AdminName          string `json:"admin_name"`
	AdminTel           string `json:"admin_tel"`
	AdminEmail         string `json:"admin_email"`
	AdminJob           string `json:"admin_job"`

	Create string `json:"create"`
	CaNum  string `json:"ca_num"`
	UserId int64  `json:"user_id"`
	Agent  string `json:"agent"`
}

type SSLPurchaseInfo struct {
	Id          string `json:"id"`
	LoadDomain  string `json:"loaddomain"`
	LoadKey     string `json:"loadkey"`
	LoadCsr     string `json:"loadcsr"`
	ProveMethod string `json:"provemethod"`
	DomainList  string `json:"domain_list"`
	LastName    string `json:"lastname"`
	FirstName   string `json:"firstname"`
	TelNum      string `json:"tel_num"`
	Email       string `json:"email"`
	Job         string `json:"job"`
	CsrState    string `json:"csrstate"`
	CsrLocality string `json:"csrlocality"`
	CsrOrg      string `json:"csrorg"`
	Department  string `json:"department"`
}
