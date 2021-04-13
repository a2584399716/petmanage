package mysql

import "web_app/models"

func GetUserList(page,limit int64)(data []*models.User,count int,err error){
	sqlStr := `select user_id, user_name,user_phone ,user_balance,create_time from user ORDER BY create_time desc limit ?,?`
	sqlStr2 := `select count(*) from user`
	data = make([]*models.User,0,2)
	err = db.Select(&data, sqlStr, (page-1)*limit, limit)
	err=db.Get(&count,sqlStr2)
	return
}
