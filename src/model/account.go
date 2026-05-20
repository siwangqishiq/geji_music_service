package model

type CreateAccountReq struct {
	Account  string `json:"account" binding:"required"`
	Nickname string `json:"nickname"`
	Password string `json:"password" binding:"required"`
	Avatar   string `json:"avatar"`
}

type LoginReq struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResp struct {
	Uid      int64  `json:"uid"`
	Account  string `json:"account"`
	Nickname string `json:"nickname"`
	Remark   string `json:"remark"`
	Age      string `json:"age"`
	Avatar   string `json:"avatar"`
	Token    string `json:"token"`
}
