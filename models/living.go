package models

type Varietie struct {
	VarietieId int8 `json:"varietie_id" db:"varietie_id"`
	VarietieName string `json:"varietie_name" db:"varietie_name"`
	VarietieNumber int `json:"varietie_number" db:"varietie_number"`
}

type LivingAnimal struct {
	LivingId int64 `json:"living_id,string" db:"living_id"`
	LivingName string `json:"living_name" db:"living_name"`
	VarietieId int8 `json:"varietie_id,string" db:"varietie_id,string"`
	LivingSex int8 `json:"living_sex,string" db:"living_sex,string"`
	CostPrice float64 `json:"cost_price,string" db:"cost_price"`
	LivingPrice float64 `json:"selling_price,string" db:"selling_price"`
	LivingState int8 `json:"living_state,string" db:"living_state"`
	CreateTime string `json:"create_time,omitempty" db:"create_time"`
}

type ApiLivingDetail struct {
	VarietieName string `json:"varietie_name" db:"varietie_name"`
	SexName string `json:"sex_name" db:"sex_name"`
	StateName string `json:"state_name" db:"state_name"`
	*LivingAnimal
}
