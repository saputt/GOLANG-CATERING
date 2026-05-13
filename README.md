# 🍱 Catering API

A simple RESTful API for catering services built with **Go (Golang)**. This project serves as a backend to manage food menus, user authentication, and customer testimonials.

---

## 📋 Table of Contents

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Folder Structure](#folder-structure)
- [Getting Started](#getting-started)
- [Environment Configuration](#environment-configuration)
- [API Endpoints](#api-endpoints)

---

## ✨ Features

- **Authentication** — Register & Login using JWT (JSON Web Tokens).
- **Role-based Access Control** — Supports `admin` and `user` roles.
- **Meal Management** — Full CRUD for food menus (Restricted: only `admin` can Create/Update/Delete).
- **Testimonials** — Users can provide, edit, and delete testimonials for specific meals.
- **Auto Migration** — Database schema is automatically generated when the server starts for the first time.

---

## 🛠️ Tech Stack

| Technology | Purpose |
|---|---|
| [Go (Golang)](https://go.dev/) | Primary programming language |
| [Chi](https://github.com/go-chi/chi) | Lightweight and idiomatic HTTP router |
| [GORM](https://gorm.io/) | ORM for database interactions |
| [PostgreSQL](https://www.postgresql.org/) | Relational database |
| [pgx](https://github.com/jackc/pgx) | PostgreSQL driver for Go |
| [golang-jwt/jwt](https://github.com/golang-jwt/jwt) | JSON Web Token creation & validation |
| [google/uuid](https://github.com/google/uuid) | UUID generation for primary keys |
| [godotenv](https://github.com/joho/godotenv) | Loads configuration from `.env` files |
| [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) | Password hashing |

---

## 📁 Folder Structure


```

catering-api/
├── cmd/
│   └── api/
│       └── main.go              # Application entry point
│
├── internal/
│   ├── auth/                    # Authentication module
│   │   ├── handler.go           # HTTP handlers (Register, Login)
│   │   ├── middleware.go        # JWT middleware & AdminOnly guard
│   │   ├── model.go             # Structs: User, LoginRequest, RegisterRequest, etc.
│   │   ├── repository.go        # Database queries for User
│   │   └── service.go           # Authentication business logic
│   │
│   ├── meal/                    # Meal/Menu module
│   │   ├── handler.go           # HTTP handlers for Meal CRUD
│   │   ├── model.go             # Structs: Meal, CreateMealRequest, UpdateMealRequest
│   │   ├── repository.go        # Database queries for Meal
│   │   └── service.go           # Meal business logic
│   │
│   ├── testimonial/             # Testimonial module
│   │   ├── handler.go           # HTTP handlers for Testimonial CRUD
│   │   ├── model.go             # Structs: Testimonial, CreateTestimonialRequest, etc.
│   │   ├── repository.go        # Database queries for Testimonial
│   │   └── service.go           # Testimonial business logic
│   │
│   ├── config/
│   │   └── config.go            # Load configuration from environment variables
│   │
│   ├── database/
│   │   └── postgres.go          # Initialize PostgreSQL connection via GORM
│   │
│   └── httpx/
│       └── response.go          # Helper for consistent HTTP responses (success & error)
│
├── .env                         # Environment configuration (git-ignored)
├── .gitignore
├── go.mod                       # Go module definition & dependencies
└── go.sum                       # Dependency checksums

```

### Module Descriptions

| Folder/File | Function |
|---|---|
| `cmd/api/` | **Entry point**. Initializes config, database, router, and all modules, then starts the HTTP server. |
| `internal/auth/` | Handles all **authentication** logic: registration, login, password hashing, JWT generation, and route protection middleware. |
| `internal/meal/` | Manages **meal menu data** (CRUD). Write operations (create, update, delete) are restricted to `admin` users only. |
| `internal/testimonial/` | Manages **customer testimonials** per meal. Authenticated users can create, update, and delete their own testimonials. |
| `internal/config/` | Reads settings from the `.env` file and provides a `Config` struct for application-wide use. |
| `internal/database/` | Initializes and returns a **GORM + PostgreSQL** connection instance used by all repositories. |
| `internal/httpx/` | Contains `WriteSuccess` and `WriteError` helper functions to ensure all HTTP responses follow a **consistent JSON format**. |

---

## 🚀 Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) version 1.21 or higher
- [PostgreSQL](https://www.postgresql.org/) instance running

### Installation Steps

**1. Clone the repository**
```bash
git clone [https://github.com/username/catering-api.git](https://github.com/username/catering-api.git)
cd catering-api

```

**2. Setup environment variables**

```bash
cp .env.example .env

```

**3. Install dependencies**

```bash
go mod tidy

```

**4. Run the application**

```bash
go run ./cmd/api/main.go

```

The server will start on `http://localhost:3000` (or the `PORT` specified in your `.env`).

> Note: Database migrations run **automatically** when the server starts.

---

## ⚙️ Environment Configuration

Create a `.env` file in the root directory using the following variables:

| Variable | Description | Default |
| --- | --- | --- |
| `PORT` | The port the server listens on | `3000` |
| `DATABASE_URL` | PostgreSQL connection string | *(Required)* |
| `JWT_SECRET` | Secret key for signing JWT tokens | `ANY_SECRET_KEY` |
| `EXPIRES_HOUR` | JWT expiration time (in hours) | `24` |

**Example `.env`:**

```env
PORT=3000
DATABASE_URL=postgres://postgres:password@localhost:5432/catering_db?sslmode=disable
JWT_SECRET=your_super_secret_key
EXPIRES_HOUR=24

```

---

## 📡 API Endpoints

### Auth

| Method | Endpoint | Access | Description |
| --- | --- | --- | --- |
| `POST` | `/auth/register` | Public | Register a new account |
| `POST` | `/auth/login` | Public | Login and receive a JWT token |

### Meal (Menu Management)

> All Meal endpoints require `Authorization: Bearer <token>` header

| Method | Endpoint | Access | Description |
| --- | --- | --- | --- |
| `GET` | `/meals` | User, Admin | Retrieve all meals |
| `GET` | `/meals/{mealId}` | User, Admin | Get details of a single meal |
| `POST` | `/meals` | **Admin only** | Create a new meal |
| `PATCH` | `/meals/{mealId}` | **Admin only** | Update an existing meal |
| `DELETE` | `/meals/{mealId}` | **Admin only** | Remove a meal |

### Testimonials

> All Testimonial endpoints require `Authorization: Bearer <token>` header

| Method | Endpoint | Access | Description |
| --- | --- | --- | --- |
| `GET` | `/testimonials/{mealId}` | User, Admin | Get all testimonials for a specific meal |
| `POST` | `/testimonials` | User, Admin | Post a new testimonial |
| `PATCH` | `/testimonials/{testiId}` | User, Admin | Update your testimonial |
| `DELETE` | `/testimonials/{testiId}` | User, Admin | Delete your testimonial |

---

### Response Examples

**Success:**

```json
{
  "message": "Get all meals success",
  "data": [...]
}

```

**Error:**

```json
{
  "message": "Meal not found"
}

```

---

## 👤 Roles

| Role | Permissions |
| --- | --- |
| `user` | Login, view meals, manage own testimonials. |
| `admin` | Full user permissions + manage (Create/Update/Delete) all meal data. |

---

## 📝 License

This project is created for educational purposes. Feel free to use and modify it as you see fit.
