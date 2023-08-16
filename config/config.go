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

	Cache struct {
		Address  string `yaml:"ADDRESS"`
		Password string `yaml:"Password"`
		DB       int    `yaml:"DB"`
	} `yaml:"CACHE"`

	ReportMail struct {
		Address string `yaml:"ADDRESS"`
		Success bool `yaml:"SUCCESS"`
		Error   bool `yaml:"ERROR"`
	} `yaml:"REPORT"`

	Mail struct {
		User     string `yaml:"USER"`
		Password string `yaml:"PASSWORD"`
		Host     string `yaml:"HOST"`
		Port     int    `yaml:"PORT"`
	} `yaml:"MAIL"`
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

	cfg.Port = getFlags()
	Data = cfg

	return cfg
}
