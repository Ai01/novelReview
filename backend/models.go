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

type Book struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"type:varchar(255);not null" json:"title"`
	Author      string         `gorm:"type:varchar(255);not null" json:"author"`
	Cover       string         `gorm:"type:varchar(512)" json:"cover"`
	Description string         `gorm:"type:text" json:"description"`
	Category    string         `gorm:"type:varchar(100)" json:"category"`
	Tags        string         `gorm:"type:varchar(255)" json:"tags"`
	Status      string         `gorm:"type:varchar(50);default:连载中" json:"status"`
	LastUpdated *time.Time     `json:"last_updated"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

type Comment struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	BookID    uint           `gorm:"not null;index:idx_comments_book_created" json:"book_id"`
	UserID    uint           `gorm:"not null" json:"user_id"`
	Content   string         `gorm:"type:text;not null" json:"content"`
	Likes     int            `gorm:"default:0" json:"likes"`
	Book      Book           `gorm:"foreignKey:BookID" json:"book,omitempty"`
	User      User           `gorm:"foreignKey:UserID" json:"user,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
