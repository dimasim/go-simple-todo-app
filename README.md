# Go Simple Todo App 🚀

Aplikasi backend Todo sederhana namun andal yang dibangun menggunakan **Go (Golang)**, **Gin Web Framework**, dan **GORM** dengan database **PostgreSQL**. Aplikasi ini mendukung autentikasi berbasis JWT (JSON Web Token) dan unggah gambar untuk setiap item todo.

---

## 🛠️ Tech Stack & Fitur

- **Go (Golang)** - Bahasa pemrograman utama yang cepat dan efisien.
- **Gin Gonic** - Framework HTTP yang berkinerja tinggi untuk membangun RESTful API.
- **GORM** - ORM Go untuk pengelolaan database PostgreSQL yang mudah.
- **JWT (JSON Web Token)** - Digunakan untuk mengamankan API (autentikasi dan otorisasi).
- **Bcrypt** - Digunakan untuk hashing password sebelum disimpan ke database.
- **Docker & Docker Compose** - Memudahkan deployment dan penyetelan database PostgreSQL.
- **User Isolation** - Keamanan penuh di mana user hanya dapat melihat, membuat, memperbarui, menghapus, atau mengunggah gambar pada todo milik mereka sendiri.

---

## 📂 Struktur Proyek

```text
├── config/              # Konfigurasi aplikasi (koneksi database)
├── controllers/         # Logika penanganan request HTTP (User & Todo)
├── middlewares/         # Middleware penanganan autentikasi JWT
├── models/              # Struktur data GORM / representasi tabel database
├── public/              # Direktori file statis (seperti gambar todo yang diunggah)
├── .env                 # File konfigurasi environment variables (lokal)
├── Dockerfile           # Docker configuration file untuk deployment production
├── docker-compose.yml   # Docker Compose untuk database PostgreSQL
├── go.mod               # Dependency tracking Go
├── main.go              # Entry point utama aplikasi
└── Makefile             # Shortcut command helper (run, build, clean)
```

---

## 🚀 Persyaratan Sistem

Pastikan Anda telah menginstal software berikut di komputer Anda:
- [Go (Golang)](https://go.dev/doc/install) (versi 1.21 ke atas direkomendasikan)
- [Docker](https://www.docker.com/products/docker-desktop) & Docker Compose
- Client database (misal: DBeaver, pgAdmin) *opsional*

---

## ⚙️ Cara Memulai & Menjalankan Aplikasi

### 1. Kloning Repositori
```bash
git clone https://github.com/dimasim/go-simple-todo-app.git
cd go-simple-todo-app
```

### 2. Konfigurasi Environment Variables
Salin atau buat file `.env` di root direktori dengan contoh konfigurasi berikut:
```env
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=gotodolist
JWT_SECRET=kunciRahasiaSuperAmanYangPanjangSekali123!@#
```

### 3. Jalankan Database PostgreSQL via Docker Compose
Jalankan perintah berikut untuk mengaktifkan database PostgreSQL dalam container Docker:
```bash
docker-compose up -d
```
*Database akan berjalan pada port `5432` secara default.*

### 4. Menjalankan Aplikasi Secara Lokal
Pastikan dependencies sudah terunduh, lalu jalankan aplikasinya:
```bash
go mod download
go run main.go
```
Aplikasi akan secara otomatis mendeteksi koneksi database, melakukan auto-migration tabel `users` dan `todos`, lalu berjalan di `http://localhost:8080`.

---

## 🐳 Menjalankan Menggunakan Docker (Opsional)

Jika ingin menjalankan seluruh aplikasi di dalam kontainer Docker, gunakan Dockerfile multi-stage yang telah disediakan:

```bash
# Build Docker image
docker build -t go-todo-app .

# Jalankan kontainer
docker run -d -p 8080:8080 --name go-todo-container --env-file .env go-todo-app
```

---

## 📋 Dokumentasi API Endpoints

Semua request berformat JSON dan mengembalikan respons JSON. Endpoint di bawah `/api/todos` memerlukan header `Authorization: Bearer <token_jwt>`.

### 🔐 Autentikasi Pengguna

#### 1. Registrasi Akun Baru
- **URL**: `/api/register`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "name": "Dimas",
    "email": "dimas@example.com",
    "password": "password123"
  }
  ```

#### 2. Login Pengguna
- **URL**: `/api/login`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "email": "dimas@example.com",
    "password": "password123"
  }
  ```
- **Response**: Mengembalikan JWT Token yang harus digunakan pada request berikutnya.
  ```json
  {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
  ```

---

### 📝 Manajemen Item Todo (Memerlukan Autentikasi JWT)

Header wajib: `Authorization: Bearer <JWT_TOKEN>`

#### 1. Mengambil Semua Todo
- **URL**: `/api/todos`
- **Method**: `GET`
- **Response**: Mengembalikan daftar todo milik user yang sedang login.

#### 2. Membuat Todo Baru
- **URL**: `/api/todos`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "title": "Membeli Susu",
    "description": "Beli susu cair 1 liter di toko terdekat"
  }
  ```

#### 3. Mengambil Detail Todo Berdasarkan ID
- **URL**: `/api/todos/:id`
- **Method**: `GET`

#### 4. Memperbarui Todo
- **URL**: `/api/todos/:id`
- **Method**: `PUT`
- **Request Body** (Semua field bersifat opsional):
  ```json
  {
    "title": "Membeli Susu & Roti",
    "is_done": true
  }
  ```

#### 5. Menghapus Todo
- **URL**: `/api/todos/:id`
- **Method**: `DELETE`

#### 6. Unggah Gambar Pendukung Todo
- **URL**: `/api/todos/:id/upload`
- **Method**: `POST`
- **Content-Type**: `multipart/form-data`
- **Form Data**:
  - `image`: *[File Gambar]*
