package repository

import (
	"github.com/aAmer0neee/comments-service-test-task/internal/config"
	"github.com/aAmer0neee/comments-service-test-task/internal/domain"
	"github.com/google/uuid"
)

type Repository interface {
	CreateArticle(article domain.Article) (domain.Article, error)
	GetArticle(id uuid.UUID) (domain.Article, error)
	GetListArticles(page, limit int) ([]domain.Article, error)
	ArticleRecordsCount(article domain.Article) (int32, error)

	CreateComment(comment domain.Comment) (domain.Comment, error)
	GetComments(articleId uuid.UUID, page, limit int) ([]domain.Comment, error)
	CommentsRecordCount(comment domain.Comment) (int32, error)
}

func InitRepository(cfg *config.Cfg) (Repository, error) {
	switch cfg.RepositoryMode {
	case "postgres":
		return ConnectPostgres(*cfg)

	case "in-memory":
		return initInMemory()
	default:
		return initInMemory()
	}
}
