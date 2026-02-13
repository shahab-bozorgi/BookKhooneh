package domain

import "gorm.io/gorm"

type Book struct {
	gorm.Model

	Title       string `gorm:"not null"`
	Author      string `gorm:"not null"`
	Description string `gorm:"type:text"`

	UserID *uint
	User   *User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	Reviews []Review `gorm:"constraint:OnDelete:CASCADE;"`
}
