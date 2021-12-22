package model

import "gorm.io/gorm"

type Follow struct {
	gorm.Model
	Email    string
	PinpinId uint `gorm:"unique"`
	IsFollow bool
}
