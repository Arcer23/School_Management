package controllers

import (
	"net/http"
	"school-app/database"
	"school-app/models"

	"github.com/gin-gonic/gin"
)

func CreateTeacher(c *gin.Context) {
	var teacher models.Teacher

	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&teacher)
	c.JSON(http.StatusOK, gin.H{"message": "Teacher Created Successfully"})
}

func GetAllTeachers(c *gin.Context) {
	var teachers []models.Teacher

	database.DB.Find(&teachers)
	c.JSON(http.StatusOK, gin.H{"message": "All teachers fetched Sucessfully"})
	
}

func GetTeacherById(c *gin.Context){
	id := c.Param("id")
	var teacher models.Teacher
	if err := c.ShouldBindJSON(&teacher) ; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message" : err})
		return 
	}

	if err := database.DB.Find(&teacher , id).Error ; err !=nil{
		c.JSON(http.StatusNotFound,gin.H{"message" : "Teacher not found "})
		return
	}
	database.DB.Save(&teacher)
	c.JSON(http.StatusOK, gin.H{"message": "Teacher fetched successfully", "teacher": teacher})
}