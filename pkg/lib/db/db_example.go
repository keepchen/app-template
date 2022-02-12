package db

import (
	"github.com/keepchen/app-template/pkg/common/db/models"
)

//ExampleUsage 使用示例
func ExampleUsage() {
	conf := Conf{
		DriverName: "mysql",
		Mysql: MysqlConf{
			Host:      "localhost",
			Port:      3306,
			Database:  "default",
			Username:  "fool",
			Password:  "bar",
			Charset:   "utf8mb4",
			ParseTime: true,
			Loc:       "Local",
		},
	}

	InitDB(conf)

	var user models.User
	GetInstance().Model(&models.User{}).First(&user)
}
