package repository

import (
	"github.com/user/novel-backend/internal/domain/user"
	"gorm.io/gorm"
)

// UserRepository 定义用户数据访问接口。
type UserRepository interface {
	FindByCredentials(login string) (*user.User, error)
}

type userRepo struct {
	db *gorm.DB
}

// NewUserRepository 创建 UserRepository 实现。
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) FindByCredentials(login string) (*user.User, error) {
	var u user.User
	if err := r.db.Where("username = ? OR email = ?", login, login).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}
