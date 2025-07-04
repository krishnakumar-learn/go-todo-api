package repository

import "github.com/krishnakumar-learn/go-todo-api/models"

// TodoRepository defines the interface for CRUD operations on Todo items.
type TodoRepository interface {
	// FindAll retrieves all Todo items.
	FindAll() ([]models.Todo, error)
	// FindByID retrieves a Todo item by its ID.
	FindByID(id uint) (*models.Todo, error)
	// Create adds a new Todo item.
	Create(todo *models.Todo) error
	// Update modifies an existing Todo item.
	Update(todo *models.Todo) error
	// Delete removes a Todo item by its ID.
	Delete(id uint) error
}
