package controllers

import (
	"final_project/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreatePhotoInput struct {
	Title     string `json:"title" binding:"required"`
	Caption   string `json:"caption" binding:"required"`
	Photo_URL string `json:"photo_url" binding:"required"`
}

type UpdatePhotoInput struct {
	Title     string `json:"title" binding:"required"`
	Caption   string `json:"caption" binding:"required"`
	Photo_URL string `json:"photo_url" binding:"required"`
}

func CreatePhoto(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input CreatePhotoInput

	// validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dt := time.Now()
	photo := models.Photo{
		Title:      input.Title,
		Caption:    input.Caption,
		Photo_URL:  input.Photo_URL,
		User_ID:    1,
		Created_At: dt.Format("2006-01-02"),
		Updated_At: dt.Format("2006-01-02"),
	}

	db.Create(&photo)

	c.JSON(http.StatusCreated, gin.H{
		"id":         photo.ID,
		"title":      photo.Title,
		"caption":    photo.Caption,
		"photo_url":  photo.Photo_URL,
		"user_id":    photo.User_ID,
		"created_at": photo.Created_At,
	})
}

func GetPhoto(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var photo []models.Photo

	if err := db.Preload("User").Find(&photo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, photo)
}

func UpdatePhoto(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var photo models.Photo

	// get model if exist
	if err := db.Where("id = ?", c.Param("photoId")).First(&photo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validate input
	var input UpdatePhotoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&photo).Updates(input)

	c.JSON(http.StatusOK, gin.H{
		"id":         photo.ID,
		"title":      photo.Title,
		"caption":    photo.Caption,
		"photo_url":  photo.Photo_URL,
		"user_id":    photo.User_ID,
		"updated_at": photo.Updated_At,
	})
}

func DeletePhoto(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var photo models.Photo

	// get model if exist
	if err := db.Where("id = ?", c.Param("photoId")).First(&photo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Delete(&photo)

	c.JSON(http.StatusOK, gin.H{"message": "Your photo has been successfully deleted"})
}
