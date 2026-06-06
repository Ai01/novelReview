package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/user/novel-backend/internal/service"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// AuthHandler 处理认证相关 HTTP 请求。
type AuthHandler struct {
	authService *service.AuthService
}

// NewAuthHandler 创建 AuthHandler。
func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Login POST /api/v1/login
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	result, err := h.authService.Login(req.Username, req.Password)
	if err != nil {
		switch err.Error() {
		case "user not found", "incorrect password":
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": result.Token,
		"user":  result.User,
	})
}
