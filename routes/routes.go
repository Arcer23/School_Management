package routes

import (
	"school-app/controllers"

	"github.com/gin-gonic/gin"
)

func StudentsRoutes(r *gin.Engine) {
	students := r.Group("/students")

	{
		students.GET("/:id", controllers.GetStudentById)
		students.GET("/", controllers.GetStudents)
		students.POST("/", controllers.CreateStudent)
		students.DELETE("/", controllers.DeleteAllStudents)
		students.DELETE("/:id", controllers.DeleteStudentById)
		students.PUT("/:id", controllers.UpdateStudent)
	}
	teachers := r.Group("/teachers")
	{
		teachers.GET("/", controllers.GetAllTeachers)
		teachers.POST("/", controllers.CreateTeacher)
		teachers.GET("/:id", controllers.GetTeacherById)
	}
}
