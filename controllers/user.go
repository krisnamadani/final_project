package controllers

import (
	"final_project/models"
	"final_project/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUserInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Age      uint   `json:"age" binding:"required,gte=8"`
}

type LoginUserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func RegisterUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input RegisterUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dt := time.Now()

	users := models.User{
		Username:   input.Username,
		Email:      input.Email,
		Password:   string(hashedPassword),
		Age:        input.Age,
		Created_At: dt.Format("2006-01-02"),
		Updated_At: dt.Format("2006-01-02"),
	}

	if err := db.Create(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":       users.ID,
		"username": users.Username,
		"email":    users.Email,
		"age":      users.Age,
	})
}

func LoginUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input LoginUserInput

	user := models.User{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("email = ?", input.Email).Take(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(input.Password), []byte(user.Password)); err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func PutUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var users []models.User
	if err := db.Where("id = ?", c.Param("id")).First(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"username": users[0].Username, "email": users[0].Email})
}

func DeleteUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var users []models.User
	db.Delete(&users)
	c.JSON(http.StatusOK, gin.H{"message": "Your account has been successfully deleted"})
}
