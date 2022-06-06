package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "root:12345678@tcp(127.0.0.1)/blogdb?charset=utf8&parseTime=True&loc=Local"
	_db, _err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if _err != nil {
		panic(_err)
	}
	db = _db
}

func GetDB() *gorm.DB {
	return db
}
