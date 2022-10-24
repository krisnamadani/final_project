package controllers

import (
	"final_project/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateCommentInput struct {
	Photo_ID uint   `json:"photo_id"`
	Message  string `json:"message" binding:"required"`
}

type UpdateCommentInput struct {
	Message string `json:"message" binding:"required"`
}

func CreateComment(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input CreateCommentInput

	// validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dt := time.Now()
	comment := models.Comment{
		User_ID:    1,
		Photo_ID:   input.Photo_ID,
		Message:    input.Message,
		Created_At: dt.Format("2006-01-02"),
		Updated_At: dt.Format("2006-01-02"),
	}

	db.Create(&comment)

	c.JSON(http.StatusCreated, gin.H{
		"id":         comment.ID,
		"user_id":    comment.User_ID,
		"photo_id":   comment.Photo_ID,
		"message":    comment.Message,
		"created_at": comment.Created_At,
	})
}

func GetComment(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var comment []models.Comment

	if err := db.Preload("User").Preload("Photo").Find(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comment)
}

func UpdateComment(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var comment models.Comment

	// get model if exist
	if err := db.Where("id = ?", c.Param("commentId")).First(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validate input
	var input UpdateCommentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&comment).Updates(input)

	c.JSON(http.StatusOK, gin.H{
		"id":         comment.ID,
		"user_id":    comment.User_ID,
		"photo_id":   comment.Photo_ID,
		"message":    comment.Message,
		"updated_at": comment.Updated_At,
	})
}

func DeleteComment(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var comment models.Comment

	// get model if exist
	if err := db.Where("id = ?", c.Param("commentId")).First(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Delete(&comment)

	c.JSON(http.StatusOK, gin.H{"message": "Your comment has been successfully deleted"})
}
