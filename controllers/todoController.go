package controllers

import (
	"net/http"

	"github.com/dimasim/go-simple-todo-app/config"
	"github.com/dimasim/go-simple-todo-app/models"
	"github.com/gin-gonic/gin"
	"path/filepath"
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
func GetTodoByID(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id") // Mengambil ID dari URL parameter

	// Mencari todo pertama yang cocok dengan ID
	if err := config.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo tidak ditemukan!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}
func UpdateTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")

	// 1. Cari dulu todo yang mau diupdate
	if err := config.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo tidak ditemukan!"})
		return
	}

	// 2. Bind data JSON baru ke struct yang sudah ditemukan
	var input models.Todo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 3. Simpan perubahan ke database
	config.DB.Model(&todo).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": todo})
}
func DeleteTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")

	// Mencari todo yang akan dihapus
	if err := config.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo tidak ditemukan!"})
		return
	}

	// Menghapus data dari database
	config.DB.Delete(&todo)

	c.JSON(http.StatusOK, gin.H{"data": true, "message": "Todo berhasil dihapus"})
}
func UploadTodoImage(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")

	// 1. Cari dulu todo yang mau di-update
	if err := config.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo tidak ditemukan!"})
		return
	}

	// 2. Ambil file dari request form-data
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tidak ada file gambar yang diupload"})
		return
	}

	// 3. Buat nama file yang unik untuk menghindari konflik
	//    Contoh sederhana: "id-todo" + "ekstensi-file-asli" -> "1.png"
	filename := id + filepath.Ext(file.Filename)
	path := "public/uploads/" + filename

	// 4. Simpan file ke path yang sudah ditentukan
	if err := c.SaveUploadedFile(file, path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan file"})
		return
	}

	// 5. Update kolom ImageURL di database
	config.DB.Model(&todo).Update("ImageURL", path)

	c.JSON(http.StatusOK, gin.H{"data": todo})
}
