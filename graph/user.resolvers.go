package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"graphql-server/domain/entity"
	"graphql-server/graph/model"

	"gorm.io/gorm"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	uid := entity.GenUuid()
	if err := r.dbc.RunInTransactionSession(func(conn *gorm.DB) error {
		e := entity.User{
			ID:   uid,
			Name: input.Name,
		}
		return conn.Omit().Create(&e).Error
	}); err != nil {
		return "", err
	}

	return uid, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	var e entity.User
	if err := r.dbc.RunInSession(func(conn *gorm.DB) error {
		return conn.Where(entity.User{ID: id}).Take(&e).Error
	}); err != nil {
		return nil, err
	}

	return &model.User{
		ID:   e.ID,
		Name: e.Name,
	}, nil
}
