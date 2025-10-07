package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Gthamsrim1/forms-backend/db"
	"github.com/Gthamsrim1/forms-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type QuestionInput struct {
	Text       string   `json:"text" binding:"required"`
	Type       string   `json:"type" binding:"required"`
	IsRequired bool     `json:"is_required"`
	Options    []string `json:"options"`
	Order      int      `json:"order"`
}

type FormInput struct {
	Title     string          `json:"title" binding:"required"`
	Questions []QuestionInput `json:"questions"`
}

func CreateForm(c *gin.Context) {
	var input FormInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIDValue, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	userID, _ := uuid.Parse(userIDValue.(string))

	form := models.Form{
		Title:   input.Title,
		OwnerID: userID,
	}

	for _, q := range input.Questions {
		optionsJSON := "null"
		if len(q.Options) > 0 {
			optionsJSONBytes, _ := json.Marshal(q.Options)
			optionsJSON = string(optionsJSONBytes)
		}
		form.Questions = append(form.Questions, models.Question{
			Text:        q.Text,
			Type:        q.Type,
			IsRequired:  q.IsRequired,
			OptionsJSON: optionsJSON,
			Order:       q.Order,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		})
	}

	if err := db.DB.Create(&form).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create form"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"form_id": form.ID})
}
