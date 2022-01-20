package config

import (
	"github.com/jinzhu/configor"
	"github.com/keepchen/app-template/pkg/utils"
	"go.uber.org/zap"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Debug      bool           `yaml:"debug" toml:"debug" json:"debug"`                   //是否是调试模式
	Secret     SecretConf     `yaml:"secret" toml:"secret" json:"secret"`                //密钥配置
	Logger     LoggerConf     `yaml:"logger" toml:"logger" json:"logger"`                //日志
	HttpServer HttpServerConf `yaml:"http_server" toml:"http_server" json:"http_server"` //http服务配置
}

type SecretConf struct {
	PublicKeyFilename  string `yaml:"public_key_filename" toml:"public_key_filename" json:"public_key_filename"`    //公钥文件地址
	PrivateKeyFilename string `yaml:"private_key_filename" toml:"private_key_filename" json:"private_key_filename"` //私钥文件地址
	PublicKeyBytes     []byte `yaml:"-" toml:"-" json:"-"`
	PrivateKeyBytes    []byte `yaml:"-" toml:"-" json:"-"`
}

type LoggerConf struct {
	Env      string `yaml:"env" toml:"env" default:"prod" json:"env"`                           //日志环境，prod：生产环境，dev：开发环境
	Level    string `yaml:"level" toml:"level" default:"info" json:"info"`                      //日志级别，debug，info，warning，error
	Filename string `yam:"filename" toml:"filename" default:"logs/running.log" json:"filename"` //日志文件名称
}

type HttpServerConf struct {
	Addr string `yaml:"addr" toml:"addr" json:"addr" default:":8080"` //监听地址
}

var C = &Config{}

//ParseConfig 解析配置
func ParseConfig(cfgPath string) {
	if len(cfgPath) != 0 {
		if err := configor.New(&configor.Config{AutoReload: true}).Load(C, cfgPath); err != nil {
			zap.L().Panic("init config fail", zap.Error(err))
		}
	} else {
		if err := configor.New(&configor.Config{AutoReload: true}).Load(C); err != nil {
			zap.L().Panic("init config fail", zap.Error(err))
		}
	}

	//载入密钥
	pub, err := utils.FileGetContents(C.Secret.PublicKeyFilename)
	if err != nil {
		zap.L().Panic("read rsa public key fail", zap.Error(err))
	}
	C.Secret.PublicKeyBytes = pub

	pri, err := utils.FileGetContents(C.Secret.PrivateKeyFilename)
	if err != nil {
		zap.L().Panic("read rsa private key fail", zap.Error(err))
	}
	C.Secret.PrivateKeyBytes = pri

	//定义全局日志组件配置
	zapLevel := zap.NewAtomicLevel()
	if err := zapLevel.UnmarshalText([]byte(C.Logger.Level)); err != nil {
		panic(err)
	}

	var zapConf zap.Config
	if env := C.Logger.Env; env == "dev" {
		zapConf = zap.NewDevelopmentConfig()
	} else {
		zapConf = zap.NewProductionConfig()
	}
	zapConf.Level = zapLevel
	zapConf.OutputPaths = []string{}
	zapConf.OutputPaths = append(zapConf.OutputPaths, C.Logger.Filename)

	logger, err := zapConf.Build(zap.Fields(zap.String("proc", filepath.Base(os.Args[0]))))
	if err != nil {
		panic(err)
	} else {
		zap.RedirectStdLog(logger)
		zap.ReplaceGlobals(logger)
	}
	defer func() {
		err := logger.Sync()
		if err != nil {
			log.Printf("zap sync error: %v", err)
		}
	}()

	zap.L().Named("configor").Info("loaded config", zap.Any("config", C))
}
