package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/amjadjibon/microservices/title/model"
)

type createTitleInput struct {
	TitleName   string `json:"title_name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Year        int    `json:"year" binding:"required"`
}

type createTitleResponse struct {
	ID          int       `json:"title_id"`
	TitleName   string    `json:"title_name"`
	Description string    `json:"description"`
	Year        int       `json:"year"`
	CreatedAt   time.Time `json:"created_at"`
	UpdateAt    time.Time `json:"updated_at"`
}

func (h *TitleHandler) CreateTitle(c *gin.Context) {
	input := createTitleInput{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	title := &model.Title{
		TitleName:   input.TitleName,
		Description: input.Description,
		Year:        input.Year,
		CreatedAt:   time.Now(),
		UpdateAt:    time.Now(),
	}

	title, err := h.repo.CreateTitle(c, title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createTitleResponse{
		ID:          title.ID,
		TitleName:   title.TitleName,
		Description: title.Description,
		Year:        title.Year,
		CreatedAt:   title.CreatedAt,
		UpdateAt:    title.UpdateAt,
	})
}
