package models

import "github.com/jinzhu/gorm"

type Edge struct {
	gorm.Model
	Name   string
	From   Node
	To     Node
	Weight int64
}
