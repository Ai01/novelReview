package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginHandler(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	strongRead := c.GetHeader("X-Read-Consistency") == "strong"
	db := DBForRead(strongRead)
	if db == nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Database not ready"})
		return
	}

	var user User
	if err := db.Where("username = ? OR email = ?", req.Username, req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
		return
	}

	token, err := GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  user,
	})
}

// ExploreItem 探索页信息流单条响应
type ExploreItem struct {
	CommentID uint   `json:"comment_id"`
	Book      Book   `json:"book"`
	User      User   `json:"user"`
	Content   string `json:"content"`
	Likes     int    `json:"likes"`
	CreatedAt string `json:"created_at"`
}

// ExploreHandler 探索页评论信息流（游标分页）
func ExploreHandler(c *gin.Context) {
	cursorStr := c.DefaultQuery("cursor", "0")
	limitStr := c.DefaultQuery("limit", "20")

	cursor, err := strconv.Atoi(cursorStr)
	if err != nil || cursor < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的游标参数"})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 20
	}
	if limit > 50 {
		limit = 20
	}

	db := DBForRead(c.GetHeader("X-Read-Consistency") == "strong")
	if db == nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "服务暂时不可用"})
		return
	}

	var comments []Comment
	query := db.Preload("Book").Preload("User")
	if cursor > 0 {
		query = query.Where("id < ?", cursor)
	}

	if err := query.Order("id DESC").Limit(limit + 1).Find(&comments).Error; err != nil {
		log.Printf("ExploreHandler query error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "服务器内部错误"})
		return
	}

	hasMore := len(comments) > limit
	if hasMore {
		comments = comments[:limit]
	}

	items := make([]ExploreItem, len(comments))
	for i, cmt := range comments {
		items[i] = ExploreItem{
			CommentID: cmt.ID,
			Book:      cmt.Book,
			User:      cmt.User,
			Content:   cmt.Content,
			Likes:     cmt.Likes,
			CreatedAt: cmt.CreatedAt.Format("2006-01-02T15:04:05Z"),
		}
	}

	nextCursor := uint(0)
	if len(items) > 0 {
		nextCursor = items[len(items)-1].CommentID
	}

	c.JSON(http.StatusOK, gin.H{
		"items":       items,
		"next_cursor": nextCursor,
		"has_more":    hasMore,
	})
}

// ExploreSearchHandler 探索页搜索
func ExploreSearchHandler(c *gin.Context) {
	q := c.Query("q")
	if q == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "搜索关键词不能为空"})
		return
	}

	cursorStr := c.DefaultQuery("cursor", "0")
	limitStr := c.DefaultQuery("limit", "20")

	cursor, err := strconv.Atoi(cursorStr)
	if err != nil || cursor < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的游标参数"})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 20
	}
	if limit > 50 {
		limit = 20
	}

	db := DBForRead(c.GetHeader("X-Read-Consistency") == "strong")
	if db == nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "服务暂时不可用"})
		return
	}

	var comments []Comment
	likePattern := "%" + q + "%"

	query := db.Preload("Book").Preload("User").
		Joins("JOIN books ON books.id = comments.book_id").
		Where("books.title LIKE ? OR books.author LIKE ? OR comments.content LIKE ?", likePattern, likePattern, likePattern)

	if cursor > 0 {
		query = query.Where("comments.id < ?", cursor)
	}

	if err := query.Order("comments.id DESC").Limit(limit + 1).Find(&comments).Error; err != nil {
		log.Printf("ExploreSearchHandler query error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "服务器内部错误"})
		return
	}

	hasMore := len(comments) > limit
	if hasMore {
		comments = comments[:limit]
	}

	items := make([]ExploreItem, len(comments))
	for i, cmt := range comments {
		items[i] = ExploreItem{
			CommentID: cmt.ID,
			Book:      cmt.Book,
			User:      cmt.User,
			Content:   cmt.Content,
			Likes:     cmt.Likes,
			CreatedAt: cmt.CreatedAt.Format("2006-01-02T15:04:05Z"),
		}
	}

	nextCursor := uint(0)
	if len(items) > 0 {
		nextCursor = items[len(items)-1].CommentID
	}

	c.JSON(http.StatusOK, gin.H{
		"items":       items,
		"next_cursor": nextCursor,
		"has_more":    hasMore,
	})
}
