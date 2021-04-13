package mysql

import (
	"web_app/models"
)

func GetGoodList(page, limit int64) (goods []*models.Good, count int, err error) {
	sqlStr := `select 
	good_id, good_name, suitable_type,good_supplier,bar_code, good_type,negative_sales_inventory, cost_price,selling_price,stock
	from goods
	ORDER BY create_time
	DESC
	limit ?,?
	`
	sqlStr2 := `select count(*) from goods`

	goods = make([]*models.Good, 0, 2)
	err = db.Select(&goods, sqlStr, (page-1)*limit, limit)
	err = db.Get(&count, sqlStr2)
	return
}

func GetGoodListByGoodTypeAndNameOrBarcode(page int64, limit int64, goodType int8, nameorBarcode string) (goods []*models.Good, count int, err error) {
	sqlStr := `SELECT good_id, good_name, bar_code, good_type,negative_sales_inventory, cost_price,selling_price,stock FROM goods 
	where concat(good_name,IFNULL(bar_code,''))
	like ? and good_type = ?
	ORDER BY create_time
	DESC
	limit ?,?`
	sqlStr2 := `select count(*) from goods
				where concat(good_name,IFNULL(bar_code,''))
				like ? and good_type = ?`
	goods = make([]*models.Good, 0, 2)
	err = db.Select(&goods, sqlStr, "%"+nameorBarcode+"%", goodType, (page-1)*limit, limit)
	err = db.Get(&count, sqlStr2, "%"+nameorBarcode+"%", goodType)
	return
}

func GetGoodListByGoodType(page int64, limit int64, goodType int8) (goods []*models.Good, count int, err error) {
	sqlStr := `SELECT good_id, good_name, bar_code, good_type,negative_sales_inventory, cost_price,selling_price,stock FROM goods 
				where good_type = ?
				ORDER BY create_time
				DESC
				limit ?,?`
	sqlStr2 := `select count(*) from goods
				where good_type = ?`
	goods = make([]*models.Good, 0, 2)
	err = db.Select(&goods, sqlStr, goodType, (page-1)*limit, limit)
	err = db.Get(&count, sqlStr2, goodType)

	return
}

func GetGoodListByNameOrBarcode(page int64, limit int64, nameorBarcode string) (goods []*models.Good, count int, err error) {
	sqlStr := `SELECT user_id, CONCAT(IFNULL(user_name,''),IFNULL(user_phone,''))
				FROM user 
				where CONCAT(IFNULL(user_name,''),IFNULL(user_phone,''))
				like ?
				ORDER BY create_time
				DESC
	limit ?,?`
	sqlStr2 := `select count(*) from goods
				where concat(good_name,IFNULL(bar_code,''))
				like ?`
	goods = make([]*models.Good, 0, 2)
	err = db.Select(&goods, sqlStr, "%"+nameorBarcode+"%", (page-1)*limit, limit)
	err = db.Get(&count, sqlStr2, "%"+nameorBarcode+"%")
	return
}

func GetGoodTypeNameByTypeID(id int8) (good_type_name string, err error) {
	sqlStr := `select type_name from goods_type where type_id = ?`
	err = db.Get(&good_type_name, sqlStr, id)
	return
}

func GetGoodsTypeList(page, limit int64) (data []*models.GoodType, count int, err error) {
	sqlStr := `select type_id, type_name from goods_type ORDER BY type_id ASC limit ?,?`
	sqlStr2 := `select count(*) from goods_type`
	data = make([]*models.GoodType, 0, 2)
	err = db.Select(&data, sqlStr, (page-1)*limit, limit)

	err = db.Get(&count, sqlStr2)
	return
}

func GetSupplierList(page, limit int64) (data []*models.Supplier, count int, err error) {
	sqlStr := `select supplier_id, contacts,supplier_phone ,supplier_name,create_time from supplier ORDER BY create_time desc limit ?,?`
	sqlStr2 := `select count(*) from supplier`
	data = make([]*models.Supplier, 0, 2)
	err = db.Select(&data, sqlStr, (page-1)*limit, limit)
	err = db.Get(&count, sqlStr2)
	return
}

func GetSuitableType(page, limit int64) (data []*models.SuitableType, count int, err error){
	sqlStr := `select suitable_id, suitable_name from suitable_type limit ?,?`
	sqlStr2 := `select count(*) from suitable_type`
	data = make([]*models.SuitableType, 0, 2)
	err = db.Select(&data, sqlStr, (page-1)*limit, limit)
	err = db.Get(&count, sqlStr2)
	return
}

