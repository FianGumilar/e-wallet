package dto

type UserData struct {
	ID       int64  `db:"id"`
	Username string `db:"username"`
	Phone    string `db:"phone"`
}
