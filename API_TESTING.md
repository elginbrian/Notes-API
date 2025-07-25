# API Testing Guide

This document provides curl examples for testing the Notes API endpoints.

## Prerequisites

Make sure the API is running with Docker Compose:

```bash
docker-compose up --build
```

The API will be available at `http://localhost:3000`

## Interactive Documentation

For a better testing experience, use the **Swagger UI** at:
**`http://localhost:3000/swagger/`**

The Swagger interface provides:

- Interactive API testing
- Automatic request/response examples
- Built-in authentication handling
- Complete API documentation

## Manual Testing with curl

## Authentication Endpoints

### Register a new user

```bash
curl -X POST http://localhost:3000/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }'
```

### Login

```bash
curl -X POST http://localhost:3000/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

**Note**: Save the JWT token from the login response for subsequent requests.

## Notes Endpoints (Protected)

Replace `YOUR_JWT_TOKEN` with the actual token from login.

### Get all notes

```bash
curl -X GET http://localhost:3000/api/notes \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### Get specific note

```bash
curl -X GET http://localhost:3000/api/notes/NOTE_ID \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### Create note (text only)

```bash
curl -X POST http://localhost:3000/api/notes \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -F "title=My First Note" \
  -F "content=This is the content of my note"
```

### Create note with image

```bash
curl -X POST http://localhost:3000/api/notes \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -F "title=Note with Image" \
  -F "content=This note has an image attached" \
  -F "image=@/path/to/your/image.jpg"
```

### Update note

```bash
curl -X PUT http://localhost:3000/api/notes/NOTE_ID \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -F "title=Updated Title" \
  -F "content=Updated content"
```

### Update note with new image

```bash
curl -X PUT http://localhost:3000/api/notes/NOTE_ID \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -F "title=Updated Title" \
  -F "content=Updated content" \
  -F "image=@/path/to/new/image.jpg"
```

### Delete note

```bash
curl -X DELETE http://localhost:3000/api/notes/NOTE_ID \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## Health Check

### Check API status

```bash
curl -X GET http://localhost:3000/health
```

## Response Examples

### Successful Registration/Login Response

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": "uuid-here",
    "name": "John Doe",
    "email": "john@example.com",
    "created_at": "2025-07-25T10:00:00Z",
    "updated_at": "2025-07-25T10:00:00Z"
  }
}
```

### Notes List Response

```json
{
  "notes": [
    {
      "id": "uuid-here",
      "title": "My Note",
      "content": "Note content",
      "image_path": "uploads/image.jpg",
      "image_url": "http://localhost:3000/uploads/image.jpg",
      "user_id": "uuid-here",
      "created_at": "2025-07-25T10:00:00Z",
      "updated_at": "2025-07-25T10:00:00Z"
    }
  ]
}
```

### Error Response

```json
{
  "error": "Description of the error"
}
```

## File Upload Notes

- Supported image formats: JPEG, PNG, GIF
- Images are stored in the `uploads/` directory
- Each uploaded image gets a unique UUID-based filename
- Image URLs are automatically generated in responses
- Old images are automatically deleted when updating notes with new images
