package comment

import (
	"time"

	"github.com/user/novel-backend/internal/domain/book"
	"github.com/user/novel-backend/internal/domain/user"
	"gorm.io/gorm"
)

// Comment 表示书籍评论实体，对应 GORM 模型映射到 comments 表。
type Comment struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	BookID    uint           `gorm:"not null;index:idx_comments_book_created" json:"book_id"`
	UserID    uint           `gorm:"not null" json:"user_id"`
	Content   string         `gorm:"type:text;not null" json:"content"`
	Likes     int            `gorm:"default:0" json:"likes"`
	Book      book.Book      `gorm:"foreignKey:BookID" json:"book,omitempty"`
	User      user.User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
