# 🍱 Catering API

RESTful API sederhana untuk layanan katering berbasis **Go (Golang)**. Proyek ini dibangun sebagai backend untuk mengelola menu makanan, autentikasi pengguna, dan testimoni pelanggan.

---

## 📋 Daftar Isi

- [Fitur](#fitur)
- [Teknologi yang Digunakan](#teknologi-yang-digunakan)
- [Struktur Folder](#struktur-folder)
- [Cara Menjalankan](#cara-menjalankan)
- [Konfigurasi Environment](#konfigurasi-environment)
- [API Endpoints](#api-endpoints)

---

## ✨ Fitur

- **Autentikasi** — Register & Login menggunakan JWT
- **Role-based Access Control** — Role `admin` dan `user`
- **Manajemen Menu (Meal)** — CRUD menu makanan, hanya `admin` yang bisa menambah/edit/hapus
- **Testimoni** — Pengguna bisa memberikan, mengubah, dan menghapus testimoni per menu
- **Auto Migration** — Skema database otomatis dibuat saat server pertama kali dijalankan

---

## 🛠️ Teknologi yang Digunakan

| Teknologi | Kegunaan |
|---|---|
| [Go (Golang)](https://go.dev/) | Bahasa pemrograman utama |
| [Chi](https://github.com/go-chi/chi) | HTTP router yang ringan dan idiomatic |
| [GORM](https://gorm.io/) | ORM untuk interaksi dengan database |
| [PostgreSQL](https://www.postgresql.org/) | Database relasional |
| [pgx](https://github.com/jackc/pgx) | PostgreSQL driver untuk Go |
| [golang-jwt/jwt](https://github.com/golang-jwt/jwt) | Pembuatan & validasi JSON Web Token |
| [google/uuid](https://github.com/google/uuid) | Generate UUID untuk primary key |
| [godotenv](https://github.com/joho/godotenv) | Load konfigurasi dari file `.env` |
| [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) | Hashing password |

---

## 📁 Struktur Folder

```
catering-api/
├── cmd/
│   └── api/
│       └── main.go              # Entry point aplikasi
│
├── internal/
│   ├── auth/                    # Modul autentikasi
│   │   ├── handler.go           # HTTP handler (Register, Login)
│   │   ├── middleware.go        # JWT middleware & AdminOnly guard
│   │   ├── model.go             # Struct: User, LoginRequest, RegisterRequest, dsb.
│   │   ├── repository.go        # Query database untuk User
│   │   └── service.go           # Business logic autentikasi
│   │
│   ├── meal/                    # Modul menu makanan
│   │   ├── handler.go           # HTTP handler CRUD Meal
│   │   ├── model.go             # Struct: Meal, CreateMealRequest, UpdateMealRequest
│   │   ├── repository.go        # Query database untuk Meal
│   │   └── service.go           # Business logic Meal
│   │
│   ├── testimonial/             # Modul testimoni
│   │   ├── handler.go           # HTTP handler CRUD Testimonial
│   │   ├── model.go             # Struct: Testimonial, CreateTestimonialRequest, dsb.
│   │   ├── repository.go        # Query database untuk Testimonial
│   │   └── service.go           # Business logic Testimonial
│   │
│   ├── config/
│   │   └── config.go            # Load konfigurasi dari environment variables
│   │
│   ├── database/
│   │   └── postgres.go          # Inisialisasi koneksi ke PostgreSQL via GORM
│   │
│   └── httpx/
│       └── response.go          # Helper untuk menulis HTTP response (success & error)
│
├── .env                         # Konfigurasi environment (tidak di-commit)
├── .gitignore
├── go.mod                       # Definisi module & dependency Go
└── go.sum                       # Checksum dependency
```

### Penjelasan Fungsi Setiap Folder

| Folder/File | Fungsi |
|---|---|
| `cmd/api/` | **Entry point** aplikasi. Menginisialisasi config, database, router, dan semua modul, lalu menjalankan HTTP server. |
| `internal/auth/` | Menangani semua hal terkait **autentikasi**: register, login, hashing password, generate JWT, dan middleware proteksi route. |
| `internal/meal/` | Mengelola **data menu makanan** (CRUD). Operasi tulis (create, update, delete) hanya bisa dilakukan oleh `admin`. |
| `internal/testimonial/` | Mengelola **testimoni pelanggan** per menu. Setiap user yang sudah login bisa membuat, mengubah, dan menghapus testimoni mereka. |
| `internal/config/` | Membaca konfigurasi dari file `.env` dan menyediakan satu struct `Config` yang dipakai oleh seluruh aplikasi. |
| `internal/database/` | Menginisialisasi dan mengembalikan instance koneksi **GORM + PostgreSQL** yang siap dipakai oleh semua repository. |
| `internal/httpx/` | Berisi helper function `WriteSuccess` dan `WriteError` untuk memastikan semua response HTTP memiliki **format JSON yang konsisten**. |

---

## 🚀 Cara Menjalankan

### Prasyarat

- [Go](https://go.dev/dl/) versi 1.21+
- [PostgreSQL](https://www.postgresql.org/) yang sudah berjalan

### Langkah-langkah

**1. Clone repository ini**
```bash
git clone https://github.com/username/catering-api.git
cd catering-api
```

**2. Salin file environment dan isi konfigurasinya**
```bash
cp .env.example .env
```

**3. Install semua dependency**
```bash
go mod tidy
```

**4. Jalankan aplikasi**
```bash
go run ./cmd/api/main.go
```

Server akan berjalan di `http://localhost:3000` (atau sesuai `PORT` di `.env`).

> Migrasi database dijalankan **secara otomatis** saat server pertama kali start.

---

## ⚙️ Konfigurasi Environment

Buat file `.env` di root project berdasarkan tabel berikut:

| Variable | Keterangan | Default |
|---|---|---|
| `PORT` | Port server berjalan | `3000` |
| `DATABASE_URL` | Connection string PostgreSQL | _(wajib diisi)_ |
| `JWT_SECRET` | Secret key untuk signing JWT | `SAUKIGANTENG` |
| `EXPIRES_HOUR` | Durasi JWT token kadaluarsa (jam) | `24` |

**Contoh `.env`:**
```env
PORT=3000
DATABASE_URL=postgres://postgres:password@localhost:5432/catering_db?sslmode=disable
JWT_SECRET=your_super_secret_key
EXPIRES_HOUR=24
```

---

## 📡 API Endpoints

### Auth

| Method | Endpoint | Akses | Deskripsi |
|---|---|---|---|
| `POST` | `/auth/register` | Public | Daftar akun baru |
| `POST` | `/auth/login` | Public | Login dan dapatkan JWT token |

### Meal (Menu Makanan)

> Semua endpoint Meal membutuhkan header `Authorization: Bearer <token>`

| Method | Endpoint | Akses | Deskripsi |
|---|---|---|---|
| `GET` | `/meals` | User, Admin | Ambil semua menu |
| `GET` | `/meals/{mealId}` | User, Admin | Ambil detail satu menu |
| `POST` | `/meals` | **Admin only** | Tambah menu baru |
| `PATCH` | `/meals/{mealId}` | **Admin only** | Update menu |
| `DELETE` | `/meals/{mealId}` | **Admin only** | Hapus menu |

### Testimonial

> Semua endpoint Testimonial membutuhkan header `Authorization: Bearer <token>`

| Method | Endpoint | Akses | Deskripsi |
|---|---|---|---|
| `GET` | `/testimonials/{mealId}` | User, Admin | Ambil semua testimoni dari satu menu |
| `POST` | `/testimonials` | User, Admin | Tambah testimoni |
| `PATCH` | `/testimonials/{testiId}` | User, Admin | Update testimoni |
| `DELETE` | `/testimonials/{testiId}` | User, Admin | Hapus testimoni |

---

### Contoh Response

**Success:**
```json
{
  "status": "success",
  "message": "Get all meals success",
  "data": [...]
}
```

**Error:**
```json
{
  "status": "error",
  "message": "Meal not found"
}
```

---

## 👤 Role

| Role | Kemampuan |
|---|---|
| `user` | Login, melihat menu, membuat/edit/hapus testimoni sendiri |
| `admin` | Semua kemampuan `user` + bisa mengelola (tambah/edit/hapus) menu |

---

## 📝 Lisensi

Proyek ini dibuat untuk keperluan belajar. Bebas digunakan dan dimodifikasi.
