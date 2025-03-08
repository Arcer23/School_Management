package controllers

import (
	"net/http"
	"school-app/database"
	"school-app/models"

	"github.com/gin-gonic/gin"
)

func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}

func GetStudents(c *gin.Context) {
	var students []models.Student

	database.DB.Find(&students)
	c.JSON(http.StatusOK, students)

}

func GetStudentById(c *gin.Context) {
	id := c.Param("id")
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	database.DB.Save(&student)
	c.JSON(http.StatusOK, student)

}
func DeleteAllStudents(c *gin.Context) {
	result := database.DB.Where("1 = 1").Delete(&models.Student{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "All students delete d successfully"})
}
func DeleteStudentById(c *gin.Context) {
	id := c.Param("id")
	var student models.Student

	if err := database.DB.First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	database.DB.Delete(&student)
	c.JSON(http.StatusOK, gin.H{"message": "Student Deleted Successfully"})
}
func UpdateStudent(c *gin.Context) {
	id := c.Param("id")

	var updatedStudent models.Student

	if err := c.ShouldBindJSON(&updatedStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}

	var student models.Student
	if err := database.DB.First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student Not Found"})
		return
	}

	student.Name = updatedStudent.Name
	student.ClassId = updatedStudent.ClassId
	student.Email = updatedStudent.Email

	if err := database.DB.Save(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save the student"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Student updated Successfully"})
}
