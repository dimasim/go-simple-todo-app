package controllers

import (
	"net/http"
	"path/filepath"

	"github.com/dimasim/go-simple-todo-app/config"
	"github.com/dimasim/go-simple-todo-app/models"
	"github.com/gin-gonic/gin"
)

// GetAllTodos mengambil semua data todo milik user yang login
func GetAllTodos(c *gin.Context) {
	var todos []models.Todo
	user, _ := c.Get("user")

	config.DB.Where("user_id = ?", user.(models.User).ID).Find(&todos)
	c.JSON(http.StatusOK, gin.H{"data": todos})
}

// CreateTodo membuat todo baru untuk user yang login
func CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, _ := c.Get("user")
	todo.UserID = user.(models.User).ID

	if err := config.DB.Create(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan todo ke database"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": todo})
}

// GetTodoByID mengambil todo berdasarkan ID, khusus untuk milik user yang login
func GetTodoByID(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")
	user, _ := c.Get("user")

	if err := config.DB.Where("id = ? AND user_id = ?", id, user.(models.User).ID).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo tidak ditemukan atau Anda tidak memiliki akses"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

// UpdateTodo memperbarui todo milik user yang login
func UpdateTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")
	user, _ := c.Get("user")

	// Pastikan todo milik user yang bersangkutan
	if err := config.DB.Where("id = ? AND user_id = ?", id, user.(models.User).ID).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo tidak ditemukan atau Anda tidak memiliki akses"})
		return
	}

	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		IsDone      *bool  `json:"is_done"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := make(map[string]interface{})
	if input.Title != "" {
		updates["title"] = input.Title
	}
	if input.Description != "" {
		updates["description"] = input.Description
	}
	if input.IsDone != nil {
		updates["is_done"] = *input.IsDone
	}

	if err := config.DB.Model(&todo).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

// DeleteTodo menghapus todo milik user yang login
func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	user, _ := c.Get("user")

	var todo models.Todo
	if err := config.DB.Where("id = ? AND user_id = ?", id, user.(models.User).ID).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo tidak ditemukan atau Anda tidak memiliki akses"})
		return
	}

	if err := config.DB.Delete(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true, "message": "Todo berhasil dihapus"})
}

// UploadTodoImage mengunggah gambar untuk todo milik user yang login
func UploadTodoImage(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")
	user, _ := c.Get("user")

	if err := config.DB.Where("id = ? AND user_id = ?", id, user.(models.User).ID).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo tidak ditemukan atau Anda tidak memiliki akses"})
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tidak ada file gambar yang diupload"})
		return
	}

	filename := id + filepath.Ext(file.Filename)
	path := "public/uploads/" + filename

	if err := c.SaveUploadedFile(file, path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan file"})
		return
	}

	if err := config.DB.Model(&todo).Update("ImageURL", path).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui Image URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}
