package config

import (
	"encoding/json"
	"os"
	"runtime"

	"github.com/BurntSushi/toml"
	"github.com/dustin/go-humanize"
	"github.com/pbnjay/memory"
)

const (
	EnvDevelopment = "development"
	EnvProduction  = "production"
)

type Config struct {
	Env              string           `toml:"env"`
	HttpServerConfig HttpServerConfig `toml:"http"`
	HttpServerStats  HttpServerStats
}

type HttpServerConfig struct {
	Port string `toml:"port"`
}

type HttpServerStats struct {
	NumCPUs     int
	TotalMemory string
}

func (cfg *Config) ToJSON() string {
	s, _ := json.MarshalIndent(cfg, "", "    ")
	return string(s)
}

func NewConfig() (*Config, error) {
	var config Config

	// Load default config
	_, err := toml.DecodeFile("config/config.toml", &config)
	if err != nil {
		return nil, err
	}

	// Override config with environment variables
	env := os.Getenv("ENV")
	if env == EnvProduction {
		config.Env = EnvProduction
	}

	port := os.Getenv("PORT")
	if port != "" {
		config.HttpServerConfig.Port = port
	}

	config.HttpServerStats = HttpServerStats{
		NumCPUs:     runtime.NumCPU(),
		TotalMemory: humanize.Bytes(memory.TotalMemory()),
	}

	return &config, nil
}
