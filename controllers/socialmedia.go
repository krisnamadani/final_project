package controllers

import (
	"final_project/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateSocialMediaInput struct {
	Name             string `json:"name" binding:"required"`
	Social_Media_URL string `json:"social_media_url" binding:"required"`
}

type UpdateSocialMediaInput struct {
	Name             string `json:"name" binding:"required"`
	Social_Media_URL string `json:"social_media_url" binding:"required"`
}

func CreateSocialMedia(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input CreateSocialMediaInput

	// validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dt := time.Now()
	socialmedia := models.SocialMedia{
		Name:             input.Name,
		Social_Media_URL: input.Social_Media_URL,
		User_ID:          1,
		Created_At:       dt.Format("2006-01-02"),
		Updated_At:       dt.Format("2006-01-02"),
	}

	db.Create(&socialmedia)

	c.JSON(http.StatusCreated, gin.H{
		"id":               socialmedia.ID,
		"name":             socialmedia.Name,
		"social_media_url": socialmedia.Social_Media_URL,
		"user_id":          socialmedia.User_ID,
		"created_at":       socialmedia.Created_At,
	})
}

func GetSocialMedia(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var socialmedia []models.SocialMedia

	if err := db.Preload("User").Find(&socialmedia).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, socialmedia)
}

func UpdateSocialMedia(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var socialmedia models.SocialMedia

	// get model if exist
	if err := db.Where("id = ?", c.Param("socialMediaId")).First(&socialmedia).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validate input
	var input UpdateSocialMediaInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&socialmedia).Updates(input)

	c.JSON(http.StatusOK, gin.H{
		"id":               socialmedia.ID,
		"name":             socialmedia.Name,
		"social_media_url": socialmedia.Social_Media_URL,
		"user_id":          socialmedia.User_ID,
		"updated_at":       socialmedia.Updated_At,
	})
}

func DeleteSocialMedia(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var socialmedia models.SocialMedia

	// get model if exist
	if err := db.Where("id = ?", c.Param("socialMediaId")).First(&socialmedia).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Delete(&socialmedia)

	c.JSON(http.StatusOK, gin.H{"message": "Your social media has been successfully deleted"})
}
