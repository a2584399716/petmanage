package models

//收银台数据
type Cashier struct {
	Id    int64 `json:"id"`
	Name  string `json:"name"`
	Price float64 `json:"price"`
	End   bool   `json:"end"`
}

type Children struct {
	Id    int64    `json:"id"`
	Name  string    `json:"name"`
	Child []*Cashier `json:"children"`
}

type CashierDetail struct {
	Cashier
	Chiledd []*Children `json:"children"`
}

type UserInfo struct {
	Name string `json:"name" db:"user_name"`
	Phone string `json:"phone" db:"user_phone"`
	Balance float64 `json:"balance" db:"user_balance"`
}