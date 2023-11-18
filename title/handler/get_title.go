package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *TitleHandler) GetTitleByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	title, err := h.repo.GetTitleById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title_id":    title.ID,
		"title_name":  title.TitleName,
		"description": title.Description,
		"year":        title.Year,
		"created_at":  title.CreatedAt,
		"updated_at":  title.UpdateAt,
	})
}
