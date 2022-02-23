package db

//Conf 配置信息
type Conf struct {
	DriverName string         `yaml:"diver_name" toml:"diver_name" json:"diver_name" default:"mysql"` //数据库类型
	Mysql      MysqlConf      `yaml:"mysql" toml:"mysql" json:"mysql"`                                //mysql配置
	Postgres   PostgresConf   `yaml:"postgres" toml:"postgres" json:"postgres"`                       //postgres配置
	Sqlserver  SqlserverConf  `yaml:"sqlserver" toml:"sqlserver" json:"sqlserver"`                    //sqlserver配置
	Sqlite     SqliteConf     `yaml:"sqlite" toml:"sqlite" json:"sqlite"`                             //sqlite配置
	Clickhouse ClickhouseConf `yaml:"clickhouse" toml:"clickhouse" json:"clickhouse"`                 //clickhouse配置
}

//MysqlConf mysql配置
type MysqlConf struct {
	Host      string `yaml:"host" toml:"host" json:"host" default:"localhost"`           //主机地址
	Port      int    `yaml:"port" toml:"port" json:"port" default:"3306"`                //端口
	Username  string `yaml:"username" toml:"username" json:"username"`                   //用户名
	Password  string `yaml:"password" toml:"password" json:"password"`                   //密码
	Database  string `yaml:"database" toml:"database" json:"database"`                   //数据库名
	Charset   string `yaml:"charset" toml:"charset" json:"charset"`                      //字符集
	ParseTime bool   `yaml:"parseTime" toml:"parseTime" json:"parseTime" default:"true"` //是否解析时间
	Loc       string `yaml:"loc" toml:"loc" json:"loc" default:"Local"`                  //位置
}

//PostgresConf postgres配置
type PostgresConf struct {
	Host     string `yaml:"host" toml:"host" json:"host" default:"localhost"`                 //主机地址
	Port     int    `yaml:"port" toml:"port" json:"port" default:"9920"`                      //端口
	Username string `yaml:"username" toml:"username" json:"username"`                         //用户名
	Password string `yaml:"password" toml:"password" json:"password"`                         //密码
	Database string `yaml:"database" toml:"database" json:"database"`                         //数据库名
	SSLMode  string `yaml:"ssl_mode" toml:"ssl_mode" json:"ssl_mode"`                         //ssl模式 enable|disable
	TimeZone string `yaml:"timezone" toml:"timezone" json:"timezone" default:"Asia/Shanghai"` //时区
}

//SqlserverConf sqlserver配置
type SqlserverConf struct {
	Host     string `yaml:"host" toml:"host" json:"host" default:"localhost"` //主机地址
	Port     int    `yaml:"port" toml:"port" json:"port" default:"9930"`      //端口
	Username string `yaml:"username" toml:"username" json:"username"`         //用户名
	Password string `yaml:"password" toml:"password" json:"password"`         //密码
	Database string `yaml:"database" toml:"database" json:"database"`         //数据库名
}

//SqliteConf sqlite配置
type SqliteConf struct {
	File string `yaml:"file" toml:"file" json:"file" default:"sqlite.db"` //数据库文件
}

//ClickhouseConf clickhouse配置
type ClickhouseConf struct {
	Host         string `yaml:"host" toml:"host" json:"host" default:"localhost"`                     //主机地址
	Port         int    `yaml:"port" toml:"port" json:"port" default:"9000"`                          //端口
	Username     string `yaml:"username" toml:"username" json:"username"`                             //用户名
	Password     string `yaml:"password" toml:"password" json:"password"`                             //密码
	Database     string `yaml:"database" toml:"database" json:"database"`                             //数据库名
	ReadTimeout  int    `yaml:"read_timeout" toml:"read_timeout" json:"read_timeout" default:"20"`    //读取超时时间
	WriteTimeout int    `yaml:"write_timeout" toml:"write_timeout" json:"write_timeout" default:"20"` //写入超时时间
}
