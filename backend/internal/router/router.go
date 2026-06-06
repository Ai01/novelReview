package router

import (
	"github.com/gin-gonic/gin"
	"github.com/user/novel-backend/internal/handler"
)

// Setup 注册所有 API 路由。
func Setup(r *gin.RouterGroup, authH *handler.AuthHandler, exploreH *handler.ExploreHandler) {
	api := r.Group("/api/v1")
	{
		api.POST("/login", authH.Login)
		api.GET("/explore", exploreH.Feed)
		api.GET("/explore/search", exploreH.Search)
	}
}
