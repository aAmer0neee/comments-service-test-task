package repository

import (
	"github.com/aAmer0neee/comments-service-test-task/internal/domain"
)

type Repository interface {
	PostArticle(article domain.Article) error
	GetArticle(id int) (domain.Article, error)
	GetListArticles(page, limit int)([]domain.Article, error)

	PostComment(comment domain.Comment)
	GetListComments(page, limit int) ([]domain.Comment, error)
	RecordsCount(a interface{})(int, error)
}
