package user

import (
	"time"

	"gorm.io/gorm"
)

// User 表示应用用户实体，对应 GORM 模型映射到 users 表。
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"type:varchar(191);uniqueIndex;not null" json:"username"`
	Email     string         `gorm:"type:varchar(191);uniqueIndex;not null" json:"email"`
	Password  string         `gorm:"not null" json:"-"`
	Avatar    string         `json:"avatar"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
