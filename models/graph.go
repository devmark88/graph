package models

// Graph : graph
type Graph struct {
	ID    uint `gorm:"primary_key;auto_increment:false"`
	Name  string
	Nodes []Node
}
