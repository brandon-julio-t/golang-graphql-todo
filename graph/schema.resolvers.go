package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/brandon-julio-t/golang-graphql-todo/graph/generated"
	"github.com/brandon-julio-t/golang-graphql-todo/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	return r.TodoService.CreateTodoFromInput(input)
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.UpdateTodo) (*model.Todo, error) {
	return r.TodoService.UpdateTodoFromInput(input)
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, input model.TodoByID) (*model.Todo, error) {
	return r.TodoService.DeleteTodo(input)
}

func (r *mutationResolver) ToggleTodoDoneStatus(ctx context.Context, input model.TodoByID) (*model.Todo, error) {
	return r.TodoService.ToggleTodoDoneStatus(input)
}

func (r *queryResolver) AllTodo(ctx context.Context) ([]*model.Todo, error) {
	return r.TodoService.GetAllTodo()
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
