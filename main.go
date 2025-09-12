package main

import (
	"log"
	"os"
	"github.com/dimasim/go-simple-todo-app/middlewares"
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
	config.DB.AutoMigrate(&models.User{}, &models.Todo{})
	log.Println("Database Migration successful!")
}

func main() {
	r := gin.Default()

	// Membuat grup rute untuk API
	api := r.Group("/api")
	{
		// api.GET("/todos", controllers.GetAllTodos)
		// api.POST("/todos", controllers.CreateTodo)
		// api.GET("/todos/:id", controllers.GetTodoByID)
		// api.PUT("/todos/:id", controllers.UpdateTodo)
		// api.DELETE("/todos/:id", controllers.DeleteTodo)	
		// api.POST("/todos/:id/upload", controllers.UploadTodoImage)
		// Rute untuk User
		api.POST("/register", controllers.Register) 
		api.POST("/login", controllers.Login)

		// Rute untuk Todos
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

	r.Run(":" + port)

	// r.Run(":8080")
}