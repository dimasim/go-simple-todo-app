package models

import "gorm.io/gorm"

// Todo adalah model untuk tabel todos di database
type Todo struct {
	gorm.Model         // Otomatis menambahkan ID, CreatedAt, UpdatedAt, DeletedAt
	Title       string `gorm:"not null"`
	Description string
	IsDone      bool   `gorm:"default:false"`
	ImageURL    string // Kolom untuk menyimpan URL gambar
}