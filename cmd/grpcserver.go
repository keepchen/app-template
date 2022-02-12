package cmd

import (
	"sync"

	"github.com/keepchen/app-template/pkg/lib/logger"

	"github.com/keepchen/app-template/pkg/utils"

	"github.com/keepchen/app-template/pkg/app/grpcserver"

	"github.com/keepchen/app-template/pkg/app/grpcserver/config"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func grpcServerCMD() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "grpcserver",
		Short: "启动grpc服务",
		Run: func(cmd *cobra.Command, args []string) {
			//启动时要执行的操作写在这里
			loggerSvc := zap.L().With(zap.String("serve", "[grpcserver]"))
			wg := &sync.WaitGroup{}

			//启动grpc接口服务
			wg.Add(1)
			go grpcserver.StartGrpcServer(loggerSvc, wg, config.C)

			//更多服务...
			utils.ListeningExitSignal(loggerSvc, wg)
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			//启动前要执行的方法写在这里，例如加载配置文件、初始化数据库连接等
			config.ParseConfig(cfgPath)
			logger.InitLoggerZap(config.C.Logger)
		},
	}

	cmd.PersistentFlags().StringVarP(&cfgPath, "config", "c", "", "配置文件路径")
	return cmd
}

func init() {
	RootCMD.AddCommand(grpcServerCMD())
}
