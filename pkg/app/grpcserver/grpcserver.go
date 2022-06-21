package grpcserver

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/keepchen/app-template/pkg/common/grpc/pb"
	"google.golang.org/grpc"

	"github.com/keepchen/app-template/pkg/app/grpcserver/handler"

	"github.com/keepchen/app-template/pkg/app/grpcserver/config"
	"go.uber.org/zap"
)

//StartGrpcServer 启动grpc服务
func StartGrpcServer(logger *zap.Logger, wg *sync.WaitGroup, cfg *config.Config) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("---- Recovered ----", zap.Any("error", err))
		}
	}()
	logger.Info("Start running grpc server", zap.String("grpc addr", cfg.GrpcServer.Addr))

	//监听退出信号
	errChan := make(chan error)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%v", <-c)
		cancel()
	}()

	logger.Sugar().Infof("Start grpc server at: %s", cfg.GrpcServer.Addr)
	//启动grpc服务
	go RunGrpcServer(ctx, logger, cfg)

	//收到退出信号
	logger.Sugar().Warnf("Exit signal: %v", <-errChan)

	wg.Done()
	logger.Sugar().Warnf("Shutting down api server at %s", cfg.GrpcServer.Addr)
}

//RunGrpcServer 启动grpc服务
func RunGrpcServer(ctx context.Context, logger *zap.Logger, cfg *config.Config) {
	listener, err := net.Listen("tcp", cfg.GrpcServer.Addr)
	if err != nil {
		logger.Fatal("Start grpc server failed", zap.Errors("error", []error{err}))
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, &handler.GreeterServer{})

	go func() {
		RunClient(cfg)
	}()

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			logger.Fatal("grpc serve failed", zap.Errors("error", []error{err}))
		}
	}()

WATCHER:
	for range ctx.Done() {
		logger.Warn("Receive quit signal, shutting down grpc server")
		grpcServer.GracefulStop()
		break WATCHER
	}
	logger.Warn("grpc server shutdown finished.")
}
