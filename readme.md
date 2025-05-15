# GoGinKickstart

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.23-blue)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

A boilerplate **Go + Gin** backend featuring:

- JWT-based authentication  
- SQLite persistence via GORM  
- Auto-migrated schema & initial admin seeding  
- Swagger UI docs  
- Static HTML login & home pages  
- Ready to integrate with a [Next.js + shadcn](https://github.com/shadcn/shadcn-ui) frontend

## Features

1. **Authentication**  
   - `/api/v1/auth/login` issues JWTs (see [`controllers/auth_controller.go`](controllers/auth_controller.go))  
   - Middleware [`middlewares/AuthMiddleware`](middlewares/auth_middleware.go) protects routes

2. **Ping Endpoint**  
   - `/api/v1/ping` returns `{ message: "pong", user: "<username>" }` if authenticated  
   - Handler in [`controllers/ping_controller.go`](controllers/ping_controller.go)

3. **Database**  
   - SQLite (default `DB_PATH=./myproject.db`)  
   - GORM auto-migrate & global `models.DB` in [`models/db.go`](models/db.go)  
   - `services.UserService` seeds an initial admin from `ADMIN_PASSWORD`  

4. **API Docs**  
   - Swagger JSON/YAML in [docs/swagger.json](docs/swagger.json) & [docs/swagger.yaml](docs/swagger.yaml)  
   - UI at `http://localhost:8080/swagger/index.html` via [`routes/swagger_routes.go`](routes/swagger_routes.go)

## Getting Started

### Prerequisites

- Go ≥1.23  
- SQLite3 (optional; bundled via `github.com/mattn/go-sqlite3`)  
- [`swag`](https://github.com/swaggo/swag#installation) for doc generation  

### Installation

```bash
git clone https://github.com/your-org/GoGinKickstart.git
cd GoGinKickstart
go mod tidy