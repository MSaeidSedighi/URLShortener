package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Email     string `gorm:"uniqueIdex; not null"`
	Password  string
	FirstName string
	LastName  string
}
