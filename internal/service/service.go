package service

import (
	"log/slog"

	"github.com/aAmer0neee/comments-service-test-task/internal/domain"
	"github.com/aAmer0neee/comments-service-test-task/internal/repository"
	"github.com/google/uuid"
)

//go:generate mockgen -source=service.go -destination=mocks/service_mock.go -package=service_mock
type Service interface {
	PostArticle(article domain.Article) (domain.Article, error)
	GetArticlesList(page, limit int32) ([]domain.Article, int32, error)
	PostComment(comment domain.Comment) (domain.Comment, error)
	GetArticle(id uuid.UUID, commentPage, pageLimit int32) (domain.Article, []domain.Comment, error)
}

func InitService(r repository.Repository, s slog.Logger) Service {
	return InitArticleService(r, s)
}
