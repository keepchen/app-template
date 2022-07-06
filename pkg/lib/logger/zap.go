package logger

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
)

var loggerSvc *zap.Logger

//GetLogger 获取日志服务实例
func GetLogger() *zap.Logger {
	return loggerSvc
}

//InitLoggerZap 初始化zap日志服务
func InitLoggerZap(cfg Conf) {
	//定义全局日志组件配置
	atomicLevel := zap.NewAtomicLevel()
	switch strings.ToLower(cfg.Level) {
	case "debug":
		atomicLevel.SetLevel(zapcore.DebugLevel)
	case "info":
		atomicLevel.SetLevel(zapcore.InfoLevel)
	case "warn":
		atomicLevel.SetLevel(zapcore.WarnLevel)
	case "error":
		atomicLevel.SetLevel(zapcore.ErrorLevel)
	case "dpanic":
		atomicLevel.SetLevel(zapcore.DPanicLevel)
	case "panic":
		atomicLevel.SetLevel(zapcore.PanicLevel)
	case "fatal":
		atomicLevel.SetLevel(zapcore.FatalLevel)
	}
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "name",
		CallerKey:      "line",
		MessageKey:     "msg",
		FunctionKey:    "func",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	writer := &lumberjack.Logger{
		Filename:   cfg.Filename,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		LocalTime:  true,
		Compress:   cfg.Compress,
	}

	zapCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(writer), atomicLevel,
	)

	loggerSvc = zap.New(zapCore, zap.AddCaller())
	defer func() {
		_ = loggerSvc.Sync()
	}()
}

//InitLoggerZapOld 初始化zap日志服务
//
//DEPRECATED
func InitLoggerZapOld(cfg Conf) {
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
