package service

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/user/novel-backend/internal/domain/user"
	"github.com/user/novel-backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

// Claims 代表 JWT 载荷。
type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateToken 为用户签发 JWT（24 小时有效期）。
func GenerateToken(userID uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// ValidateToken 校验 JWT 并返回 Claims。
func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}

// LoginResult 登录返回。
type LoginResult struct {
	Token string    `json:"token"`
	User  user.User `json:"user"`
}

// AuthService 认证业务服务。
type AuthService struct {
	userRepo repository.UserRepository
}

// NewAuthService 创建 AuthService。
func NewAuthService(userRepo repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

// Login 验证用户凭据并返回 JWT 和用户信息。
func (s *AuthService) Login(username, password string) (*LoginResult, error) {
	u, err := s.userRepo.FindByCredentials(username)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return nil, errors.New("incorrect password")
	}

	token, err := GenerateToken(u.ID)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &LoginResult{Token: token, User: *u}, nil
}
