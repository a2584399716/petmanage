package mysql

import "web_app/models"

func GetServeList(page,limit int64)(data []*models.Serve, count int,err error){
	sqlStr := `select serve_id, serve_name,selling_price,serve_type  from serve ORDER BY create_time desc limit ?,?`
	sqlStr2 := `select count(*) from serve`
	data = make([]*models.Serve,0,2)
	err = db.Select(&data, sqlStr, (page-1)*limit, limit)
	err=db.Get(&count,sqlStr2)
	return
}

func GetServeTypeList(page,limit int64)(data []*models.ServeType, count int,err error){
	sqlStr := `select type_id, type_name from serve_type ORDER BY create_time desc limit ?,?`
	sqlStr2 := `select count(*) from serve_type`
	data = make([]*models.ServeType,0,2)
	err = db.Select(&data, sqlStr, (page-1)*limit, limit)
	err=db.Get(&count,sqlStr2)
	return

}

func GetServeTypeNameByTypeID(id int8)(serve_type_name string,err error){
	sqlStr := `select type_name from serve_type where type_id = ?`
	err = db.Get(&serve_type_name,sqlStr,id)
	return
}

func SearchPepeatByServiceName(name string,id int64)(err error){
	sqlStr := `select count(*) from serve where serve_name = ? and serve_id <> ? limit 1`
	var count int
	err = db.Get(&count, sqlStr, name, id)
	if count != 0 {
		err = ErrorGoodsNameExist
	}
	return
}

func SearchPepeatByServiceTypeName(name string,id int8)(err error){
	sqlStr := `select count(*) from serve_type where type_name = ? and type_id <> ? limit 1`
	var count int
	err = db.Get(&count, sqlStr, name, id)
	if count != 0 {
		err = ErrorGoodsNameExist
	}
	return
}

func AddService(s *models.Serve)(err error){
	sqlStr := `insert into serve(
				serve_name, selling_price, serve_type)
				values (?, ?, ?)
	`
	_, err = db.Exec(sqlStr, s.ServeName, s.ServePrice, s.ServeType)
	return
}

func AddServiceType(s *models.ServeType)(err error){
	sqlStr := `insert into serve_type(type_name)
				values (?)
	`
	_, err = db.Exec(sqlStr, s.TypeName)
	return
}

func EditServiceById(g *models.Serve) (err error) {
	sqlStr := `update serve set serve_name=?,selling_price=?,serve_type=? where serve_id = ?`
	_, err = db.Exec(sqlStr,g.ServeName,g.ServePrice,g.ServeType,g.ServeId)
	return
}

func DelServiceById(id int64)(err error){
	sqlStr := `delete from serve where serve_id = ?`
	_, err = db.Exec(sqlStr, id)
	return
}

func DelServiceTypeById(id int8)(err error){
	sqlStr := `delete from serve_type where type_id = ?`
	_, err = db.Exec(sqlStr, id)
	return
}
