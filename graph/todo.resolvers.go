package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graphql-server/domain/entity"
	"graphql-server/graph/generated"
	"graphql-server/graph/model"

	"gorm.io/gorm"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	tid := entity.GenUuid()
	if err := r.dbc.RunInTransactionSession(func(conn *gorm.DB) error {
		e := entity.Todo{
			ID:     tid,
			Title:  input.Title,
			Done:   false,
			UserID: input.UserID,
		}

		return conn.Omit().Create(&e).Error
	}); err != nil {
		return nil, err
	}

	return &model.Todo{}, nil
}

func (r *queryResolver) Todos(ctx context.Context, userID string) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *todoResolver) Title(ctx context.Context, obj *model.Todo) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }
