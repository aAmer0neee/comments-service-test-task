package repository

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/aAmer0neee/comments-service-test-task/internal/domain"
	"github.com/google/uuid"
)

type MemoryRepository struct {
	mu sync.RWMutex

	articles          map[uuid.UUID]domain.Article // ключ - идентификатор статьи
	createdAtArticles []articleWithDate

	comments          map[uuid.UUID]domain.Comment //ключ - идентификатор комментария

	commentReply map[uuid.UUID][]*domain.Comment
}

type articleWithDate struct {
	id        uuid.UUID
	createdAt time.Time
}

type commentWithDate struct {
	id        uuid.UUID
	createdAt time.Time
}

func initInMemory() (*MemoryRepository, error) {
	return &MemoryRepository{
		mu:                sync.RWMutex{},
		articles:          make(map[uuid.UUID]domain.Article),
		createdAtArticles: make([]articleWithDate, 0),
		comments:          make(map[uuid.UUID]domain.Comment),
		commentReply: map[uuid.UUID][]*domain.Comment{},
	}, nil
}

func (r *MemoryRepository) CreateArticle(article domain.Article) (domain.Article, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.articles[article.ID]; !exists {
		r.articles[article.ID] = article

		newRecord := articleWithDate{
			id:        article.ID,
			createdAt: article.CreatedAt,
		}

		recordIndex := sort.Search(len(r.createdAtArticles), func(i int) bool {
			return r.createdAtArticles[i].createdAt.Before(newRecord.createdAt)
		})

		r.createdAtArticles = append(r.createdAtArticles, articleWithDate{})

		copy(r.createdAtArticles[recordIndex+1:],
			r.createdAtArticles[recordIndex:])

		r.createdAtArticles[recordIndex] = newRecord
		return article, nil
	}
	return domain.Article{}, fmt.Errorf("try to add a non-unique value %+v", article.ID)
}

func (r *MemoryRepository) GetArticle(id uuid.UUID) (domain.Article, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if article, exists := r.articles[id]; exists {
		return article, nil

	}
	return domain.Article{}, fmt.Errorf("article id: %v not exists", id)
}

func (r *MemoryRepository) GetListArticles(page, limit int) ([]domain.Article, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if limit <= 0 || page <= 0 {
		return nil, fmt.Errorf("invalid page size: page=%d limit=%d", page, limit)
	}

	offset := (page - 1) * limit
	if offset >= len(r.createdAtArticles) {
		return []domain.Article{}, nil
	}

	end := offset + limit
	if end > len(r.createdAtArticles) {
		end = len(r.createdAtArticles)
	}

	result := make([]domain.Article, end-offset)
	for _, record := range r.createdAtArticles[offset:end] {
		if article, exists := r.articles[record.id]; exists {
			result = append(result, article)
		}
	}

	return result, nil
}

func (r *MemoryRepository) ArticleRecordsCount(article domain.Article) (int32, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return int32(len(r.articles)), nil
}

func (r *MemoryRepository) CreateComment(comment domain.Comment) (domain.Comment, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.comments[comment.ID]; !exists {
		r.comments[comment.ID] = comment

		if comment.ParentID != uuid.Nil {
			r.commentReply[comment.ParentID] = append(r.commentReply[comment.ParentID], &comment)
		}

		return comment, nil
	}
	return domain.Comment{}, fmt.Errorf("adding a non-unique value %+v", comment.ID)
}

func (r *MemoryRepository) GetComments(articleId uuid.UUID, page, limit int) ([]domain.Comment, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	
	rootComments := r.rootComments(articleId)

	start := (page - 1) * limit
	if start >= len(rootComments) {
		return []domain.Comment{}, nil
	}

	end := start + limit
	if end > len(rootComments) {
		end = len(rootComments)
	}

	commentPage := rootComments[start:end]

	for i := range commentPage {
		r.addReply(&commentPage[i])
	}

	return commentPage, nil
}

func (r *MemoryRepository) CommentsRecordCount(comment domain.Comment) (int32, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return int32(len(r.comments)), nil
}

func (r *MemoryRepository) rootComments(articleId uuid.UUID) []domain.Comment {

	rootComments := []domain.Comment{}
	for _, comment := range r.comments {
		if comment.ParentID == uuid.Nil && comment.ArticleID == articleId {
			rootComments = append(rootComments, comment)
		}
	}

	sort.Slice(rootComments, func(i, j int) bool {
		return rootComments[i].CreatedAt.Before(rootComments[j].CreatedAt)
	})

	return rootComments
}

func (r *MemoryRepository) addReply(comment *domain.Comment) {
	if replies, ok := r.commentReply[comment.ID]; ok {
		for _, reply := range replies {
			r.addReply(reply)
			comment.Replies = append(comment.Replies, *reply)
		}
		sort.Slice(comment.Replies, func(i, j int) bool {
			return comment.Replies[i].CreatedAt.Before(comment.Replies[j].CreatedAt)
		})
	}
}
