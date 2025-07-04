package repository

import (
	"github.com/krishnakumar-learn/go-todo-api/models"
	"gorm.io/gorm"
)

// mysqlTodoRepository implements the TodoRepository interface using a MySQL database via GORM.
type mysqlTodoRepository struct {
	db *gorm.DB
}

// NewMySQLTodoRepository creates a new TodoRepository backed by MySQL using GORM.
func NewMySQLTodoRepository(db *gorm.DB) TodoRepository {
	return &mysqlTodoRepository{db}
}

// FindAll retrieves all todos from the MySQL database.
func (r *mysqlTodoRepository) FindAll() ([]models.Todo, error) {
	var todos []models.Todo
	err := r.db.Find(&todos).Error
	return todos, err
}

// FindByID retrieves a todo by its ID from the MySQL database.
func (r *mysqlTodoRepository) FindByID(id uint) (*models.Todo, error) {
	var todo models.Todo
	err := r.db.First(&todo, id).Error
	return &todo, err
}

// Create inserts a new todo into the MySQL database.
func (r *mysqlTodoRepository) Create(todo *models.Todo) error {
	return r.db.Create(todo).Error
}

// Update saves changes to an existing todo in the MySQL database.
func (r *mysqlTodoRepository) Update(todo *models.Todo) error {
	return r.db.Save(todo).Error
}

// Delete removes a todo by its ID from the MySQL database.
func (r *mysqlTodoRepository) Delete(id uint) error {
	return r.db.Delete(&models.Todo{}, id).Error
}
