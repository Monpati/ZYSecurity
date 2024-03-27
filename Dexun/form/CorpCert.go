package form

type Corp struct {
	CorpDoc       string `json:"corp_doc"`
	CorpName      string `json:"corp_name"`
	CorpNum       string `json:"corp_num"`
	LgPersonFront string `json:"lgperson_front"`
	LgPersonBack  string `json:"lgperson_back"`
	LgMan         string `json:"lgman"`
	LgPersonNum   string `json:"lgperson_num"`
}

type CorpFilter struct {
	Field string `json:"field"`
	Value string `json:"value"`
	PageForm
}
