package repository

import (
	"fmt"

	"github.com/aAmer0neee/comments-service-test-task/internal/config"
	"github.com/aAmer0neee/comments-service-test-task/internal/domain"
	"github.com/google/uuid"
)

type Repository interface {
	CreateArticle(article domain.Article) (domain.Article, error)
	GetArticle(id uuid.UUID) (domain.Article, error)
	GetListArticles(page, limit int) ([]domain.Article, error)

	CreateComment(comment domain.Comment) (domain.Comment, error)
	GetRootComments(page, limit int) ([]domain.Comment, error)
	RecordsCount(a interface{}) (int32, error)
}

func SwitchRepository(cfg *config.Cfg) (Repository, error) {
	switch cfg.RepositoryMode {
	case "postgres":
		return ConnectPostgres(*cfg)

	case "memory":
		return nil, fmt.Errorf("not implemented")
	default:
		return nil, fmt.Errorf("not implemented")
	}
}
