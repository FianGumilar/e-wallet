package dto

type AuthRes struct {
	UserID int64  `json:"-"`
	Token  string `json:"token"`
}
