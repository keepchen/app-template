package httpserver

import (
	"context"
	"fmt"
	"github.com/keepchen/app-template/pkg/app/httpserver/config"
	"github.com/keepchen/app-template/pkg/app/httpserver/routes"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

//StartApiServer 启动api服务
func StartApiServer(logger *zap.Logger, wg *sync.WaitGroup, cfg *config.Config) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("---- Recovered ----", zap.Any("error", err))
		}
	}()
	logger.Info("Start running http server", zap.String("http addr", cfg.HttpServer.Addr))

	//监听退出信号
	errChan := make(chan error)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%v", <-c)
		cancel()
	}()

	//启动http路由服务
	logger.Sugar().Infof("Start http server at: %s", cfg.HttpServer.Addr)
	wg.Add(1)
	go routes.RunServer(ctx, cfg, wg)

	logger.Sugar().Infof("Start grpc server at: %s", cfg.HttpServer.Addr)

	//收到退出信号
	logger.Sugar().Warnf("Exit signal: %v", <-errChan)

	wg.Done()
	logger.Sugar().Warnf("Shutting down api server at %s", cfg.HttpServer.Addr)
}
