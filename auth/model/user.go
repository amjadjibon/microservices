package model

import (
	"database/sql"
	"time"
)

// User represents the User table in the database.
type User struct {
	ID         int          `json:"id"`
	Username   string       `json:"username"`
	Name       string       `json:"name"`
	Email      string       `json:"email"`
	IsVerified bool         `json:"is_verified"`
	Gender     string       `json:"gender"`
	Password   string       `json:"password"`
	Role       string       `json:"role"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
	DeletedAt  sql.NullTime `json:"deleted_at"`
}

// UserPermission represents the UserPermission table in the database (if added).
type UserPermission struct {
	ID             int          `json:"id"`
	UserID         int          `json:"user_id"`
	PermissionName string       `json:"permission_name"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
	DeletedAt      sql.NullTime `json:"deleted_at"`
}
