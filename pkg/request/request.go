package request

// RegistrationReq Incoming reference
type RegistrationReq struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	HeaderImg string `json:"header_img"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type FondArticleReq struct {
	UserUuid string `json:"user_uuid"`
	Title    string `json:"title"`
	Context  string `json:"context"`
	Tag      []int  `json:"tag"`
	Category int    `json:"Category"`
}

type GetArticleReq struct {
	UserUuid string `json:"user_uuid"`
	Page     int    `json:"page"`
	Size     int    `json:"size"`
}
