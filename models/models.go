package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"-" gorm:"not null"`
	Name      string    `json:"name" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Notes     []Note    `json:"notes,omitempty" gorm:"foreignKey:UserID"`
}

type AuthUser struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Note struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Title       string    `json:"title" gorm:"not null"`
	Content     string    `json:"content"`
	ImagePath   string    `json:"image_path,omitempty"`
	ImageURL    string    `json:"image_url,omitempty" gorm:"-"`
	UserID      uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type RegisterRequest struct {
	Name     string `json:"name" validate:"required,min=2"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type CreateNoteRequest struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content"`
}

type UpdateNoteRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Standard API Response wrapper following REST best practices
type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// Auth response payload
type AuthData struct {
	Token string   `json:"token"`
	User  AuthUser `json:"user"`
}

// Notes list response payload
type NotesData struct {
	Notes []Note `json:"notes"`
	Count int    `json:"count"`
}

// Single note response payload (for create, update, get)
type NoteData struct {
	Note Note `json:"note"`
}

// Success message response payload
type MessageData struct {
	Message string `json:"message"`
}

// Specific response types for Swagger documentation
type AuthSuccessResponse struct {
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Data    AuthData `json:"data"`
}

type NotesSuccessResponse struct {
	Status  string    `json:"status"`
	Message string    `json:"message"`
	Data    NotesData `json:"data"`
}

type NoteSuccessResponse struct {
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Data    NoteData `json:"data"`
}

type MessageSuccessResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    MessageData `json:"data"`
}

type ErrorResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}
