package redis

//Conf 配置信息
type Conf struct {
	Host     string `yaml:"host" toml:"host" json:"host" default:"localhost"` //主机地址
	Port     int    `yaml:"port" toml:"port" json:"port" default:"6379"`      //端口
	Password string `yaml:"password" toml:"password" json:"password"`         //密码
	Database int    `yaml:"database" toml:"database" json:"database"`         //数据库名
}
