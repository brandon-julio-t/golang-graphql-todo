package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/brandon-julio-t/golang-graphql-todo/graph/generated"
	"github.com/brandon-julio-t/golang-graphql-todo/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		ID:          fmt.Sprintf("T%d", rand.Int()),
		Text:        input.Text,
		Done:        false,
		AssistantID: input.AssistantInitial, // TODO
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, input model.FindTodoByID) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ToggleTodoDoneStatus(ctx context.Context, input model.FindTodoByID) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateAssistant(ctx context.Context, input model.NewAssistant) (*model.Assistant, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateAssistant(ctx context.Context, input model.NewAssistant) (*model.Assistant, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteAssistant(ctx context.Context, input model.FindAssistantByID) (*model.Assistant, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Assistants(ctx context.Context) ([]*model.Assistant, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	if len(r.todos) == 0 {
		r.todos = append(r.todos, &model.Todo{
			ID:          fmt.Sprintf("T%d", rand.Int()),
			Text:        "Todo #1",
			Done:        false,
			AssistantID: "Random User",
		})
		r.todos = append(r.todos, &model.Todo{
			ID:          fmt.Sprintf("T%d", rand.Int()),
			Text:        "Todo #1",
			Done:        true,
			AssistantID: "Random User",
		})
		r.todos = append(r.todos, &model.Todo{
			ID:          fmt.Sprintf("T%d", rand.Int()),
			Text:        "Todo #1",
			Done:        false,
			AssistantID: "Random User",
		})
	}

	return r.todos, nil
}

func (r *todoResolver) Assistant(ctx context.Context, obj *model.Todo) (*model.Assistant, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
