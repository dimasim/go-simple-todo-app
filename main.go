package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/dimasim/go-simple-todo-app/config"// <-- Ganti dengan path modul Anda
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
}

func main() {
	fmt.Println("Aplikasi berjalan...")
}