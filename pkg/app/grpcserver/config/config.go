package config

import (
	"github.com/jinzhu/configor"
	"github.com/keepchen/app-template/pkg/lib/db"
	"github.com/keepchen/app-template/pkg/lib/logger"
	"github.com/keepchen/app-template/pkg/lib/redis"
	"github.com/keepchen/app-template/pkg/utils"
	"go.uber.org/zap"
)

type Config struct {
	Debug      bool           `yaml:"debug" toml:"debug" json:"debug"`                   //是否是调试模式
	Secret     SecretConf     `yaml:"secret" toml:"secret" json:"secret"`                //密钥配置
	Logger     logger.Conf    `yaml:"logger" toml:"logger" json:"logger"`                //日志
	GrpcServer GrpcServerConf `yaml:"grpc_server" toml:"grpc_server" json:"grpc_server"` //grpc服务配置
	DB         db.Conf        `yaml:"db" toml:"db" json:"db"`                            //数据库配置
	Redis      redis.Conf     `yaml:"redis" toml:"redis" json:"redis"`                   //redis配置
}

type SecretConf struct {
	PublicKeyFilename  string `yaml:"public_key_filename" toml:"public_key_filename" json:"public_key_filename"`    //公钥文件地址
	PrivateKeyFilename string `yaml:"private_key_filename" toml:"private_key_filename" json:"private_key_filename"` //私钥文件地址
	PublicKeyBytes     []byte `yaml:"-" toml:"-" json:"-"`
	PrivateKeyBytes    []byte `yaml:"-" toml:"-" json:"-"`
}

type GrpcServerConf struct {
	Addr string `yaml:"addr" toml:"addr" json:"addr" default:":8081"` //监听地址
}

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

	//载入密钥
	pub, err := utils.FileGetContents(C.Secret.PublicKeyFilename)
	if err != nil {
		panic(err)
	}
	C.Secret.PublicKeyBytes = pub

	pri, err := utils.FileGetContents(C.Secret.PrivateKeyFilename)
	if err != nil {
		panic(err)
	}
	C.Secret.PrivateKeyBytes = pri

	zap.L().Named("configor").Info("loaded config", zap.Any("config", C))
}
