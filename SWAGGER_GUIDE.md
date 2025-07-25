# Swagger Documentation Generation

This document explains how to maintain and update the Swagger documentation for the Notes API.

## Prerequisites

Install the Swag CLI tool:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

## Generating Documentation

After adding or modifying Swagger annotations in your code, regenerate the documentation:

```bash
swag init
```

This will generate/update the following files:

- `docs/docs.go`
- `docs/swagger.json`
- `docs/swagger.yaml`

## Swagger Annotations Reference

### General API Info (main.go)

```go
// @title Notes API
// @version 1.0
// @description A comprehensive Notes API built with Go Fiber
// @host localhost:3000
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
```

### Function Documentation

```go
// FunctionName godoc
// @Summary Brief description
// @Description Detailed description
// @Tags Category
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Parameter description"
// @Success 200 {object} models.ModelName "Success description"
// @Failure 400 {object} map[string]string "Error description"
// @Router /api/endpoint [method]
```

### Common Parameters

- `@Param id path string true "Parameter description"` - Path parameter
- `@Param request body models.Model true "Request body"` - JSON body
- `@Param title formData string true "Form field"` - Form data
- `@Param image formData file false "File upload"` - File upload

### Security

For protected endpoints, add:

```go
// @Security BearerAuth
```

## Accessing Documentation

Once the application is running, access the Swagger UI at:
`http://localhost:3000/swagger/`

## Tips

1. Always regenerate docs after modifying annotations
2. Use descriptive summaries and descriptions
3. Include all possible response codes
4. Document all parameters clearly
5. Use appropriate tags to group related endpoints
