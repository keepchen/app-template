package routes

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/keepchen/app-template/pkg/app/httpserver/config"
	"github.com/keepchen/app-template/pkg/app/httpserver/docs"
	"github.com/keepchen/app-template/pkg/app/httpserver/handler"
	mdlw "github.com/keepchen/app-template/pkg/app/httpserver/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gin-contrib/pprof"
)

//RunServer 启动路由服务
func RunServer(ctx context.Context, cfg *config.Config, wg *sync.WaitGroup) {
	defer wg.Done()
	//swagger配置
	docs.SwaggerInfo.Title = "NFT metadata Server API"
	docs.SwaggerInfo.Description = "This is a document for NFT server."
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Host = cfg.HttpServer.Addr
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.Version = "1.0"

	var r *gin.Engine
	if cfg.Debug {
		gin.SetMode(gin.DebugMode)
		r = gin.Default()
		pprof.Register(r, "debug/pprof")
	} else {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
		r = gin.New()
	}

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "%s", "pong")
	})
	api := r.Group("/api/v1")
	{
		api.Use(mdlw.AuthCheck()).
			GET("/say-hello", handler.SayHello)
	}

	//swagger 必须放在所有路由定义的最后
	if cfg.Debug {
		url := ginSwagger.URL("/swagger/doc.json") // The url pointing to API definition
		//access /swagger/index.html
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}

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
