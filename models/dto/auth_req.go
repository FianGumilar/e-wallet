package dto

type AuthReq struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
