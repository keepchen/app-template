package logger

//Conf 日志配置
type Conf struct {
	Env        string `yaml:"env" toml:"env" json:"env" default:"prod"`                            //日志环境，prod：生产环境，dev：开发环境
	Level      string `yaml:"level" toml:"level" json:"level" default:"info"`                      //日志级别，debug，info，warn，error
	Filename   string `yaml:"filename" toml:"filename" json:"filename" default:"logs/running.log"` //日志文件名称
	MaxSize    int    `yaml:"max_size" toml:"max_size" json:"max_size" default:"100"`              //日志大小限制，单位MB
	MaxBackups int    `yaml:"max_backups" toml:"max_backups" json:"max_backups" default:"10"`      //最大历史文件保留数量
	Compress   bool   `yaml:"compress" toml:"compress" json:"compress" default:"true"`             //是否压缩历史日志文件
}
