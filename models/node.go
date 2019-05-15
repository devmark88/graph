package models

import "github.com/jinzhu/gorm"

// Node : graph node
type Node struct {
	gorm.Model
	ID      uint `gorm:"primary_key;auto_increment:false"`
	Name    string
	GraphID uint
	NodeID  *uint
	Nodes   []Node
}

// Find => find index of element in nodes by id
func Find(id uint, nodes []Node) int {
	for i, n := range nodes {
		if n.ID == id {
			return i
		}
	}
	return -1
}
