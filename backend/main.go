package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/user/novel-backend/internal/handler"
	"github.com/user/novel-backend/internal/infra"
	"github.com/user/novel-backend/internal/repository"
	"github.com/user/novel-backend/internal/router"
	"github.com/user/novel-backend/internal/service"
)

func main() {
	// 1. 初始化基础设施（数据库连接）
	if err := infra.InitDB(); err != nil {
		log.Printf("Database init failed (service will start not-ready): %v", err)
	}

	// 2. 组装依赖链
	readDB := infra.DBForRead(false)

	// Repository 层
	userRepo := repository.NewUserRepository(readDB)
	commentRepo := repository.NewCommentRepository(readDB)

	// Service 层
	authService := service.NewAuthService(userRepo)
	exploreService := service.NewExploreService(commentRepo)

	// Handler 层
	authHandler := handler.NewAuthHandler(authService)
	exploreHandler := handler.NewExploreHandler(exploreService)

	// 3. 启动 Gin 引擎
	r := gin.Default()

	// 健康检查（不走 router，直接注册）
	r.GET("/healthz", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	r.GET("/readyz", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 1*time.Second)
		defer cancel()
		if err := infra.PingWrite(ctx); err != nil {
			c.String(http.StatusServiceUnavailable, "not ready")
			return
		}
		c.String(http.StatusOK, "ready")
	})
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://127.0.0.1:5173", "http://localhost", "http://127.0.0.1"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// 注册业务路由
	router.Setup(&r.RouterGroup, authHandler, exploreHandler)

	r.Run(":8080")
}
