package logger

import (
	"log"
	"os"
	"path/filepath"

	"go.uber.org/zap"
)

//InitLoggerZap 初始化zap日志服务
func InitLoggerZap(cfg Conf) {
	//定义全局日志组件配置
	zapLevel := zap.NewAtomicLevel()
	if err := zapLevel.UnmarshalText([]byte(cfg.Level)); err != nil {
		panic(err)
	}

	var zapConf zap.Config
	if env := cfg.Env; env == "dev" {
		zapConf = zap.NewDevelopmentConfig()
	} else {
		zapConf = zap.NewProductionConfig()
	}
	zapConf.Level = zapLevel
	zapConf.OutputPaths = []string{}
	zapConf.OutputPaths = append(zapConf.OutputPaths, cfg.Filename)

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
}
