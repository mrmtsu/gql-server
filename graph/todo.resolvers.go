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

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	tid := entity.GenUuid()
	if err := r.dbc.RunInTransactionSession(func(conn *gorm.DB) error {
		e := entity.Todo{
			ID:     tid,
			Title:  input.Title,
			Done:   false,
			UserID: entity.UserID(input.UserID),
		}

		return conn.Omit().Create(&e).Error
	}); err != nil {
		return nil, err
	}

	return &model.Todo{}, nil
}

func (r *queryResolver) Todos(ctx context.Context, userID string) ([]*model.Todo, error) {
	var es []entity.Todo
	if err := r.dbc.RunInSession(func(conn *gorm.DB) error {
		return conn.Where(entity.Todo{UserID: entity.UserID(userID)}).Find(&es).Error
	}); err != nil {
		return nil, err
	}

	items := make([]*model.Todo, 0, len(es))
	for _, e := range es {
		items = append(items, &model.Todo{
			ID:     e.ID,
			Title:  e.Title,
			Done:   e.Done,
			UserID: e.UserID.String(),
		})
	}

	return items, nil
}

func (r *queryResolver) Todo(ctx context.Context, id string) (*model.Todo, error) {
	var e entity.Todo
	if err := r.dbc.RunInSession(func(conn *gorm.DB) error {
		return conn.Where(entity.Todo{ID: id}).Take(&e).Error
	}); err != nil {
		return nil, err
	}

	return &model.Todo{
		ID:     e.ID,
		Title:  e.Title,
		Done:   e.Done,
		UserID: e.UserID.String(),
	}, nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return r.Query().User(ctx, obj.UserID)
}

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }
