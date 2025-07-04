package repository

import (
	"github.com/krishnakumar-learn/go-todo-api/models"
	"gorm.io/gorm"
)

type mysqlTodoRepository struct {
	db *gorm.DB
}

func NewMySQLTodoRepository(db *gorm.DB) TodoRepository {
	return &mysqlTodoRepository{db}
}

func (r *mysqlTodoRepository) FindAll() ([]models.Todo, error) {
	var todos []models.Todo
	err := r.db.Find(&todos).Error
	return todos, err
}

func (r *mysqlTodoRepository) FindByID(id uint) (*models.Todo, error) {
	var todo models.Todo
	err := r.db.First(&todo, id).Error
	return &todo, err
}

func (r *mysqlTodoRepository) Create(todo *models.Todo) error {
	return r.db.Create(todo).Error
}

func (r *mysqlTodoRepository) Update(todo *models.Todo) error {
	return r.db.Save(todo).Error
}

func (r *mysqlTodoRepository) Delete(id uint) error {
	return r.db.Delete(&models.Todo{}, id).Error
}
