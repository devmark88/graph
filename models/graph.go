package models

import "github.com/jinzhu/gorm"

// Graph : graph
type Graph struct {
	gorm.Model
	Name  string
	Nodes []Node
	Edges []Edge
}
