package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Username string `gorm:"uniqueIndex;not null"`
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"type:varchar(10);default:'user'"`

	Reviews []Review `gorm:"constraint:OnDelete:CASCADE;"`
	Books   []Book
}
