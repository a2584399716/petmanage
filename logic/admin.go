package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
)

func Login(p *models.ParamLogin) error {
	admin := &models.Admin{
		Username: p.Username,
		Password: p.Password,
	}
	return mysql.Login(admin)

}
