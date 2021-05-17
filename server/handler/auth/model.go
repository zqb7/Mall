package auth

type ReqRegister struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Birthday int64  `json:"birthday"`
	Mobile   string `json:"mobile"`
	Avatar   string `json:"avatar" binding:"max=255"`
}

type ReqLogin struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}
