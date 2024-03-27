package form

type PageForm struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
type PageOrigin struct {
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
}

type Message struct {
	Content []Content `json:"content"`
	Err     string    `json:"err"`
	Sign    int       `json:"sign"`
	MD5     string    `json:"md5"`
}

type Content struct {
	Cid     string `json:"cid"`
	Title   string `json:"title"`
	About   string `json:"about"`
	Content string `json:"content"`
}
