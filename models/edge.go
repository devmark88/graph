package models

// Edge => the place of nodes
type Edge struct {
	ID     uint `gorm:"primary_key;auto_increment:false"`
	From   uint
	To     uint
	Weight float64
}
