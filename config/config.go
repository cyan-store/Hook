package config

import (
	"os"

	"github.com/cyan-store/hook/log"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Port       int    `yaml:"PORT"`
	DSN        string `yaml:"DSN"`
	StripeKey  string `yaml:"STRIPE_KEY"`
	StripeHook string `yaml:"STRIPE_HOOK"`
	Cache      struct {
		Address  string `yaml:"ADDRESS"`
		Password string `yaml:"Password"`
		DB       int    `yaml:"DB"`
	} `yaml:"CACHE"`
}

var Data Config

func Load() Config {
	config, err := os.ReadFile("config.yaml")
	cfg := Config{}

	if err != nil {
		log.Error.Println("[config.Load] Could not find config.yaml.")
		os.Exit(1)
	}

	if err = yaml.Unmarshal(config, &cfg); err != nil {
		log.Error.Println("[config.Load] Could not unmarshal config -", err)
		os.Exit(1)
	}

	Data = cfg
	return cfg
}
