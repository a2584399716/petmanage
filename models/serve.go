package models

type Serve struct {
	ServeId int64 `json:"serve_id,string,omitempty" db:"serve_id"`
	ServeType int8 `json:"serve_type,string" db:"serve_type"`
	ServeName string `json:"serve_name" db:"serve_name"`
	ServePrice float64 `json:"selling_price,string" db:"selling_price"`
}

type ApiServeDetail struct {
	Serve_type_name string `json:"serve_type_name" db:"serve_name"`
	*Serve
}

type ServeType struct{
	TypeId int8 `json:"type_id" db:"type_id"`
	TypeName string `json:"type_name" db:"type_name"`
}
