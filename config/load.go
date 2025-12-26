package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

func Load() Config {
	var cfg Config

	// load configs
	var raw = koanf.New(".")
	if err := raw.Load(file.Provider("config.yml"), yaml.Parser()); err != nil {
		log.Fatalf("can't load config from config.yml file : %v \n", err)
	}
	if err := raw.Load(file.Provider(".env"), dotenv.Parser()); err != nil {
		log.Fatalf("can't load config from .env file : %v \n", err)

	}

	// map all key and set to nromalize
	var normalize = koanf.New(".")
	for key, value := range raw.All() {
		nKey := strings.ToLower(strings.Replace(key, "_", ".", 1))
		normalize.Set(nKey, value)
	}

	// Unmashal configs to golang struct 
	if err := normalize.Unmarshal("", &cfg); err != nil {
		fmt.Printf("can't Unmarshal config : %v \n", err)
	}
	return cfg
}
