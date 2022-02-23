package db

import (
	"time"

	"gorm.io/gorm"
)

//数据库连接实例
var dbInstance *gorm.DB

//InitDB 初始化数据库连接
func InitDB(conf Conf) {
	dialect := conf.GenDialector()
	dbPtr, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := dbPtr.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	dbInstance = dbPtr
}

//GetInstance 获取数据库实例
func GetInstance() *gorm.DB {
	return dbInstance
}
