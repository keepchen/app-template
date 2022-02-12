package db

import (
	"time"

	"gorm.io/gorm"
)

//数据库连接实例
var dbInstance *gorm.DB

//InitDB 初始化数据库连接
func InitDB(fields DsnFields) {
	dialect := fields.GenDialector()
	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	dbInstance = db
}

//GetInstance 获取数据库实例
func GetInstance() *gorm.DB {
	return dbInstance
}
