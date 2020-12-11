package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"math/rand"

	"github.com/brandon-julio-t/golang-graphql-todo/graph/generated"
	"github.com/brandon-julio-t/golang-graphql-todo/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		ID:   fmt.Sprintf("T%d", rand.Int()),
		Text: input.Text,
		Done: false,
	}

	r.allTodo = append(r.allTodo, todo)
	return todo, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.UpdateTodo) (*model.Todo, error) {
	for i, todo := range r.allTodo {
		if todo.ID == input.ID {
			todo.Text = input.Text
			r.allTodo[i] = todo
			return todo, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("Todo with ID %s not found.", input.ID))
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, input model.TodoByID) (*model.Todo, error) {
	allTodo := make([]*model.Todo, 0)
	var deleted *model.Todo = nil

	for _, todo := range r.allTodo {
		if todo.ID != input.ID {
			allTodo = append(allTodo, todo)
		} else {
			deleted = todo
		}
	}

	if deleted == nil {
		return nil, errors.New(fmt.Sprintf("Todo with ID %s not found.", input.ID))
	}

	r.allTodo = allTodo
	return deleted, nil
}

func (r *mutationResolver) ToggleTodoDoneStatus(ctx context.Context, input model.TodoByID) (*model.Todo, error) {
	for i, todo := range r.allTodo {
		if todo.ID == input.ID {
			todo.Done = !todo.Done
			r.allTodo[i] = todo
			return todo, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("Todo with ID %s not found", input.ID))
}

func (r *queryResolver) AllTodo(ctx context.Context) ([]*model.Todo, error) {
	if len(r.allTodo) == 0 {
		r.allTodo = append(r.allTodo, &model.Todo{
			ID:   fmt.Sprintf("T%d", rand.Int()),
			Text: "Todo #1",
			Done: false,
		})
		r.allTodo = append(r.allTodo, &model.Todo{
			ID:   fmt.Sprintf("T%d", rand.Int()),
			Text: "Todo #2",
			Done: true,
		})
		r.allTodo = append(r.allTodo, &model.Todo{
			ID:   fmt.Sprintf("T%d", rand.Int()),
			Text: "Todo #3",
			Done: false,
		})
	}

	return r.allTodo, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
