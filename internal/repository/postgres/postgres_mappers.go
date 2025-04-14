package postgres

import (
	"github.com/aAmer0neee/comments-service-test-task/internal/domain"
)

func convertArticleToGorm(article domain.Article) Article {
	return Article{
		ID:                article.ID,
		Content:           article.Content,
		CreatedAt:         article.CreatedAt,
		CommentPermission: article.CommentPermission,
	}
}

func convertGormToArticle(article Article) domain.Article {
	return domain.Article{
		ID:                article.ID,
		Content:           article.Content,
		CreatedAt:         article.CreatedAt,
		CommentPermission: article.CommentPermission,
	}
}

func convertCommentToGorm(comment domain.Comment) Comment {
	return Comment{
		ID:        comment.ID,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
		ArticleID: comment.ArticleID,
		ParentID:  comment.ParentID,
	}
}

func convertGormToComment(comment Comment) domain.Comment {
	return domain.Comment{
		ID:        comment.ID,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
		ArticleID: comment.ArticleID,
		ParentID:  comment.ParentID,
	}
}
