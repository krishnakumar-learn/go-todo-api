package repository

import "github.com/krishnakumar-learn/go-todo-api/models"

type TodoRepository interface {
	FindAll() ([]models.Todo, error)
	FindByID(id uint) (*models.Todo, error)
	Create(todo *models.Todo) error
	Update(todo *models.Todo) error
	Delete(id uint) error
}
