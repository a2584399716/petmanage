package models

import "time"

type Good struct {
	Good_name              string  `json:"good_name" db:"good_name" binding:"required"`
	Bar_Code               string  `json:"bar_code" db:"bar_code"`
	CostPrice              float64 `json:"cost_price,string" db:"cost_price" binding:"required"`
	Selling_Price          float64 `json:"selling_price,string" db:"selling_price" binding:"required"`
	Good_id                int64   `json:"good_id,omitempty,string" db:"good_id"`
	Stock                  int     `json:"stock,string" db:"stock" binding:"required"`
	Good_Type              int8    `json:"good_type,string" db:"good_type"`
	SuitableType           int8    `json:"suitable_type,string" db:"suitable_type"`
	Good_Supplier          int8    `json:"good_supplier,string" db:"good_supplier"`
	NegativeSalesInventory int8    `json:"negative_sales_inventory,string" db:"negative_sales_inventory"` //是否支持负库存销售，0为不支持，1为支持，默认不支持
}

type ApiGoodDetail struct {
	Good_type_name string `json:"good_type_name" db:"type_name"`
	*Good
}

type GoodType struct {
	TypeId   int8   `json:"type_id,omitempty" db:"type_id"`
	TypeName string `json:"type_name" db:"type_name"`
}

type Supplier struct {
	SupplierId    int8      `json:"supplier_id,omitempty,string" db:"supplier_id"`
	SupplierName  string    `json:"supplier_name" db:"supplier_name"`
	Contacts      string    `json:"contacts" db:"contacts"` //联系人
	SupplierPhone string    `json:"supplier_phone" db:"supplier_phone"`
	CreateTime    time.Time `json:"create_time,omitempty" db:"create_time"`
}

type SuitableType struct {
	SuitableId int8 `json:"suitable_id,omitempty" db:"suitable_id"`
	SuitableName string `json:"suitable_name" db:"suitable_name"`
}
