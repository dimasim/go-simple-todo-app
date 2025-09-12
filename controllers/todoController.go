package controllers

import (
	"net/http"

	"github.com/dimasim/go-simple-todo-app/config"
	"github.com/dimasim/go-simple-todo-app/models"
	"github.com/gin-gonic/gin"
)

// GetAllTodos mengambil semua data todo
func GetAllTodos(c *gin.Context) {
	var todos []models.Todo
	// Mengambil semua record dari tabel todos
	config.DB.Find(&todos)

	c.JSON(http.StatusOK, gin.H{"data": todos})
}

// CreateTodo membuat todo baru
func CreateTodo(c *gin.Context) {
	var todo models.Todo

	// Binding JSON request body ke struct Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Menyimpan data baru ke database
	config.DB.Create(&todo)

	c.JSON(http.StatusCreated, gin.H{"data": todo})
}