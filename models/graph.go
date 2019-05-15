package models

import "github.com/jinzhu/gorm"

// Graph : graph
type Graph struct {
	gorm.Model
	ID    uint `gorm:"primary_key;auto_increment:false"`
	Name  string
	Nodes []Node
	Edges []Edge
}
