package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	BookID  uint
	UserID  uint
	Rating  int
	Comment string

	Book Book
	User User
}
