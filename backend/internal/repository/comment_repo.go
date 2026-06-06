package repository

import (
	"github.com/user/novel-backend/internal/domain/book"
	"github.com/user/novel-backend/internal/domain/comment"
	"gorm.io/gorm"
)

// CommentRepository 定义评论数据访问接口。
type CommentRepository interface {
	FindExploreList(cursor string, limit int) ([]comment.Comment, bool, error)
	FindByKeyword(q string, cursor string, limit int) ([]comment.Comment, bool, error)
}

type commentRepo struct {
	db *gorm.DB
}

// NewCommentRepository 创建 CommentRepository 实现。
func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepo{db: db}
}

func (r *commentRepo) FindExploreList(cursor string, limit int) ([]comment.Comment, bool, error) {
	var comments []comment.Comment
	query := r.db.Preload("Book").Preload("User")
	if cursor != "0" {
		query = query.Where("created_at < ?", cursor)
	}

	if err := query.Order("created_at DESC").Limit(limit + 1).Find(&comments).Error; err != nil {
		return nil, false, err
	}

	hasMore := len(comments) > limit
	if hasMore {
		comments = comments[:limit]
	}
	return comments, hasMore, nil
}

func (r *commentRepo) FindByKeyword(q string, cursor string, limit int) ([]comment.Comment, bool, error) {
	var comments []comment.Comment
	likePattern := "%" + q + "%"

	bookSubQuery := r.db.Model(&book.Book{}).Select("id").
		Where("title LIKE ? OR author LIKE ?", likePattern, likePattern)

	query := r.db.Preload("Book").Preload("User").
		Where(r.db.Where("book_id IN (?)", bookSubQuery).Or("content LIKE ?", likePattern))

	if cursor != "0" {
		query = query.Where("created_at < ?", cursor)
	}

	if err := query.Order("created_at DESC").Limit(limit + 1).Find(&comments).Error; err != nil {
		return nil, false, err
	}

	hasMore := len(comments) > limit
	if hasMore {
		comments = comments[:limit]
	}
	return comments, hasMore, nil
}
