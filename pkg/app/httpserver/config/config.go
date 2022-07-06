package config

import (
	"log"

	"github.com/jinzhu/configor"
	"github.com/keepchen/app-template/pkg/lib/db"
	"github.com/keepchen/app-template/pkg/lib/jwt"
	"github.com/keepchen/app-template/pkg/lib/logger"
	"github.com/keepchen/app-template/pkg/lib/redis"
)

//Config 整体的配置信息
type Config struct {
	Debug      bool           `yaml:"debug" toml:"debug" json:"debug"`                   //是否是调试模式
	Logger     logger.Conf    `yaml:"logger" toml:"logger" json:"logger"`                //日志
	DB         db.Conf        `yaml:"db" toml:"db" json:"db"`                            //数据库配置
	Redis      redis.Conf     `yaml:"redis" toml:"redis" json:"redis"`                   //redis配置
	JWT        jwt.Conf       `yaml:"jwt" toml:"jwt" json:"jwt"`                         //jwt配置
	HttpServer HttpServerConf `yaml:"http_server" toml:"http_server" json:"http_server"` //http服务配置
}

//HttpServerConf http服务配置
type HttpServerConf struct {
	Addr string `yaml:"addr" toml:"addr" json:"addr" default:":8080"` //监听地址
}

//C 全局配置变量
var C = &Config{}

//ParseConfig 解析配置
func ParseConfig(cfgPath string) {
	if len(cfgPath) != 0 {
		if err := configor.New(&configor.Config{AutoReload: true}).Load(C, cfgPath); err != nil {
			panic(err)
		}
	} else {
		if err := configor.New(&configor.Config{AutoReload: true}).Load(C); err != nil {
			panic(err)
		}
	}

	//解析jwt配置
	C.JWT.Load()

	if C.Debug {
		log.Printf("loaded config: %#v", C)
	}
}
