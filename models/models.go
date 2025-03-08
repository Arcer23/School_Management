package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	ID      uint   `gorm:"primaryKey`
	Name    string `gorm:"size:255;not null`
	Email   string `gorm:"uniqueIndex;not null"`
	ClassId uint
}

type Teacher struct {
	gorm.Model
	Name    string `gorm:"not null"`
	Email   string `gorm:"uniqueIndex;not null`
	Subject string `gorm:"not null"`
}

type Class struct {
	gorm.Model
	Name      string `gorm:"unique;not null`
	TeacherId uint
}
