package utils

import (
	"go.uber.org/zap"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// ListeningExitSignal 监听系统退出信号
func ListeningExitSignal(logger *zap.Logger, wg *sync.WaitGroup) {
	signals := make(chan os.Signal, 1) // 监听退出
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	sig := <-signals

	wg.Wait()
	logger.Sugar().Infof("Receive signal: %v,program exited.", zap.Any("signal", sig))
}
