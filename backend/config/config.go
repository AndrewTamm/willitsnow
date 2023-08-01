package config

import (
	"embed"
	"flag"
	"fmt"
	"github.com/caarlos0/env/v7"
	"gopkg.in/yaml.v2"
	"io/fs"
)

// Name of the struct tag used in examples.
const tagName = "envkey"

//go:embed yaml
var configFS embed.FS

type Config struct {
	Server struct {
		Host string `yaml:"host" env:"ServerHost"`
		Port string `yaml:"port" env:"ServerPort"`
		Cors bool   `yaml:"cors" eng:"ServerRequireCors"`
	} `yaml:"server"`
	Weather struct {
		Location string `yaml:"location" env:"WeatherLocation"`
	} `yaml:"weather"`
}

func NewConfig(configPath string) (*Config, error) {
	config := &Config{}

	yamls, err := fs.Sub(configFS, "yaml")
	if err != nil {
		return nil, err
	}

	file, err := yamls.Open(configPath)

	if err != nil {
		return nil, err
	}
	defer func(file fs.File) {
		_ = file.Close()
	}(file)

	if err := yaml.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}

	err = env.Parse(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func ParseFlags() (string, error) {
	// String that contains the configured configuration path
	var configPath string

	// Set up a CLI flag called "-config" to allow users to supply the configuration file
	flag.StringVar(&configPath, "config", "config.yaml", "path to config file")

	flag.Parse()

	if err := ValidateConfigPath(configPath); err != nil {
		return "", err
	}

	return configPath, nil
}

func ValidateConfigPath(path string) error {
	yamls, err := fs.Sub(configFS, "yaml")
	if err != nil {
		return err
	}

	f, err := yamls.Open(path)
	if err != nil {
		return err
	}

	s, err := f.Stat()
	if err != nil {
		return err
	}

	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}