func AddGoods(g *models.Good) (err error) {

	sqlStr := `insert into goods(
		good_name, bar_code, good_type, good_supplier, suitable_type,cost_price,selling_price,stock,negative_sales_inventory)
	values (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err = db.Exec(sqlStr, g.Good_name, g.Bar_Code, g.Good_Type, g.Good_Supplier, g.SuitableType, g.CostPrice, g.Selling_Price, g.Stock, g.NegativeSalesInventory)
	return
}

func AddSupplier(sup *models.Supplier)(err error){
	sqlStr := `insert into supplier(
				supplier_name, contacts, supplier_phone)
				values (?, ?, ?)
	`
	_, err = db.Exec(sqlStr, sup.SupplierName, sup.Contacts, sup.SupplierPhone)
	return
}

func AddGoodsType(t *models.GoodType)(err error){
	sqlStr := `insert into goods_type (type_name) values (?)
	`
	_, err = db.Exec(sqlStr, t.TypeName)
	return
}

func AddSuitableType(s *models.SuitableType)(err error){
	sqlStr := `insert into suitable_type (suitable_name) values (?)
	`
	_, err = db.Exec(sqlStr, s.SuitableName)
	return
}


func EditGoods(g *models.Good) (err error) {
	sqlStr := `update goods set good_name=?,bar_code=?,good_type=?,good_supplier=?,suitable_type=?,cost_price=?,selling_price=?,negative_sales_inventory=?,stock=? where good_id = ?`
	_, err = db.Exec(sqlStr, g.Good_name, g.Bar_Code, g.Good_Type, g.Good_Supplier, g.SuitableType, g.CostPrice, g.Selling_Price, g.NegativeSalesInventory, g.Stock,g.Good_id)
	return
}

func EditSupplier(sup *models.Supplier) (err error) {
	sqlStr := `update supplier set supplier_name=?,contacts=?,supplier_phone=? where supplier_id = ?`
	_, err = db.Exec(sqlStr, sup.SupplierName,sup.Contacts,sup.SupplierPhone,sup.SupplierId)
	return
}

func SearchPepeatByGoodsName(name string, id int64) (err error) {
	sqlStr := `select count(*) from goods where good_name = ? and good_id <> ? limit 1`
	var count int
	err = db.Get(&count, sqlStr, name, id)
	if count != 0 {
		err = ErrorGoodsNameExist
	}
	return
}

func SearchPepeatBySupplier(name string, id int8) (err error) {
	sqlStr := `select count(*) from supplier where supplier_name = ? and supplier_id <> ? limit 1`
	var count int
	err = db.Get(&count, sqlStr, name, id)
	if count != 0 {
		err = ErrorGoodsNameExist
	}
	return
}

func SearchPepeatBySupplierName(name string,id int8)(err error){
	sqlStr := `select count(*) from supplier where supplier_name = ? and supplier_id <> ? limit 1`
	var count int
	err = db.Get(&count, sqlStr, name, id)
	if count != 0 {
		err = ErrorGoodsNameExist
	}
	return
}
func SearchPepeatByGoodsTypeName(name string)(err error){
	sqlStr := `select count(*) from goods_type where type_name = ?  limit 1`
	var count int
	err = db.Get(&count, sqlStr, name)
	if count != 0 {
		err = ErrorGoodsNameExist
	}
	return
}


func SearchPepeatBySuitableTypeName(name string)(err error){
	sqlStr := `select count(*) from suitable_type where suitable_name = ?  limit 1`
	var count int
	err = db.Get(&count, sqlStr, name)
	if count != 0 {
		err = ErrorGoodsNameExist
	}
	return
}
func DelGoodsById(id int64) (err error) {
	sqlStr := `delete from goods where good_id = ?`
	_, err = db.Exec(sqlStr, id)
	return

}

func DelGoodsTypeById(id int8)(err error){
	sqlStr := `delete from goods_type where type_id = ?`
	_, err = db.Exec(sqlStr, id)
	return
}

func DelSuitableTypeById(id int8)(err error){
	sqlStr := `delete from suitable_type where suitable_id = ?`
	_, err = db.Exec(sqlStr, id)
	return
}

func DelSupplierById(id int8)(err error){
	sqlStr := `delete from supplier where supplier_id = ?`
	_, err = db.Exec(sqlStr, id)
	return
}