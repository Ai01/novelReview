package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	if err := InitDB(); err != nil {
		log.Printf("Database init failed (service will start not-ready): %v", err)
	}

	r := gin.Default()

	r.GET("/healthz", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	r.GET("/readyz", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 1*time.Second)
		defer cancel()
		if err := PingWrite(ctx); err != nil {
			c.String(http.StatusServiceUnavailable, "not ready")
			return
		}
		c.String(http.StatusOK, "ready")
	})

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// 使用官方 CORS 中间件
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://127.0.0.1:5173", "http://localhost", "http://127.0.0.1"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	api := r.Group("/api/v1")
	{
		api.POST("/login", LoginHandler)
		api.GET("/explore", ExploreHandler)
		api.GET("/explore/search", ExploreSearchHandler)
	}

	r.Run(":8080")
}
