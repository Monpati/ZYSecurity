package form

type Person struct {
	UserId    int64  `json:"user_id"`
	RealName  string `json:"real_name"`
	CardId    string `json:"card_id"`
	CardFront string `json:"card_front"`
	CardBack  string `json:"card_back"`
}

type PersonFilter struct {
	RealName string `json:"real_name"`
	CardId   string `json:"card_id"`
	Field    string `json:"field"`
	Value    string `json:"value"`
	PageForm
}

type CertResponseData struct {
	Data    Data   `json:"data"`
	Message string `json:"message"`
	Status  int64  `json:"status"`
	Time    int64  `json:"time"`
}

type Data struct {
	// 真实姓名
	FullName string `json:"full_name"`
	// 出生日期
	SmrzBirthday string `json:"smrz_birthday"`
	// 实名地区
	SmrzCity string `json:"smrz_city"`
	// 身份证
	SmrzGnum string `json:"smrz_gnum"`
	// 性别
	SmrzSex string `json:"smrz_sex"`
}
