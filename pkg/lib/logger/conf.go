package logger

//Conf 日志配置
type Conf struct {
	Env      string `yaml:"env" toml:"env" json:"env" default:"prod"`                           //日志环境，prod：生产环境，dev：开发环境
	Level    string `yaml:"level" toml:"level" json:"level" default:"info"`                     //日志级别，debug，info，warning，error
	Filename string `yam:"filename" toml:"filename" json:"filename" default:"logs/running.log"` //日志文件名称
}
