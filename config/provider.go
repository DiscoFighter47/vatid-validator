package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/spf13/viper"
)

// provider interface for config provider
type provider interface {
	read(*viper.Viper) error
}

type providerFunc func(*viper.Viper) error

func (f providerFunc) read(v *viper.Viper) error {
	return f(v)
}

var providers = map[string]provider{
	"file": providerFunc(fileProvider),
}

// fileProvider reads config from a file
func fileProvider(v *viper.Viper) error {
	e := struct {
		File string `env:"CONFIG_FILE,required"`
	}{}
	if err := env.Parse(&e); err != nil {
		return fmt.Errorf("unable to process environment variable: %w", err)
	}
	v.SetConfigFile(e.File)
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("unable to read config file: %w", err)
	}
	return nil
}
