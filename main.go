package main

import (
	"fmt"

	"school-app/database"
	"school-app/models"
	"school-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()

	err := database.DB.AutoMigrate(&models.Student{}, &models.Class{}, &models.Teacher{})
	if err != nil {
		fmt.Println("Migration failed", err)
	} else {
		fmt.Println("database migrated successfully")
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "This is pranish"})
	})
	routes.StudentsRoutes(r)
	r.Run(":8080")

}
