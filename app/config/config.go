package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	App struct {
		Version string `yaml:"version"`
	} `yaml:"version"`
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Url string `yaml:"url"`
	} `yaml:"database"`
}

func Load(configPath string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
