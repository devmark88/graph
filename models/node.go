package models

// Node : graph node
type Node struct {
	ID      uint `gorm:"primary_key;auto_increment:false"`
	Name    string
	GraphID uint
}

// Find => find index of element in nodes by id
func Find(id uint, nodes []Node) *Node {
	for i, n := range nodes {
		if n.ID == id {
			return &nodes[i]
		}
	}
	return nil
}
