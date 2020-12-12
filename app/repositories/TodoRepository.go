package repositories

import (
	"github.com/brandon-julio-t/golang-graphql-todo/graph/model"
	"gorm.io/gorm"
)

type TodoRepository struct {
	DB *gorm.DB
}

func (r TodoRepository) Save(todo *model.Todo) (*model.Todo, error) {
	if err := r.DB.Create(todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

func (r TodoRepository) GetAll() ([]*model.Todo, error) {
	var todo []*model.Todo

	if err := r.DB.Find(&todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

func (r TodoRepository) GetById(id string) (*model.Todo, error) {
	todo := &model.Todo{}

	if err := r.DB.First(todo, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

func (r TodoRepository) Update(todo *model.Todo) (*model.Todo, error) {
	if err := r.DB.Save(todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

func (r TodoRepository) Delete(todo *model.Todo) (*model.Todo, error) {
	if err := r.DB.Delete(todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}
