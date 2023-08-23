package models

import "gorm.io/gorm"

type League struct {
	gorm.Model
	Name      string
	CountryID uint
	Country   Country
	Teams     []Team
}
