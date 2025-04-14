package mappers

import (
	"github.com/aAmer0neee/comments-service-test-task/graph/model"
	"github.com/aAmer0neee/comments-service-test-task/internal/domain"
	"github.com/google/uuid"
)

func InputToDomainArticle(article model.ArticleCreateInput) domain.Article {
	return domain.Article{
		Content:           article.Content,
		CommentPermission: *article.CommentPermission,
	}
}

func DomainArticleToResponse(article domain.Article) *model.Article {
	return &model.Article{
		ID:                article.ID,
		Content:           article.Content,
		CreatedAt:         article.CreatedAt.String(),
		CommentPermission: article.CommentPermission,
	}
}

func DomainArticlesListToResponse(articles []domain.Article) []*model.Article {
	response := []*model.Article{}
	for _, artcle := range articles {
		converted := DomainArticleToResponse(artcle)
		if converted != nil {
			response = append(response, converted)
		}
	}
	return response
}

func InputToDomainComment(comment model.CommentCreateInput) domain.Comment {
	nilUuid := uuid.Nil
	if comment.ParentID != nil {
		nilUuid = *comment.ParentID
	}
	return domain.Comment{
		Content:   comment.Content,
		ArticleID: comment.ArticleID,
		ParentID:  nilUuid,
	}
}

func DomainCommentToResponse(comment domain.Comment) *model.Comment {
	var replies []*model.Comment
	if len(comment.Replies) > 0 {
		replies = DomainCommentsListToResponse(comment.Replies)
	}
	return &model.Comment{
		ID:        comment.ID,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt.String(),
		ArticleID: comment.ArticleID,

		ParentID: &comment.ParentID,
		Replies:  replies,
	}
}

func DomainCommentsListToResponse(articles []domain.Comment) []*model.Comment {
	response := []*model.Comment{}
	for _, artcle := range articles {
		converted := DomainCommentToResponse(artcle)
		if converted != nil {
			response = append(response, converted)
		}
	}
	return response
}
