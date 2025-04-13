package service

import (
	"log/slog"
	_ "sort"

	"github.com/aAmer0neee/comments-service-test-task/internal/domain"
	"github.com/aAmer0neee/comments-service-test-task/internal/repository"
	"github.com/google/uuid"
)

type Service struct {
	Repo repository.Repository
	Log  slog.Logger
}

func InitService(r repository.Repository, l slog.Logger) Service {
	return Service{
		Repo: r,
		Log:  l,
	}
}

func (s *Service) PostArticle(article domain.Article) (domain.Article, error) {
	article.ID = uuid.New()
	updateArticle, err := s.Repo.CreateArticle(article)
	if err != nil {
		s.Log.Info("ошибка добавления записи", "message", err.Error())
		return domain.Article{}, err
	}
	return updateArticle, nil
}

func (s *Service) GetArticlesList(page, limit int32) ([]domain.Article, int32, error) {
	articles, err := s.Repo.GetListArticles(int(page), int(limit))
	if err != nil {
		s.Log.Info("ошибка получения списка статей", "message", err.Error())
		return []domain.Article{}, 0, err
	}
	recordCount, err := s.Repo.RecordsCount(domain.Article{})
	if err != nil {
		s.Log.Info("ошибка получения количества статей", "message", err.Error())
		return []domain.Article{}, 0, err
	}
	return articles, recordCount, nil
}

func (s *Service) PostComment(comment domain.Comment) (domain.Comment, error) {
	comment.ID = uuid.New()
	updateComment, err := s.Repo.CreateComment(comment)
	if err != nil {
		s.Log.Info("ошибка добавления записи", "message", err.Error())
		return domain.Comment{}, err
	}
	return updateComment, nil
}

func (s *Service) GetArticle(id uuid.UUID, commentPage, pageLimit int32) (domain.Article, []domain.Comment, error) {
	article, err := s.Repo.GetArticle(id)
	if err != nil {
		s.Log.Info("ошибка получения статьи", "message", err.Error())
		return domain.Article{}, []domain.Comment{}, err
	}
	comments, err := s.Repo.GetRootComments(int(commentPage), int(pageLimit))
	if err != nil {
		s.Log.Info("ошибка получения списка комментариев", "message", err.Error())
		return domain.Article{}, []domain.Comment{}, nil
	}
	return article, comments, nil
}

/* func (s *Service)normalzeComentsTree(tree []domain.Comment)[]domain.Comment{
	commentMap := make (map[uuid.UUID][]domain.Comment)

	rootComments := []domain.Comment{}
	for _, comment := range tree {
		if comment.ParentID == uuid.Nil {
			rootComments = append(rootComments, comment)
		} else {
			commentMap[comment.ParentID] = append(commentMap[comment.ParentID], comment)
		}
	}

	var addComment func(*domain.Comment)

	addComment = func(c *domain.Comment) {
		if replies, ok := commentMap[c.ID]; ok {

			s.sortCreatedAt(replies)

			for i:= range replies {
				addComment(&replies[i])
				c.Replies = append(c.Replies, replies[i])
			}

		}
	}

	for i := range rootComments {
		addComment(&rootComments[i])
	}
	return rootComments
}

func (s *Service)sortCreatedAt(comments []domain.Comment){
	sort.Slice(comments, func(i, j int) bool {
		return comments[i].CreatedAt.Before(comments[j].CreatedAt)
	})
} */
