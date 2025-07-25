package handlers

import (
	"notes-api/database"
	"notes-api/models"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// Register godoc
// @Summary Register a new user
// @Description Create a new user account with name, email, and password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body models.RegisterRequest true "User registration data"
// @Success 201 {object} models.AuthSuccessResponse "User registered successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid request body"
// @Failure 409 {object} models.ErrorResponse "User already exists"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/auth/register [post]
func Register(c *fiber.Ctx) error {
	var req models.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Status: "error",
			Error:  "Invalid request body",
		})
	}

	// Check if user already exists
	var existingUser models.User
	if err := database.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		return c.Status(fiber.StatusConflict).JSON(models.ErrorResponse{
			Status: "error",
			Error:  "User with this email already exists",
		})
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Status: "error",
			Error:  "Failed to hash password",
		})
	}

	// Create user
	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Status: "error",
			Error:  "Failed to create user",
		})
	}

	// Generate JWT token
	token, err := generateJWT(user.ID.String())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Status: "error",
			Error:  "Failed to generate token",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(models.AuthSuccessResponse{
		Status:  "success",
		Message: "User registered successfully",
		Data: models.AuthData{
			Token: token,
			User: models.AuthUser{
				ID:        user.ID,
				Email:     user.Email,
				Name:      user.Name,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
			},
		},
	})
}

// Login godoc
// @Summary User login
// @Description Authenticate user with email and password, returns JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body models.LoginRequest true "User login credentials"
// @Success 200 {object} models.AuthSuccessResponse "Login successful"
// @Failure 400 {object} models.ErrorResponse "Invalid request body"
// @Failure 401 {object} models.ErrorResponse "Invalid credentials"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/auth/login [post]
func Login(c *fiber.Ctx) error {
	var req models.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Status: "error",
			Error:  "Invalid request body",
		})
	}

	// Find user
	var user models.User
	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Status: "error",
			Error:  "Invalid credentials",
		})
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Status: "error",
			Error:  "Invalid credentials",
		})
	}

	// Generate JWT token
	token, err := generateJWT(user.ID.String())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Status: "error",
			Error:  "Failed to generate token",
		})
	}

	return c.JSON(models.AuthSuccessResponse{
		Status:  "success",
		Message: "Login successful",
		Data: models.AuthData{
			Token: token,
			User: models.AuthUser{
				ID:        user.ID,
				Email:     user.Email,
				Name:      user.Name,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
			},
		},
	})
}

func generateJWT(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(), // 7 days
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
