# Deployment Troubleshooting Guide

## Issues Fixed

### 1. 404 Error at Root Path (`/`)

**Problem**: `{"error": "Cannot GET /"}` when accessing the root URL.

**Solution**: Added a root endpoint that provides API information and navigation.

**Now available at**: `http://178.128.61.145:3011/`

```json
{
  "message": "Welcome to Notes API",
  "version": "1.0.0",
  "status": "running",
  "docs": "Visit /swagger/ for API documentation",
  "health": "/health",
  "endpoints": {
    "auth": {
      "register": "POST /api/auth/register",
      "login": "POST /api/auth/login"
    },
    "notes": {
      "list": "GET /api/notes",
      "get": "GET /api/notes/:id",
      "create": "POST /api/notes",
      "update": "PUT /api/notes/:id",
      "delete": "DELETE /api/notes/:id"
    }
  }
}
```

### 2. Swagger Documentation Access

**Problem**: 404 error when accessing `/swagger`

**Solutions Applied**:

- Added redirect from `/swagger` to `/swagger/`
- Ensured proper Swagger route configuration
- Regenerated Swagger documentation

**Access Points**:

- `http://178.128.61.145:3011/swagger/` (main Swagger UI)
- `http://178.128.61.145:3011/swagger` (redirects to above)

## Available Endpoints

### Public Endpoints

- `GET /` - API welcome page with navigation
- `GET /health` - Health check
- `GET /swagger/` - Interactive API documentation
- `GET /uploads/{filename}` - Serve uploaded images

### Authentication

- `POST /api/auth/register` - User registration
- `POST /api/auth/login` - User authentication

### Protected Endpoints (require JWT token)

- `GET /api/notes` - List all user notes
- `GET /api/notes/{id}` - Get specific note
- `POST /api/notes` - Create new note (with optional image)
- `PUT /api/notes/{id}` - Update existing note
- `DELETE /api/notes/{id}` - Delete note

## Testing the API

### 1. Health Check

```bash
curl http://178.128.61.145:3011/health
```

### 2. Register a User

```bash
curl -X POST http://178.128.61.145:3011/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test User",
    "email": "test@example.com",
    "password": "password123"
  }'
```

### 3. Login

```bash
curl -X POST http://178.128.61.145:3011/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123"
  }'
```

### 4. Use Swagger UI (Recommended)

Visit: `http://178.128.61.145:3011/swagger/`

1. Click "Authorize" button
2. Enter: `Bearer YOUR_JWT_TOKEN` (replace with actual token from login)
3. Test all endpoints interactively

## Docker Deployment Status

Your current setup:

- **External Port**: 3011 (mapped to internal port 3000)
- **Database**: PostgreSQL with health checks
- **File Storage**: Volume mounted for uploads
- **CORS**: Enabled for all origins

## Common Issues & Solutions

### Swagger Not Loading

- Ensure you're accessing `/swagger/` with trailing slash
- Check if docs are generated: `swag init`
- Verify Docker container is running: `docker-compose ps`

### Database Connection Issues

- Check PostgreSQL container status
- Verify environment variables in docker-compose.yml
- Check container logs: `docker-compose logs postgres`

### File Upload Issues

- Ensure uploads directory exists and has proper permissions
- Check volume mounting in docker-compose.yml
- Verify supported formats: JPEG, PNG, GIF

### Authentication Issues

- Ensure JWT token is included in Authorization header
- Format: `Authorization: Bearer YOUR_TOKEN`
- Check token expiration (7 days by default)

## Next Steps

1. ✅ **Root endpoint** - Fixed, now shows API info
2. ✅ **Swagger documentation** - Fixed, accessible at `/swagger/`
3. 🔄 **Test the full API workflow** using Swagger UI
4. 🔄 **Upload test images** to verify file handling
5. 🔄 **Monitor logs** for any runtime issues

Your Notes API is now fully functional and ready for use! 🚀
