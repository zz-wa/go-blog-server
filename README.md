# go-blog-server

A RESTful blog backend built with Go, featuring JWT authentication, RBAC authorization, and a clean layered architecture.

## Tech Stack

- **Framework**: [Echo v5](https://echo.labstack.com/)
- **ORM**: [GORM](https://gorm.io/) + PostgreSQL / SQLite
- **Auth**: JWT ([golang-jwt/jwt v5](https://github.com/golang-jwt/jwt))
- **RBAC**: [Casbin v3](https://casbin.org/)
- **Cache / Rate Limit**: [Redis](https://redis.io/) (go-redis v9)
- **Config**: [Viper](https://github.com/spf13/viper)
- **Logging**: [Zap](https://github.com/uber-go/zap)

## Features

- User registration & login with JWT
- Role-based access control (Casbin): visitor / user / admin
- Article management with category & tag associations (many-to-many)
- Article list filtering by status / category / tag / keyword with pagination
- Article archive (grouped by year/month)
- Article view count tracking
- Comment system (users post comments, admin deletes)
- Like system for articles and comments (Redis Set + DB persistence)
- Admin dashboard statistics (article count, views, categories, tags, users)
- Public home page aggregation endpoint
- System config management (key-value store, includes "about" page)
- Admin user management (list, update, reset password, enable/disable)
- Role & menu management
- File upload (local storage)
- Login log & operation log
- Redis-based login rate limiting (5 req/min per IP)
- Admin auto-initialization from config

## Project Structure

```
.
├── main.go
├── internal/
│   ├── api/          # HTTP handlers
│   ├── service/      # Business logic
│   ├── repository/   # Database access
│   ├── model/        # GORM models
│   ├── middleware/   # Auth, permission, logging, rate limit
│   ├── router/       # Route definitions
│   ├── pkg/jwt/      # JWT helpers
│   ├── global/       # Global variables & config
│   └── casbin/       # RBAC model & policy
└── uploads/          # Uploaded files
```

## Getting Started

### Prerequisites

- Go 1.21+
- PostgreSQL (or use SQLite for local dev)
- Redis

### Setup

1. Clone the repo

```bash
git clone https://github.com/butter-July/go-blog-server.git
cd go-blog-server
```

2. Copy the example config and fill in your values

```bash
cp internal/global/config.yaml.example internal/global/config.yaml
```

Edit `config.yaml`:
- Set your database credentials under `Pgsql`
- Set a strong `JWT.Secret`
- Set `Admin.Email` and `Admin.Password` for the initial admin account

3. Start PostgreSQL and Redis (Docker)

```bash
docker compose up -d
```

4. Run

```bash
go run main.go
```

The server starts on port `8080` by default.

## API Overview

| Group | Prefix | Auth Required |
|-------|--------|--------------|
| Public | `/api/v1/public/` | No |
| User | `/api/v1/user/` | JWT (any logged-in user) |
| Admin | `/api/v1/admin/` | JWT + Admin role |

### Key Endpoints

```
# Public
POST   /api/v1/public/register
POST   /api/v1/public/login
GET    /api/v1/public/home
GET    /api/v1/public/articles
GET    /api/v1/public/articles/:id
GET    /api/v1/public/archive
GET    /api/v1/public/commentList/:article_id
GET    /api/v1/public/about

# User (login required)
POST   /api/v1/user/comment/:article_id
POST   /api/v1/user/like/:like_type/:target_id

# Admin
GET    /api/v1/admin/articles
POST   /api/v1/admin/articles
PUT    /api/v1/admin/articles/:id
DELETE /api/v1/admin/articles/:id
DELETE /api/v1/admin/comment/:id

GET    /api/v1/admin/dashboard
GET    /api/v1/admin/configs
PUT    /api/v1/admin/configs/:key
PUT    /api/v1/admin/about

GET    /api/v1/admin/userlist
PUT    /api/v1/admin/users/:id
PUT    /api/v1/admin/users/:id/password
PUT    /api/v1/admin/users/:id/status

GET    /api/v1/admin/login-logs
GET    /api/v1/admin/operation-logs
```

## Reference

This project was built as a learning exercise, referencing the architecture and feature set of [Gin-Vue-Admin/gin-vue-blog](https://github.com/Tjyy-1223/Gin-Vue-Admin) (Gin + Vue full-stack blog). Rewritten from scratch using Echo v5 instead of Gin, with a simplified feature set focused on the core backend.

## License

[MIT](LICENSE)
