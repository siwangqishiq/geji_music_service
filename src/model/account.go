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
