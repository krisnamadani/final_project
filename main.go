package main

import (
	"final_project/controllers"
	"final_project/middlewares"
	"final_project/models"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func SetupDB() *gorm.DB {
	USER := "root"
	PASS := ""
	HOST := "localhost"
	PORT := "3306"
	DBNAME := "final_project"

	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)

	db, err := gorm.Open("mysql", URL)

	if err != nil {
		panic(err.Error())
	}

	return db
}

func main() {
	r := gin.Default()

	db := SetupDB()

	db.AutoMigrate(
		&models.User{},
		&models.Photo{},
		&models.Comment{},
		&models.SocialMedia{},
	)

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.POST("/users/register", controllers.RegisterUser)
	r.POST("/users/login", controllers.LoginUser)

	r.Use(middlewares.JwtAuthMiddleware())

	r.PUT("/users/:id", controllers.PutUser)
	r.DELETE("/users", controllers.DeleteUser)

	r.POST("/photos", controllers.CreatePhoto)
	r.GET("/photos", controllers.GetPhoto)
	r.PUT("/photos/:photoId", controllers.UpdatePhoto)
	r.DELETE("/photos/:photoId", controllers.DeletePhoto)

	r.POST("/comments", controllers.CreateComment)
	r.GET("/comments", controllers.GetComment)
	r.PUT("/comments/:commentId", controllers.UpdateComment)
	r.DELETE("/comments/:commentId", controllers.DeleteComment)

	r.POST("/socialmedias", controllers.CreateSocialMedia)
	r.GET("/socialmedias", controllers.GetSocialMedia)
	r.PUT("/socialmedias/:socialMediaId", controllers.UpdateSocialMedia)
	r.DELETE("/socialmedias/:socialMediaId", controllers.DeleteSocialMedia)

	r.Run()
}
