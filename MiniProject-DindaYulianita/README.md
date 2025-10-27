# MiniProject-DindaYulianita

Skeleton REST API in Go (Clean Architecture).

## Features included
- Auth (register/login) with JWT. Register auto-creates a store.
- User, Store, Address, Category (admin-only), Product, Transaction services.
- File upload endpoint for product images.
- GORM (MySQL) + auto-migrations.
- Gorilla/mux router.
- Pagination & basic filtering.
- Clean project layout: cmd, internal.

## Quick start
1. Copy `.env.example` to `.env` and edit DB credentials and JWT secret.
2. `go mod tidy`
3. Start MySQL and ensure the database exists.
4. `go run ./cmd`
5. Server listens on port from `.env` (default 8080).

