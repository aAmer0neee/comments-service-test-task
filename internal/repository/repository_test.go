package repository

import (
	"testing"
	"time"

	"github.com/aAmer0neee/comments-service-test-task/internal/domain"
	"github.com/aAmer0neee/comments-service-test-task/internal/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateArticle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository_mock.NewMockRepository(ctrl)

	article := domain.Article{
		ID:                uuid.New(),
		Content:           "Test Article",
		CreatedAt:         time.Now(),
		CommentPermission: true,
	}

	mockRepo.EXPECT().CreateArticle(article).Return(article, nil).Times(1)

	result, err := mockRepo.CreateArticle(article)

	assert.NoError(t, err)
	assert.Equal(t, article.Content, result.Content)
	assert.Equal(t, article.CommentPermission, result.CommentPermission)
}

func TestGetArticle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository_mock.NewMockRepository(ctrl)
	articleID := uuid.New()

	article := domain.Article{
		ID:                articleID,
		Content:           "Test Article",
		CreatedAt:         time.Now(),
		CommentPermission: true,
	}

	mockRepo.EXPECT().GetArticle(articleID).Return(article, nil).Times(1)

	result, err := mockRepo.GetArticle(articleID)

	assert.NoError(t, err)
	assert.Equal(t, articleID, result.ID)
	assert.Equal(t, article.Content, result.Content)
}

func TestGetListArticles(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository_mock.NewMockRepository(ctrl)

	article1 := domain.Article{
		ID:                uuid.New(),
		Content:           "Test Article 1",
		CreatedAt:         time.Now(),
		CommentPermission: true,
	}

	article2 := domain.Article{
		ID:                uuid.New(),
		Content:           "Test Article 2",
		CreatedAt:         time.Now(),
		CommentPermission: true,
	}

	mockRepo.EXPECT().GetListArticles(1, 2).Return([]domain.Article{article1, article2}, nil).Times(1)

	result, err := mockRepo.GetListArticles(1, 2)

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, article1.Content, result[0].Content)
	assert.Equal(t, article2.Content, result[1].Content)
}

func TestArticleRecordsCount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository_mock.NewMockRepository(ctrl)

	article := domain.Article{
		ID:                uuid.New(),
		Content:           "Test Article",
		CreatedAt:         time.Now(),
		CommentPermission: true,
	}

	mockRepo.EXPECT().ArticleRecordsCount(article).Return(int32(5), nil).Times(1)

	result, err := mockRepo.ArticleRecordsCount(article)

	assert.NoError(t, err)
	assert.Equal(t, int32(5), result)
}

func TestCreateComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository_mock.NewMockRepository(ctrl)

	comment := domain.Comment{
		ID:        uuid.New(),
		Content:   "Test Comment",
		CreatedAt: time.Now(),
		ArticleID: uuid.New(),
		ParentID:  uuid.New(),
	}

	mockRepo.EXPECT().CreateComment(comment).Return(comment, nil).Times(1)

	result, err := mockRepo.CreateComment(comment)

	assert.NoError(t, err)
	assert.Equal(t, comment.Content, result.Content)
	assert.Equal(t, comment.ArticleID, result.ArticleID)
}

func TestGetComments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository_mock.NewMockRepository(ctrl)
	articleID := uuid.New()

	comment1 := domain.Comment{
		ID:        uuid.New(),
		Content:   "Test Comment 1",
		CreatedAt: time.Now(),
		ArticleID: articleID,
	}

	comment2 := domain.Comment{
		ID:        uuid.New(),
		Content:   "Test Comment 2",
		CreatedAt: time.Now(),
		ArticleID: articleID,
	}

	mockRepo.EXPECT().GetComments(articleID, 1, 2).Return([]domain.Comment{comment1, comment2}, nil).Times(1)

	result, err := mockRepo.GetComments(articleID, 1, 2)

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, comment1.Content, result[0].Content)
	assert.Equal(t, comment2.Content, result[1].Content)
}

func TestCommentsRecordCount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository_mock.NewMockRepository(ctrl)

	comment := domain.Comment{
		ID:        uuid.New(),
		Content:   "Test Comment",
		CreatedAt: time.Now(),
		ArticleID: uuid.New(),
	}

	mockRepo.EXPECT().CommentsRecordCount(comment).Return(int32(3), nil).Times(1)

	result, err := mockRepo.CommentsRecordCount(comment)

	assert.NoError(t, err)
	assert.Equal(t, int32(3), result)
}
