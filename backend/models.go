package main

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"type:varchar(191);uniqueIndex;not null" json:"username"`
	Email     string         `gorm:"type:varchar(191);uniqueIndex;not null" json:"email"`
	Password  string         `gorm:"not null" json:"-"` // 密码不返回给前端
	Avatar    string         `json:"avatar"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func Migrate() {
	err := DB.AutoMigrate(&User{})
	if err != nil {
		panic("Failed to migrate database: " + err.Error())
	}
}
