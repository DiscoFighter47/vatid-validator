package config

import (
	"log"
	"os"
	"sync"

	"github.com/spf13/viper"
	"gopkg.in/validator.v2"
)

// Config holds all required configurations'
type Config struct {
	App    App
	Euvies Euvies
}

// App related config
type App struct {
	Port int64 `validate:"min=3000,max=9999"`
}

// Euvies client config
type Euvies struct {
	Timeout int `validate:"nonzero"`
}

var cnf = Config{}
var once = sync.Once{}

func GetConfig() Config {
	once.Do(func() {
		v := viper.New()
		key := os.Getenv("CONFIG_PROVIDER")
		if key == "" {
			log.Fatalln("environment variable missing: [CONFIG_PROVIDER]")
		}
		p, ok := providers[key]
		if !ok {
			log.Fatalln("invalid config provider:", key)
		}
		if err := p.read(v); err != nil {
			log.Fatalln("unable to read config:", err)
		}
		if err := v.Unmarshal(&cnf); err != nil {
			log.Fatalln("unable to unmarshal config:", err)
		}
		if err := validator.Validate(&cnf); err != nil {
			log.Fatalln("invalid config:", err)
		}
	})
	return cnf
}
