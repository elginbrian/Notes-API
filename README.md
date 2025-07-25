# Notes API

A RESTful API for managing notes built with Go Fiber, featuring JWT authentication, image uploads, and PostgreSQL database.

## Features

- CRUD operations for notes
- JWT-based authentication
- Image upload with multipart form data
- PostgreSQL database with GORM
- Docker Compose for easy deployment
- Environment-based configuration
- **Swagger/OpenAPI documentation**

## Quick Start

1. Clone the repository
2. Copy `.env.example` to `.env` and configure your environment variables
3. Run with Docker Compose:

```bash
docker-compose up --build
```

The API will be available at `http://localhost:3000`

## Documentation

### Swagger UI

Access the interactive API documentation at: `http://localhost:3000/swagger/`

The Swagger documentation provides:

- Complete API endpoint documentation
- Interactive testing interface
- Request/response examples
- Authentication setup instructions

## API Endpoints

### Authentication

- `POST /api/auth/register` - Register a new user
- `POST /api/auth/login` - Login user

### Notes (Protected routes)

- `GET /api/notes` - Get all notes for authenticated user
- `GET /api/notes/:id` - Get specific note
- `POST /api/notes` - Create new note (supports image upload)
- `PUT /api/notes/:id` - Update note
- `DELETE /api/notes/:id` - Delete note

### Image Upload

Notes can include images by sending multipart form data with an `image` field.

## Environment Variables

- `DB_HOST` - Database host
- `DB_USER` - Database username
- `DB_PASSWORD` - Database password
- `DB_NAME` - Database name
- `DB_PORT` - Database port
- `JWT_SECRET` - JWT signing secret
- `PORT` - Application port

## Development

To run locally without Docker:

1. Install dependencies: `go mod tidy`
2. Set up PostgreSQL database
3. Configure `.env` file
4. Run: `go run main.go`
