package models

import "gorm.io/gorm"

type Coffee struct {
	gorm.Model
	Name        string	`gorm:"unique"`
	Description string
	Link        string
	Image       string
}
