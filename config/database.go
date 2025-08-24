package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB adalah variabel global untuk menampung koneksi database
var DB *gorm.DB

// ConnectDB adalah fungsi untuk menghubungkan aplikasi ke database
func ConnectDB() {
	var err error
	// Membuat Data Source Name (DSN) dari environment variables
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Membuka koneksi ke database menggunakan GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal terhubung ke database!")
	}

	fmt.Println("Berhasil terhubung ke database!")
	DB = db
}