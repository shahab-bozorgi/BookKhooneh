package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model

	BookID uint `gorm:"not null"`
	UserID uint `gorm:"not null"`

	Rating  int    `gorm:"not null;check:rating >= 1 AND rating <= 10"`
	Comment string `gorm:"type:text"`

	Book Book `gorm:"constraint:OnDelete:CASCADE;"`
	User User `gorm:"constraint:OnDelete:CASCADE;"`
}
