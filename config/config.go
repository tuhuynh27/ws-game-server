package config

import (
	"bytes"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Mongo Mongo `yaml:"mongo" mapstructure:"mongo"`
}

type Mongo struct {
	Host         string `yaml:"host" mapstructure:"host"`
	DatabaseName string `yaml:"database_name" mapstructure:"database_name"`
}

var defaultConfig = []byte(`
mongo:
    host: mongodb://localhost:27017
    database_name: odd-game
`)

func Load() *Config {
	var cfg = &Config{}
	viper.SetConfigType("yaml")

	err := viper.ReadConfig(bytes.NewBuffer(defaultConfig))
	if err != nil {
		log.Fatal("Failed to read viper config", err)
	}

	viper.AutomaticEnv()

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatal("Failed to unmarshal config", err)
	}

	log.Println("Config loaded")
	return cfg
}
