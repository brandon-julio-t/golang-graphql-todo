package services

import (
	"github.com/brandon-julio-t/golang-graphql-todo/app/repositories"
	"github.com/brandon-julio-t/golang-graphql-todo/graph/model"
	"github.com/google/uuid"
)

type TodoService struct {
	Repository *repositories.TodoRepository
}

func (s TodoService) CreateTodoFromInput(input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		ID:   uuid.Must(uuid.NewRandom()).String(),
		Text: input.Text,
		Done: false,
	}

	return s.Repository.Save(todo)
}

func (s *TodoService) GetAllTodo() ([]*model.Todo, error) {
	return s.Repository.GetAll()
}

func (s TodoService) ToggleTodoDoneStatus(input model.TodoByID) (*model.Todo, error) {
	todo, err := s.Repository.GetById(input.ID)

	if err != nil {
		return nil, err
	}

	todo.Done = !todo.Done

	return s.Repository.Update(todo)
}

func (s TodoService) UpdateTodoFromInput(input model.UpdateTodo) (*model.Todo, error) {
	todo, err := s.Repository.GetById(input.ID)

	if err != nil {
		return nil, err
	}

	todo.Text = input.Text

	return s.Repository.Update(todo)
}

func (s TodoService) DeleteTodo(input model.TodoByID) (*model.Todo, error) {
	todo, err := s.Repository.GetById(input.ID)

	if err != nil {
		return nil, err
	}

	return s.Repository.Delete(todo)
}
