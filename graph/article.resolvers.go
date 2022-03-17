package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"graphql-server/domain/entity"
	"graphql-server/graph/generated"
	"graphql-server/graph/model"

	"gorm.io/gorm"
)

func (r *articleResolver) User(ctx context.Context, obj *model.Article) (*model.User, error) {
	return r.Query().User(ctx, obj.UserID)
}

func (r *mutationResolver) CreateArticle(ctx context.Context, input model.NewArticle) (*model.Article, error) {
	tid := entity.GenUuid()
	if err := r.dbc.RunInTransactionSession(func(conn *gorm.DB) error {
		e := entity.Article{
			ID:     tid,
			Title:  input.Title,
			Body:   input.Body,
			UserID: entity.UserID(input.UserID),
		}

		return conn.Omit().Create(&e).Error
	}); err != nil {
		return nil, err
	}

	return &model.Article{}, nil
}

func (r *queryResolver) Articles(ctx context.Context, userID string) ([]*model.Article, error) {
	var es []entity.Article
	if err := r.dbc.RunInSession(func(conn *gorm.DB) error {
		return conn.Where(entity.Article{UserID: entity.UserID(userID)}).Find(&es).Error
	}); err != nil {
		return nil, err
	}

	items := make([]*model.Article, 0, len(es))
	for _, e := range es {
		items = append(items, &model.Article{
			ID:     e.ID,
			Title:  e.Title,
			Body:   e.Body,
			UserID: e.UserID.String(),
		})
	}

	return items, nil
}

func (r *queryResolver) Article(ctx context.Context, id string) (*model.Article, error) {
	var e entity.Article
	if err := r.dbc.RunInSession(func(conn *gorm.DB) error {
		return conn.Where(entity.Article{ID: id}).Take(&e).Error
	}); err != nil {
		return nil, err
	}

	return &model.Article{
		ID:     e.ID,
		Title:  e.Title,
		Body:   e.Body,
		UserID: e.UserID.String(),
	}, nil
}

// Article returns generated.ArticleResolver implementation.
func (r *Resolver) Article() generated.ArticleResolver { return &articleResolver{r} }

type articleResolver struct{ *Resolver }
