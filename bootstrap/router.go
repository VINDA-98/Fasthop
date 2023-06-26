package bootstrap

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/VINDA-98/Fasthop/app/middleware"
	"github.com/VINDA-98/Fasthop/global"
	"github.com/VINDA-98/Fasthop/routes"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	if global.App.Config.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	// gin.SetMode(gin.ReleaseMode) // 生产模式下，设置该选项，将不会记录debug的日志
	router.Use(gin.Logger(), middleware.CustomRecovery())
	router.MaxMultipartMemory = 200 << 20 // 200 MiB todo 考虑到视频上传

	// 跨域处理
	router.Use(middleware.Cors())

	// 前端项目静态资源
	router.StaticFile("/", "./static/dist/index.html")
	router.StaticFile("/login", "./static/dist/auth/login.html")
	router.StaticFile("/home", "./static/dist/auth/index.html")
	router.Static("/assets", "./static/dist/assets")
	router.StaticFile("/favicon.ico", "./static/dist/favicon.ico")
	// 其他静态资源
	router.Static("/public", "./static")
	router.Static("/storage", "./storage/app/public")

	// 注册 api 分组路由
	apiGroup := router.Group("/api")
	routes.SetApiGroupRoutes(apiGroup)

	return router
}

func RunServer() {
	r := setupRouter()

	srv := &http.Server{
		Addr:    ":" + global.App.Config.App.Port,
		Handler: r,
	}
	log.Printf("Server Run: localhost%s \n", srv.Addr)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 设置优雅退出
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("正在注销服务")
	// 等待中断信号以优雅地关闭服务器（设置 3 秒的超时时间）
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("服务退出成功")
}
