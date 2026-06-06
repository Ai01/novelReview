package service

import (
	"github.com/user/novel-backend/internal/domain/comment"
	"github.com/user/novel-backend/internal/repository"
)

// ExploreItem 探索页信息流单条响应。
type ExploreItem struct {
	CommentID uint        `json:"comment_id"`
	Book      interface{} `json:"book"`
	User      interface{} `json:"user"`
	Content   string      `json:"content"`
	Likes     int         `json:"likes"`
	CreatedAt string      `json:"created_at"`
}

// ExploreResult 探索页查询结果。
type ExploreResult struct {
	Items      []ExploreItem `json:"items"`
	NextCursor string        `json:"next_cursor"`
	HasMore    bool          `json:"has_more"`
}

// ExploreService 探索业务服务。
type ExploreService struct {
	commentRepo repository.CommentRepository
}

// NewExploreService 创建 ExploreService。
func NewExploreService(commentRepo repository.CommentRepository) *ExploreService {
	return &ExploreService{commentRepo: commentRepo}
}

// GetFeed 获取探索页评论信息流。
func (s *ExploreService) GetFeed(cursor string, limit int) (*ExploreResult, error) {
	comments, hasMore, err := s.commentRepo.FindExploreList(cursor, limit)
	if err != nil {
		return nil, err
	}
	return s.buildResult(comments, hasMore), nil
}

// Search 搜索书籍评论。
func (s *ExploreService) Search(q string, cursor string, limit int) (*ExploreResult, error) {
	comments, hasMore, err := s.commentRepo.FindByKeyword(q, cursor, limit)
	if err != nil {
		return nil, err
	}
	return s.buildResult(comments, hasMore), nil
}

func (s *ExploreService) buildResult(comments []comment.Comment, hasMore bool) *ExploreResult {
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

	nextCursor := "0"
	if len(items) > 0 {
		nextCursor = items[len(items)-1].CreatedAt
	}

	return &ExploreResult{
		Items:      items,
		NextCursor: nextCursor,
		HasMore:    hasMore,
	}
}
