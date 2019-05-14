package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Specs => global app config
type Specs struct {
	Debug            bool   `split_words:"true" default:"false"`
	Port             int    `split_words:"true" default:"9999"`
	DatabaseHost     string `split_words:"true" default:"127.0.0.1"`
	DatabaseUsername string `split_words:"true" default:"mark" `
	DatabasePort     string `split_words:"true" default:"5432"`
	DatabasePassword string `split_words:"true" default:"123456"`
	DatabaseName     string `split_words:"true" default:"unireg"`
}

// InitSpecs => initial config
func InitSpecs() Specs {
	var c Specs
	err := envconfig.Process("unireg", &c)
	if err != nil {
		panic(err)
	}
	return c
}
