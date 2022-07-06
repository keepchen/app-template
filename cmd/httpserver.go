package cmd

import (
	"sync"

	"github.com/keepchen/app-template/pkg/lib/logger"

	"github.com/keepchen/app-template/pkg/app/httpserver"
	"github.com/keepchen/app-template/pkg/app/httpserver/config"
	"github.com/keepchen/app-template/pkg/utils"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func httpServerCMD() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "httpserver",
		Short: "启动http服务",
		Run: func(cmd *cobra.Command, args []string) {
			//启动时要执行的操作写在这里
			logSvc := logger.GetLogger()
			loggerSvc := logSvc.With(zap.String("serve", "[httpserver]"))
			wg := &sync.WaitGroup{}

			//启动http接口服务
			wg.Add(1)
			go httpserver.StartHttpServer(loggerSvc, wg, config.C)

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
	RootCMD.AddCommand(httpServerCMD())
}
