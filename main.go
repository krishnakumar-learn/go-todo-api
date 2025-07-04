package main

import (
	"github.com/krishnakumar-learn/go-todo-api/db"
	"github.com/krishnakumar-learn/go-todo-api/handlers"
	"github.com/krishnakumar-learn/go-todo-api/repository"
	"github.com/krishnakumar-learn/go-todo-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	db := db.Connect()
	db.AutoMigrate(&models.Todo{})

	repo := repository.NewMySQLTodoRepository(db)
	handler := handlers.NewTodoHandler(repo)

	r := gin.Default()
	todos := r.Group("/todos")
	{
		todos.GET("/", handler.GetTodos)
		todos.POST("/", handler.CreateTodo)
		todos.GET("/:id", handler.GetTodo)
		todos.PUT("/:id", handler.UpdateTodo)
		todos.DELETE("/:id", handler.DeleteTodo)
	}

	r.Run(":8080")
}
