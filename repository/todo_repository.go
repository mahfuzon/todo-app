package repository

import (
	"github.com/mahfuzon/to-do-list-app/model"
	"gorm.io/gorm"
)

type TodoRepository interface {
	Create(todo model.Todo) (model.Todo, error)
	Find(id int) (model.Todo, error)
	All() ([]model.Todo, error)
	Update(todo model.Todo) (model.Todo, error)
	Delete(todo model.Todo) error
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{db: db}
}

func (todoRepository *todoRepository) Create(todo model.Todo) (model.Todo, error) {
	err := todoRepository.db.Create(&todo).Error
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (todoRepository *todoRepository) Find(id int) (model.Todo, error) {
	todo := model.Todo{}
	err := todoRepository.db.First(&todo, id).Error
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (todoRepository *todoRepository) All() ([]model.Todo, error) {
	var todos []model.Todo
	err := todoRepository.db.Find(&todos).Error
	if err != nil {
		return todos, err
	}

	return todos, nil
}

func (todoRepository *todoRepository) Update(todo model.Todo) (model.Todo, error) {
	err := todoRepository.db.Save(&todo).Error
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (todoRepository *todoRepository) Delete(todo model.Todo) error {
	err := todoRepository.db.Delete(&todo).Error
	if err != nil {
		return err
	}

	return nil
}
