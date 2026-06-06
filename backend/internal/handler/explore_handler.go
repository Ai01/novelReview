package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/user/novel-backend/internal/service"
)

// ExploreHandler 处理探索页相关 HTTP 请求。
type ExploreHandler struct {
	exploreService *service.ExploreService
}

// NewExploreHandler 创建 ExploreHandler。
func NewExploreHandler(exploreService *service.ExploreService) *ExploreHandler {
	return &ExploreHandler{exploreService: exploreService}
}

// Feed GET /api/v1/explore
func (h *ExploreHandler) Feed(c *gin.Context) {
	cursorStr := c.DefaultQuery("cursor", "0")
	limitStr := c.DefaultQuery("limit", "20")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 20
	}
	if limit > 50 {
		limit = 20
	}

	result, err := h.exploreService.GetFeed(cursorStr, limit)
	if err != nil {
		log.Printf("ExploreHandler.Feed error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "服务器内部错误"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items":       result.Items,
		"next_cursor": result.NextCursor,
		"has_more":    result.HasMore,
	})
}

// Search GET /api/v1/explore/search
func (h *ExploreHandler) Search(c *gin.Context) {
	q := c.Query("q")
	if q == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "搜索关键词不能为空"})
		return
	}

	cursorStr := c.DefaultQuery("cursor", "0")
	limitStr := c.DefaultQuery("limit", "20")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 20
	}
	if limit > 50 {
		limit = 20
	}

	result, err := h.exploreService.Search(q, cursorStr, limit)
	if err != nil {
		log.Printf("ExploreHandler.Search error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "服务器内部错误"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items":       result.Items,
		"next_cursor": result.NextCursor,
		"has_more":    result.HasMore,
	})
}
