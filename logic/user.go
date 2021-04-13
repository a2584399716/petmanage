package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
)

func GetUserList(page,limit int64)(data []*models.User,count int,err error){
	data,count,err = mysql.GetUserList(page,limit)
	if err != nil {
		return nil, 0,err
	}
	return
}
