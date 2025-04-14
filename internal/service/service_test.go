package service

import (
	"testing"

	"github.com/aAmer0neee/comments-service-test-task/internal/domain"
	"github.com/aAmer0neee/comments-service-test-task/internal/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

func TestPostArticle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service_mock.NewMockService(ctrl)

	article := domain.Article{
		ID:      uuid.New(),
		Content: "Test Article",
	}

	mockService.EXPECT().PostArticle(article).Return(article, nil).Times(1)

	result, err := mockService.PostArticle(article)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result.Content != article.Content {
		t.Errorf("Expected content %s, got %s", article.Content, result.Content)
	}
}

func TestGetArticlesList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service_mock.NewMockService(ctrl)

	articles := []domain.Article{
		{ID: uuid.New(), Content: "Article 1"},
		{ID: uuid.New(), Content: "Article 2"},
	}

	mockService.EXPECT().GetArticlesList(int32(1), int32(10)).Return(articles, int32(2), nil).Times(1)

	result, count, err := mockService.GetArticlesList(int32(1), int32(10))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(result) != len(articles) {
		t.Errorf("Expected %d articles, got %d", len(articles), len(result))
	}
	if count != 2 {
		t.Errorf("Expected count %d, got %d", 2, count)
	}
}

func TestPostComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service_mock.NewMockService(ctrl)

	comment := domain.Comment{
		ID:      uuid.New(),
		Content: "Test Comment",
	}

	mockService.EXPECT().PostComment(comment).Return(comment, nil).Times(1)

	result, err := mockService.PostComment(comment)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.Content != comment.Content {
		t.Errorf("Expected content %s, got %s", comment.Content, result.Content)
	}
}

func TestGetArticle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service_mock.NewMockService(ctrl)

	articleID := uuid.New()
	expectedArticle := domain.Article{
		ID:      articleID,
		Content: "Sample Article",
	}
	expectedComments := []domain.Comment{
		{ID: uuid.New(), Content: "Comment 1"},
		{ID: uuid.New(), Content: "Comment 2"},
	}

	mockService.EXPECT().
		GetArticle(articleID, int32(1), int32(10)).
		Return(expectedArticle, expectedComments, nil).
		Times(1)

	resultArticle, resultComments, err := mockService.GetArticle(articleID, int32(1), int32(10))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resultArticle.ID != expectedArticle.ID {
		t.Errorf("Expected article ID %v, got %v", expectedArticle.ID, resultArticle.ID)
	}
	if len(resultComments) != len(expectedComments) {
		t.Errorf("Expected %d comments, got %d", len(expectedComments), len(resultComments))
	}
}
