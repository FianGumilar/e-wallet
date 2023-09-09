package entitty

type User struct {
	ID       int64  `db:"id"`
	Username string `db:"username"`
	Phone    string `db:"phone"`
	Password string `db:"password"`
}
