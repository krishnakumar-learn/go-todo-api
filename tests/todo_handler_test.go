package tests

import (
	"bytes"
	"encoding/json"
	"github.com/krishnakumar-learn/go-todo-api/handlers"
	"github.com/krishnakumar-learn/go-todo-api/repository"
	"github.com/krishnakumar-learn/go-todo-api/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)
func setupRouter(h *handlers.TodoHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	todos := r.Group("/todos")
	{
		todos.GET("/", h.GetTodos)
		todos.POST("/", h.CreateTodo)
		todos.GET("/:id", h.GetTodo)
		todos.PUT("/:id", h.UpdateTodo)
		todos.DELETE("/:id", h.DeleteTodo)
	}
	return r
}

func TestGetTodos(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockTodoRepository(ctrl)
	mockRepo.EXPECT().FindAll().Return([]models.Todo{
		{ID: 1, Title: "Task1", Description: "Test", Completed: false},
	}, nil)

	handler := handlers.NewTodoHandler(mockRepo)
	r := setupRouter(handler)

	req, _ := http.NewRequest("GET", "/todos/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockTodoRepository(ctrl)
	mockRepo.EXPECT().Create(gomock.Any()).Return(nil)

	handler := handlers.NewTodoHandler(mockRepo)
	r := setupRouter(handler)

	todo := models.Todo{Title: "Task2", Description: "Create Test", Completed: false}
	body, _ := json.Marshal(todo)

	req, _ := http.NewRequest("POST", "/todos/", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockTodoRepository(ctrl)
	mockRepo.EXPECT().FindByID(uint(1)).Return(&models.Todo{ID: 1, Title: "Task1"}, nil)

	handler := handlers.NewTodoHandler(mockRepo)
	r := setupRouter(handler)

	req, _ := http.NewRequest("GET", "/todos/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	oldTodo := &models.Todo{ID: 1, Title: "Old", Description: "Old Desc", Completed: false}
	updated := &models.Todo{ID: 1, Title: "Updated", Description: "Updated Desc", Completed: true}

	mockRepo := repository.NewMockTodoRepository(ctrl)
	mockRepo.EXPECT().FindByID(uint(1)).Return(oldTodo, nil)
	mockRepo.EXPECT().Update(oldTodo).Return(nil)

	handler := handlers.NewTodoHandler(mockRepo)
	r := setupRouter(handler)

	body, _ := json.Marshal(updated)
	req, _ := http.NewRequest("PUT", "/todos/1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockTodoRepository(ctrl)
	mockRepo.EXPECT().Delete(uint(1)).Return(nil)

	handler := handlers.NewTodoHandler(mockRepo)
	r := setupRouter(handler)

	req, _ := http.NewRequest("DELETE", "/todos/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}