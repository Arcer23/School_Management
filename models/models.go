package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-"`
	Role     string `json:"role"`
}
type Student struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey" json:"user_id"`
	User     User   `gorm:"foreignKey:ID;not null"`
	ClassId  uint   `json:"class_id"`
	Class    Class  `gorm:"foreignKey:ClassId"`
	RollNo   string `json:"roll_no"`
	ParentId uint   `json:"parent_id"`
	Parent   User   `gorm:"foreignKey:ParentId"`
}

type Teacher struct {
	gorm.Model
	UserID  uint   `json:"user_id"`
	User    User   `gorm:"foreignKey:ID"`
	Subject string `gorm:"not null" json:"subject"`
}

type Class struct {
	gorm.Model
	Name        string    `gorm:"unique;not null" json:"name"`
	Description string    `json:"description"`
	Students    []Student `gorm:"foreignKey:ClassId"`
	TeacherId   uint
}

type Attendence struct {
	gorm.Model
	StudentID uint    `json:"student_id"`
	Student   Student `gorm:"foreignKey:StudentID"`
	Date      string  `json:"date"`
	Status    string  `json:"status"`
}

type Fee struct {
	gorm.Model
	StudentID uint    `json:"student_id"`
	Student   Student `gorm:"foreignKey:StudentID"`
	Amount    float64 `json:"amount"`
	Status    string  `json:"status"`
	DueDate   string  `json:"due_date"`
}

type Exam struct {
	gorm.Model
	Title     string  `json:"title"`
	SubjectID uint    `json:"subject_id"`
	Subject   Subject `gorm:"foreignKey:SubjectID"`
	Date      string  `json:"date"`
}

type Subject struct {
	gorm.Model
	Name    string `json:"name"`
	Code    string `json:"code"`
	ClassID uint   `json:"class_id"`
	Class   Class  `gorm:"foreignKey:ClassID"`
}
