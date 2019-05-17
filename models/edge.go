package models

// Edge => the place of nodes
type Edge struct {
	ID     uint `gorm:"primary_key;auto_increment:false"`
	From   Node `gorm:"association_foreignkey:FromID"`
	To     Node `gorm:"association_foreignkey:ToID"`
	FromID uint
	ToID   uint
	Weight float64
}
