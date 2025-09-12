package main

import (
	"log"

	"github.com/dimasim/go-simple-todo-app/models"
	"github.com/dimasim/go-simple-todo-app/config"
	"github.com/dimasim/go-simple-todo-app/controllers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

)

// init() akan berjalan sebelum fungsi main()
func init() {
	// Memuat file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error memuat file .env")
	}

	// Menghubungkan ke database
	config.ConnectDB()
	config.DB.AutoMigrate(&models.Todo{}) 
	log.Println("Database Migration successful!")
}

func main() {
	r := gin.Default()

	// Membuat grup rute untuk API
	api := r.Group("/api")
	{
		api.GET("/todos", controllers.GetAllTodos)
		api.POST("/todos", controllers.CreateTodo)
		// Rute lain akan ditambahkan di sini nanti
	}

	// Menjalankan server di port 8080
	r.Run(":8080")
}