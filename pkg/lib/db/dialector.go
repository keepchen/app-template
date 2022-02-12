package db

import (
	"fmt"
	"strings"

	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"

	"gorm.io/gorm"
)

//数据库类型标识
const (
	DriverNameMysql      = "mysql"      //数据库类型标识:mysql
	DriverNamePostgres   = "postgres"   //数据库类型标识:postgres sql
	DriverNameSqlite     = "sqlite"     //数据库类型标识:sqlite
	DriverNameSqlserver  = "sqlserver"  //数据库类型标识:sqlserver
	DriverNameClickHouse = "clickhouse" //数据库类型标识:clickhouse

	//more ...
)

//组装mysql的dsn
//@see https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
func mysqlDsn(conf Conf) string {
	if len(conf.Mysql.Charset) == 0 {
		conf.Mysql.Charset = "utf8mb4"
	}
	if len(conf.Mysql.Loc) == 0 {
		conf.Mysql.Loc = "Local"
	}
	var parseTime string
	if conf.Mysql.ParseTime {
		parseTime = "True"
	} else {
		parseTime = "False"
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%d)%s?charset=%s&parseTime=%s&loc=%s",
		conf.Mysql.Username, conf.Mysql.Password, conf.Mysql.Host, conf.Mysql.Port,
		conf.Mysql.Database, conf.Mysql.Charset, parseTime, conf.Mysql.Loc)
}

//组装postgres的dsn
//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
func postgresDsn(conf Conf) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		conf.Postgres.Host, conf.Postgres.Username, conf.Postgres.Password, conf.Postgres.Database,
		conf.Postgres.Port, conf.Postgres.SSLMode, conf.Postgres.TimeZone)
}

//组装sqlite的dsn
//dsn := "sqlite.db"
func sqliteDsn(conf Conf) string {
	return conf.Sqlite.File
}

//组装sqlserver的dsn
//dsn := "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
func sqlserverDsn(conf Conf) string {
	return fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
		conf.Sqlserver.Username, conf.Sqlserver.Password, conf.Sqlserver.Host,
		conf.Sqlserver.Port, conf.Sqlserver.Database)
}

//组装clickhouse的dsn
//dsn := "tcp://localhost:9000?database=gorm&username=gorm&password=gorm&read_timeout=10&write_timeout=20"
func clickhouseDsn(conf Conf) string {
	return fmt.Sprintf("tcp://%s:%d?database=%s&username=%s&password=%s&read_timeout=%d&write_timeout=%d",
		conf.Clickhouse.Host, conf.Clickhouse.Port, conf.Clickhouse.Database,
		conf.Clickhouse.Username, conf.Clickhouse.Password,
		conf.Clickhouse.ReadTimeout, conf.Clickhouse.WriteTimeout)
}

//GenDialector 生成数据库连接方言(驱动)
func (conf Conf) GenDialector() gorm.Dialector {
	switch strings.ToLower(conf.DriverName) {
	case DriverNameMysql:
		return mysql.Open(mysqlDsn(conf))
	case DriverNamePostgres:
		return postgres.Open(postgresDsn(conf))
	case DriverNameSqlite:
		return sqlite.Open(sqliteDsn(conf))
	case DriverNameSqlserver:
		return sqlserver.Open(sqlserverDsn(conf))
	case DriverNameClickHouse:
		return clickhouse.Open(clickhouseDsn(conf))
	default:
		panic("not supported database driver")
	}
}
