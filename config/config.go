package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port  string `yaml:"port" mapstructure:"port"`
	Mongo Mongo  `yaml:"mongo" mapstructure:"mongo"`
}

type Mongo struct {
	Host         string `yaml:"host" mapstructure:"host"`
	DatabaseName string `yaml:"database_name" mapstructure:"database_name"`
}

func Load() *Config {
	var cfg = &Config{}
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
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
