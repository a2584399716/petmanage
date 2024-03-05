package mysql

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"web_app/models"
	"web_app/settings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Init(cfg *settings.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db2, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db2.AutoMigrate(&models.Admin{},
		&models.UserInfo{},
		models.Customer{},
		&models.Good{},
		&models.ApiGoodDetail{},
		&models.GoodType{},
		&models.Supplier{},
		&models.SuitableType{},
		&models.Varietie{},
		&models.LivingAnimal{},
		&models.ApiLivingDetail{},
		&models.Serve{},
		&models.ApiServeDetail{},
		&models.ServeType{},
		&models.User{},
	)
	if err != nil {
		panic(err)
	}
	return
}

func Close() {
	_ = db.Close()
}
