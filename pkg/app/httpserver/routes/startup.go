package routes

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/gin-gonic/gin"
	"github.com/keepchen/app-template/pkg/app/httpserver/config"
	"github.com/keepchen/app-template/pkg/app/httpserver/docs"
	"go.uber.org/zap"
)

//RunServer 启动路由服务
func RunServer(ctx context.Context, cfg *config.Config, wg *sync.WaitGroup) {
	defer wg.Done()

	var r *gin.Engine
	if cfg.Debug {
		gin.SetMode(gin.DebugMode)
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
		r = gin.New()
	}

	//注册路由
	registerRoutes(r, cfg)

	//swagger 必须放在所有路由定义的最后
	runSwaggerServerOnDebugMode(r, cfg)

	//手动监听服务并检测退出信号从而实现优雅退出
	srv := &http.Server{
		Addr:    cfg.HttpServer.Addr,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			zap.L().Info("listen error", zap.Errors("error", []error{err}))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("Shutting down server...")

	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Error("Server forced to shutdown", zap.Errors("error", []error{err}))
	}

	zap.L().Info("Server exiting")
}

//启动swagger接口文档服务
//
//仅在调试模式下才会开启
func runSwaggerServerOnDebugMode(r *gin.Engine, cfg *config.Config) {
	if !cfg.Debug {
		//如果不是调试模式就不注册swagger路由
		return
	}
	//swagger配置
	docs.SwaggerInfo.Title = "demo Server API"
	docs.SwaggerInfo.Description = "This is a document for demo server."
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Host = cfg.HttpServer.Addr
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.Version = "1.0"

	//注册路由
	url := ginSwagger.URL("/swagger/doc.json") // The url pointing to API definition
	//access /swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
