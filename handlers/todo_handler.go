package handlers

import (
	"net/http"
	"strconv"

	"github.com/krishnakumar-learn/go-todo-api/models"
	"github.com/krishnakumar-learn/go-todo-api/repository"

	"github.com/gin-gonic/gin"
)

// TodoHandler handles HTTP requests for todo operations.
type TodoHandler struct {
	repo repository.TodoRepository
}

// NewTodoHandler creates a new TodoHandler with the given repository.
func NewTodoHandler(repo repository.TodoRepository) *TodoHandler {
	return &TodoHandler{repo}
}

// GetTodos handles GET requests to retrieve all todos.
func (h *TodoHandler) GetTodos(c *gin.Context) {
	todos, err := h.repo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}

// GetTodo handles GET requests to retrieve a single todo by ID.
func (h *TodoHandler) GetTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := h.repo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

// CreateTodo handles POST requests to create a new todo.
func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.repo.Create(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, todo)
}

// UpdateTodo handles PUT requests to update an existing todo by ID.
func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := h.repo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	if err := c.ShouldBindJSON(todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.repo.Update(todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)
}

// DeleteTodo handles DELETE requests to remove a todo by ID.
func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.repo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
