package routes

import (
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/keepchen/app-template/pkg/app/httpserver/config"

	"github.com/gin-gonic/gin"
	"github.com/keepchen/app-template/pkg/app/httpserver/handler"
	mdlw "github.com/keepchen/app-template/pkg/app/httpserver/middleware"
)

//注册路由
func registerRoutes(r *gin.Engine, cfg *config.Config) {
	if cfg.Debug {
		//仅在调试模式下才开始pprof检测
		pprof.Register(r, "debug/pprof")
	}

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "%s", "pong")
	})
	api := r.Group("/api/v1")
	{
		api.Use(mdlw.AuthCheck()).
			GET("/say-hello", handler.SayHello)
	}
	//@see https://gin-gonic.com/docs/examples/html-rendering/
	r.LoadHTMLGlob("pkg/app/httpserver/templates/*")
	r.GET("/", handler.Index)
}
