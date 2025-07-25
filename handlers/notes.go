package handlers

import (
	"fmt"
	"notes-api/database"
	"notes-api/middleware"
	"notes-api/models"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetNotes godoc
// @Summary Get all notes for authenticated user
// @Description Retrieve all notes belonging to the authenticated user with image URLs
// @Tags Notes
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.NotesResponse "List of notes" 
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/notes [get]
func GetNotes(c *fiber.Ctx) error {
	userID := middleware.GetUserID(c)

	var notes []models.Note
	if err := database.DB.Where("user_id = ?", userID).Find(&notes).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch notes",
		})
	}

	for i := range notes {
		if notes[i].ImagePath != "" {
			notes[i].ImageURL = fmt.Sprintf("%s://%s/uploads/%s", 
				c.Protocol(), c.Get("Host"), filepath.Base(notes[i].ImagePath))
		}
	}

	return c.JSON(fiber.Map{
		"notes": notes,
	})
}

// GetNote godoc
// @Summary Get a specific note
// @Description Retrieve a specific note by ID for the authenticated user
// @Tags Notes
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Note ID"
// @Success 200 {object} models.Note "Note details"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 404 {object} models.ErrorResponse "Note not found"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/notes/{id} [get]
func GetNote(c *fiber.Ctx) error {
	userID := middleware.GetUserID(c)
	noteID := c.Params("id")

	var note models.Note
	if err := database.DB.Where("id = ? AND user_id = ?", noteID, userID).First(&note).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Note not found",
		})
	}

	if note.ImagePath != "" {
		note.ImageURL = fmt.Sprintf("%s://%s/uploads/%s", 
			c.Protocol(), c.Get("Host"), filepath.Base(note.ImagePath))
	}

	return c.JSON(note)
}

// CreateNote godoc
// @Summary Create a new note
// @Description Create a new note with optional image upload using multipart form data
// @Tags Notes
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param title formData string true "Note title"
// @Param content formData string false "Note content"
// @Param image formData file false "Image file (JPEG, PNG, GIF)"
// @Success 201 {object} models.Note "Note created successfully"
// @Failure 400 {object} models.ErrorResponse "Bad request"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/notes [post]
func CreateNote(c *fiber.Ctx) error {
	userID := middleware.GetUserID(c)

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid multipart form",
		})
	}

	title := ""
	content := ""
	
	if titleValues := form.Value["title"]; len(titleValues) > 0 {
		title = titleValues[0]
	}
	if contentValues := form.Value["content"]; len(contentValues) > 0 {
		content = contentValues[0]
	}

	if title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Title is required",
		})
	}

	userUUID, _ := uuid.Parse(userID)
	note := models.Note{
		Title:   title,
		Content: content,
		UserID:  userUUID,
	}

	if files := form.File["image"]; len(files) > 0 {
		file := files[0]
		
		if !isValidImageType(file.Header.Get("Content-Type")) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid image type. Only JPEG, PNG, and GIF are allowed",
			})
		}

		ext := filepath.Ext(file.Filename)
		filename := fmt.Sprintf("%s%s", uuid.New().String(), ext)
		
		uploadsDir := "uploads"
		if err := os.MkdirAll(uploadsDir, 0755); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create uploads directory",
			})
		}

		savePath := filepath.Join(uploadsDir, filename)
		if err := c.SaveFile(file, savePath); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to save image",
			})
		}

		note.ImagePath = savePath
	}

	if err := database.DB.Create(&note).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create note",
		})
	}

	if note.ImagePath != "" {
		note.ImageURL = fmt.Sprintf("%s://%s/uploads/%s", 
			c.Protocol(), c.Get("Host"), filepath.Base(note.ImagePath))
	}

	return c.Status(fiber.StatusCreated).JSON(note)
}

// UpdateNote godoc
// @Summary Update an existing note
// @Description Update a note's title, content, and/or image using multipart form data
// @Tags Notes
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param id path string true "Note ID"
// @Param title formData string false "Note title"
// @Param content formData string false "Note content"
// @Param image formData file false "Image file (JPEG, PNG, GIF)"
// @Success 200 {object} models.Note "Note updated successfully"
// @Failure 400 {object} models.ErrorResponse "Bad request"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 404 {object} models.ErrorResponse "Note not found"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/notes/{id} [put]
func UpdateNote(c *fiber.Ctx) error {
	userID := middleware.GetUserID(c)
	noteID := c.Params("id")

	var note models.Note
	if err := database.DB.Where("id = ? AND user_id = ?", noteID, userID).First(&note).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Note not found",
		})
	}

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid multipart form",
		})
	}

	if titleValues := form.Value["title"]; len(titleValues) > 0 && titleValues[0] != "" {
		note.Title = titleValues[0]
	}
	if contentValues := form.Value["content"]; len(contentValues) > 0 {
		note.Content = contentValues[0]
	}

	if files := form.File["image"]; len(files) > 0 {
		file := files[0]
		
		if !isValidImageType(file.Header.Get("Content-Type")) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid image type. Only JPEG, PNG, and GIF are allowed",
			})
		}

		if note.ImagePath != "" {
			os.Remove(note.ImagePath)
		}

		ext := filepath.Ext(file.Filename)
		filename := fmt.Sprintf("%s%s", uuid.New().String(), ext)
		
		uploadsDir := "uploads"
		if err := os.MkdirAll(uploadsDir, 0755); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create uploads directory",
			})
		}

		savePath := filepath.Join(uploadsDir, filename)
		if err := c.SaveFile(file, savePath); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to save image",
			})
		}

		note.ImagePath = savePath
	}

	if err := database.DB.Save(&note).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update note",
		})
	}

	if note.ImagePath != "" {
		note.ImageURL = fmt.Sprintf("%s://%s/uploads/%s", 
			c.Protocol(), c.Get("Host"), filepath.Base(note.ImagePath))
	}

	return c.JSON(note)
}

// DeleteNote godoc
// @Summary Delete a note
// @Description Delete a note and its associated image file
// @Tags Notes
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Note ID"
// @Success 200 {object} models.MessageResponse "Note deleted successfully"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 404 {object} models.ErrorResponse "Note not found"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/notes/{id} [delete]
func DeleteNote(c *fiber.Ctx) error {
	userID := middleware.GetUserID(c)
	noteID := c.Params("id")

	var note models.Note
	if err := database.DB.Where("id = ? AND user_id = ?", noteID, userID).First(&note).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Note not found",
		})
	}

	if note.ImagePath != "" {
		os.Remove(note.ImagePath)
	}

	if err := database.DB.Delete(&note).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete note",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Note deleted successfully",
	})
}

func isValidImageType(contentType string) bool {
	validTypes := []string{
		"image/jpeg",
		"image/jpg", 
		"image/png",
		"image/gif",
	}

	for _, validType := range validTypes {
		if strings.ToLower(contentType) == validType {
			return true
		}
	}
	return false
}
