package model

import (
	"database/sql"
	"time"
)

// Title represents the Titles table
type Title struct {
	ID          int          `json:"title_id"`
	TitleName   string       `json:"title_name"`
	Description string       `json:"description"`
	Year        int          `json:"year"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdateAt    time.Time    `json:"updated_at"`
	DeletedAt   sql.NullTime `json:"deleted_at"`
}

// Language represents the Languages table
type Language struct {
	ID           int          `json:"language_id"`
	LanguageName string       `json:"language_name"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdateAt     time.Time    `json:"updated_at"`
	DeletedAt    sql.NullTime `json:"deleted_at"`
}

// TitleType represents the Title_Type table
type TitleType struct {
	ID        int          `json:"type_id"`
	TypeName  string       `json:"type_name"`
	CreatedAt time.Time    `json:"created_at"`
	UpdateAt  time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}

// TitleContent represents the Title_Content table
type TitleContent struct {
	ID          int          `json:"content_id"`
	TitleID     int          `json:"title_id"`
	ContentName string       `json:"content_name"`
	ContentType int          `json:"content_type"`
	Year        int          `json:"year"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdateAt    time.Time    `json:"updated_at"`
	DeltedAt    sql.NullTime `json:"deleted_at"`
}

// TitleContentLanguage represents the Title_Content_Languages table
type TitleContentLanguage struct {
	ContentID  int `json:"content_id"`
	LanguageID int `json:"language_id"`
}
