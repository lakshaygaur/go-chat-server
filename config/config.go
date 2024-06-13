package config

import (
	"github.com/stvp/go-toml-config"
)

// var (
// 	country            = config.String("country", "Unknown")
// 	atlantaEnabled     = config.Bool("atlanta.enabled", false)
// 	alantaPopulation   = config.Int("atlanta.population", 0)
// 	atlantaTemperature = config.Float("atlanta.temperature", 0)
// )

func SetVars() {
	config.Parse("/home/lakshay/go-playground/chat-server/config/config.conf")
}