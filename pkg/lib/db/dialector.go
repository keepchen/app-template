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

//DsnFields 连接字符串字段
type DsnFields struct {
	DriverName   string //数据库类型名称
	Username     string //用户名
	Password     string //密码
	Host         string //主机地址
	Port         int    //端口
	Database     string //数据库名
	Charset      string //字符集
	ParseTime    bool   //是否解析时间
	Loc          string //位置
	SSLMode      string //ssl模式 enable|disable
	TimeZone     string //时区
	SqliteFile   string //sqlite文件
	ReadTimeout  int    //读取超时时间
	WriteTimeout int    //写入超时时间
}

//组装mysql的dsn
//@see https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
func mysqlDsn(fields DsnFields) string {
	if len(fields.Charset) == 0 {
		fields.Charset = "utf8mb4"
	}
	if len(fields.Loc) == 0 {
		fields.Loc = "Local"
	}
	var parseTime string
	if fields.ParseTime {
		parseTime = "True"
	} else {
		parseTime = "False"
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%d)%s?charset=%s&parseTime=%s&loc=%s",
		fields.Username, fields.Password, fields.Host, fields.Port, fields.Database, fields.Charset, parseTime, fields.Loc)
}

//组装postgres的dsn
//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
func postgresDsn(fields DsnFields) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		fields.Host, fields.Username, fields.Password, fields.Database, fields.Port, fields.SSLMode, fields.TimeZone)
}

//组装sqlite的dsn
//dsn := "sqlite.db"
func sqliteDsn(fields DsnFields) string {
	return fields.SqliteFile
}

//组装sqlserver的dsn
//dsn := "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
func sqlserverDsn(fields DsnFields) string {
	return fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
		fields.Username, fields.Password, fields.Host, fields.Port, fields.Database)
}

//组装clickhouse的dsn
//dsn := "tcp://localhost:9000?database=gorm&username=gorm&password=gorm&read_timeout=10&write_timeout=20"
func clickhouseDsn(fields DsnFields) string {
	return fmt.Sprintf("tcp://%s:%d?database=%s&username=%s&password=%s&read_timeout=%d&write_timeout=%d",
		fields.Host, fields.Port, fields.Database, fields.Username, fields.Password, fields.ReadTimeout, fields.WriteTimeout)
}

//GenDialector 生成数据库连接方言(驱动)
func (fields DsnFields) GenDialector() gorm.Dialector {
	switch strings.ToLower(fields.DriverName) {
	case DriverNameMysql:
		return mysql.Open(mysqlDsn(fields))
	case DriverNamePostgres:
		return postgres.Open(postgresDsn(fields))
	case DriverNameSqlite:
		return sqlite.Open(sqliteDsn(fields))
	case DriverNameSqlserver:
		return sqlserver.Open(sqlserverDsn(fields))
	case DriverNameClickHouse:
		return clickhouse.Open(clickhouseDsn(fields))
	default:
		panic("not supported database driver")
	}
}
