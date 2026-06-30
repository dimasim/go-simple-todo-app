package main

import (
	"log"
	"os"

	"github.com/dimasim/go-simple-todo-app/config"
	"github.com/dimasim/go-simple-todo-app/controllers"
	"github.com/dimasim/go-simple-todo-app/middlewares"
	"github.com/dimasim/go-simple-todo-app/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Info: File .env tidak ditemukan, menggunakan environment variables sistem")
	}
	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{}, &models.Todo{})
	log.Println("Database Migration successful!")
}

func main() {
	r := gin.Default()

	// 1. Definisikan dan gunakan Middleware CORS (SEBELUM RUTE)
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"} // Izinkan React App
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	r.Use(cors.New(corsConfig))

	// 2. Sajikan File Statis
	r.Static("/public", "./public")

	// 3. Definisikan Rute-Rute API Anda
	api := r.Group("/api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		todos := api.Group("/todos")
		todos.Use(middlewares.RequireAuth)
		{
			todos.GET("", controllers.GetAllTodos)
			todos.POST("", controllers.CreateTodo)
			todos.GET("/:id", controllers.GetTodoByID)
			todos.PUT("/:id", controllers.UpdateTodo)
			todos.DELETE("/:id", controllers.DeleteTodo)
			todos.POST("/:id/upload", controllers.UploadTodoImage)
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)
	r.Run(":" + port)
}