package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.70

import (
	"context"
	"fmt"

	"github.com/aAmer0neee/comments-service-test-task/graph/model"
	"github.com/aAmer0neee/comments-service-test-task/graph/runtime"
	"github.com/aAmer0neee/comments-service-test-task/internal/mappers"
)

// CreateArticle is the resolver for the createArticle field.
func (r *mutationResolver) CreateArticle(ctx context.Context, input model.ArticleCreateInput) (model.ArticleCreateResponse, error) {
	if response, err := r.Service.PostArticle(mappers.InputToDomainArticle(input)); err != nil {
		return model.ArticleCreateBadRequest{
			Message: "ошибка при создании статьи",
		}, fmt.Errorf("ошибка при создании статьи")
	} else {
		return model.ArticleCreateOk{
			Article: mappers.DomainArticleToResponse(response),
		}, nil
	}
}

// Mutation returns runtime.MutationResolver implementation.
func (r *Resolver) Mutation() runtime.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
