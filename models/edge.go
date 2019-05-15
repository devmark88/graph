package models

import "github.com/jinzhu/gorm"

type Edge struct {
	gorm.Model
	ID     uint `gorm:"primary_key;auto_increment:false"`
	Name   string
	From   uint
	To     uint
	Weight float64
}
