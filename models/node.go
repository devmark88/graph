package models
import "github.com/jinzhu/gorm"

// Node : graph node
type Node struct {
	gorm.Model
	Name  string
	GraphID uint
	Childs []Node
}