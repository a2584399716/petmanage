package models

type Admin struct {
	Username string `db:"username"`
	Password string `db:"password"`
}
