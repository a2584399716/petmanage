package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"web_app/models"
)

const secret = "authorChen"

func Login(admin *models.Admin) (err error) {
	oldPassword := admin.Password //用户登录的密码

	sqlStr := `select username,password from admin where username = ?`
	err = db.Get(admin, sqlStr, admin.Username)
	if err == sql.ErrNoRows {
		return errors.New("用户不存在")
	}
	if err != nil {
		//查询数据库失败
		return err
	}
	//判断密码是否正确

	password := encryptPassword(oldPassword)
	if password != admin.Password {
		return errors.New("密码错误")
	}
	return

}

func encryptPassword(oldPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oldPassword)))
}
