package config

import (
	"github.com/jinzhu/gorm"
)

// AppContext => General Application Context
type AppContext struct {
	Config Specs
	DB     *gorm.DB
}
