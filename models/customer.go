package models

type Customer struct {
	IdValue int64 `json:"value" db:"user_id"`
	NameAndPhone string `json:"name" db:"info"`
}
