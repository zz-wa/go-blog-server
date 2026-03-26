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
- Role-based access control (Casbin)
- Article management with category & tag associations (many-to-many)
- Article list filtering by status / category / tag / keyword with pagination
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

3. Start PostgreSQL (Docker)

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
| User | `/api/v1/user/` | JWT |
| Admin | `/api/v1/admin/` | JWT + Admin role |

### Key Endpoints

```
POST   /api/v1/public/register
POST   /api/v1/public/login
GET    /api/v1/public/articles
GET    /api/v1/public/articles/:id

GET    /api/v1/user/profile
POST   /api/v1/user/upload

GET    /api/v1/admin/articles
POST   /api/v1/admin/articles
PUT    /api/v1/admin/articles/:id
DELETE /api/v1/admin/articles/:id

GET    /api/v1/admin/userlist
PUT    /api/v1/admin/users/:id
PUT    /api/v1/admin/users/:id/password
PUT    /api/v1/admin/users/:id/status

GET    /api/v1/admin/login-logs
GET    /api/v1/admin/operation-logs
```

## License

[MIT](LICENSE)
