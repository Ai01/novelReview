package book

import (
	"time"

	"gorm.io/gorm"
)

// Book 表示书籍实体，对应 GORM 模型映射到 books 表。
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
